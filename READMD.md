# Go Library

## Test Init

```mod
github.com/zhangji2/golibrary v1.0.1
github.com/zhangji2/golibrary latest
```

```bash
go get github.com/zhangji2/golibrary@main
go mod tidy
go run ./main.go

# 运行当前包所有单测
go test -v

# 运行并显示覆盖率
go test -v -cover
```
