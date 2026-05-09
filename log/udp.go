package log

import (
	"bytes"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

// ############################################################
// 生产级 UDP Hook：服务器挂了 / 断网 完全不影响业务
// ############################################################

// UDPHook 安全异步UDP日志Hook
type UDPHook struct {
	addr   string
	conn   net.Conn
	mu     sync.RWMutex
	levels []logrus.Level
	sendCh chan []byte // 异步队列，防止阻塞
	closed bool
}

// NewUDPHook 创建安全的UDP Hook
func NewUDPHook(addr string, levels []logrus.Level) *UDPHook {
	hook := &UDPHook{
		addr:   addr,
		levels: levels,
		sendCh: make(chan []byte, 1000), // 最多缓冲1000条日志（防OOM）
	}

	// 后台协程异步发送，业务完全不阻塞
	go hook.sendLoop()

	// 首次立即尝试建连，避免仅靠 ticker 时首包要等一个周期
	go hook.tryDialOnce()

	// 自动重连
	go hook.autoReconnect()

	return hook
}

// Levels 实现 logrus.Hook
func (h *UDPHook) Levels() []logrus.Level {
	return h.levels
}

// Fire 日志入口（只入队，不执行任何IO）
func (h *UDPHook) Fire(entry *logrus.Entry) error {
	// 序列化日志
	bs, err := entry.Logger.Formatter.Format(entry)
	if err != nil {
		return nil
	}
	bs = append(bytes.TrimSpace(bs), '\n')

	// 非阻塞入队（队列满直接丢弃，绝不阻塞业务）
	select {
	case h.sendCh <- bs:
	default:
		// 队列满了丢弃，不影响业务
	}
	return nil
}

// ############################################################
// 下面是内部安全发送逻辑
// ############################################################

// sendLoop 异步发送协程
func (h *UDPHook) sendLoop() {
	for bs := range h.sendCh {
		h.mu.RLock()
		conn := h.conn
		h.mu.RUnlock()

		if conn == nil {
			fmt.Println("[UDP_HOOK] 无连接，跳过发送")
			continue
		}

		// 最多50ms超时，绝不卡住
		_ = conn.SetWriteDeadline(time.Now().Add(50 * time.Millisecond))
		_, _ = conn.Write(bs)

		fmt.Println(string(bs))
	}
}

// autoReconnect 自动重连（后台）
func (h *UDPHook) autoReconnect() {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for range ticker.C {

		// fmt.Println("[UDP_HOOK] autoReconnect: 心跳:3秒跳一下:", time.Now())

		h.mu.RLock()
		isClosed := h.closed
		h.mu.RUnlock()

		if isClosed {
			return
		}

		// 无连接 或 连接失效 → 重建
		h.mu.RLock()
		valid := h.conn != nil
		h.mu.RUnlock()

		if !valid {
			h.tryDialOnce()
		}
	}
}

// tryDialOnce 在 conn 为空且未关闭时尝试建立 UDP 连接（与 autoReconnect 共用逻辑）
func (h *UDPHook) tryDialOnce() {
	h.mu.RLock()
	if h.closed || h.conn != nil {
		h.mu.RUnlock()
		return
	}
	h.mu.RUnlock()

	fmt.Println("[UDP_HOOK] tryDialOnce: 尝试连接:", h.addr)
	conn, err := net.DialTimeout("udp", h.addr, 1*time.Second)
	if err != nil {
		return
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.closed {
		_ = conn.Close()
		return
	}
	if h.conn != nil {
		_ = conn.Close()
		return
	}
	h.conn = conn
	fmt.Println("[UDP_HOOK] tryDialOnce: 连接成功:", h.addr)
}
