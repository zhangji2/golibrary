package log

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type LogConfigStruct struct {
	Path     string
	FileName string
	Debug    bool
	UdpIp    string
}

var LogClient *logrus.Logger

// NewLogger new logger
func NewLogger(logConfig LogConfigStruct) *logrus.Logger {
	// if Log != nil {
	// 	return Log
	// }

	maxAgeDays := 7
	maxAge := time.Duration(maxAgeDays) * 24 * time.Hour

	// 创建不同日志级别的writer
	createWriter := func(level string) io.Writer {
		// 使用时间戳而不是strftime格式
		currentTime := time.Now().Format("20060102")
		fileName := fmt.Sprintf("service.%s.%s.%s.log", logConfig.FileName, level, currentTime)
		linkName := fmt.Sprintf("service.%s.%s.current.log", logConfig.FileName, level)
		writer, err := rotatelogs.New(
			path.Join(logConfig.Path, fileName),
			rotatelogs.WithLinkName(path.Join(logConfig.Path, linkName)),
			rotatelogs.WithMaxAge(maxAge),
			rotatelogs.WithRotationTime(24*time.Hour),
		)
		if err != nil {
			fmt.Printf("创建%s日志writer失败: %v\n", level, err)
			return os.Stdout // 失败时输出到控制台
		}
		return writer
	}

	// 为每个日志级别创建独立的writer
	infoWriter := createWriter("info")
	warnWriter := createWriter("warn")
	errorWriter := createWriter("error")
	debugWriter := createWriter("debug")
	fatalWriter := createWriter("fatal")
	panicWriter := createWriter("panic")

	// 配置writer map，每个级别对应不同的文件
	writerMap := lfshook.WriterMap{
		logrus.DebugLevel: debugWriter,
		logrus.InfoLevel:  infoWriter,
		logrus.WarnLevel:  warnWriter,
		logrus.ErrorLevel: errorWriter,
		logrus.FatalLevel: fatalWriter,
		logrus.PanicLevel: panicWriter,
	}

	log := logrus.New()

	//生产环境不输出到终端
	if !logConfig.Debug {
		log.SetOutput(ioutil.Discard)
		log.SetOutput(io.Discard)
	}

	// UDP Hook（完全安全）
	if len(logConfig.UdpIp) > 0 {
		udpHook := NewUDPHook(logConfig.UdpIp, []logrus.Level{
			logrus.InfoLevel,
			logrus.WarnLevel,
			logrus.ErrorLevel,
		})
		log.AddHook(udpHook)
	}

	log.Hooks.Add(lfshook.NewHook(
		writerMap,
		&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05", PrettyPrint: false},
	))
	return log
}

// WithField 结构化字段写入
// func WithField(key string, value interface{}) *logrus.Entry {
// 	return Log.WithFields(map[string]interface{}{key: value})
// }

// WithFields .
// func WithFields(fields map[string]interface{}) *logrus.Entry {
// 	return Log.WithFields(fields)
// }
