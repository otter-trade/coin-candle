# coin-candle

基于本地文件编写的 数字货币 K 线模块。

## 使用文档

## 使用到的第三方库

github.com/handy-golang/go-tools

## 本地调试使用 go work

```bash
go work init
```

文件 `./go.work`

```go
go 1.22.4

use (
./
)

replace(
  github.com/handy-golang/go-tools => /root/handy-golang/go-tools
)
```
