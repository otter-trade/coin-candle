package global

import (
	"fmt"
	"log"

	"github.com/handy-golang/go-tools/m_log"
	"github.com/handy-golang/go-tools/m_time"
)

var (
	Log         *log.Logger // 系统日志 & 重大错误或者事件
	RunLog      *log.Logger // 运行日志
	ExchangeLog *log.Logger // 交易所 日志
)

func init_Log() {
	Log = m_log.NewLog(m_log.NewLogParam{
		Path: Path.LogPath,
		Name: "Sys",
	})
	RunLog = m_log.NewLog(m_log.NewLogParam{
		Path: Path.LogPath,
		Name: "Run",
	})
	ExchangeLog = m_log.NewLog(m_log.NewLogParam{
		Path: Path.LogPath,
		Name: "Exchange",
	})
}

// 删除 7 天之前的日志文件
func ClearLog() {
	m_log.Clear(m_log.ClearParam{
		Path:      Path.LogPath,
		ClearTime: m_time.UnixTimeInt64.Day * 7,
	})
}

func LogErr(sum ...any) {
	str := fmt.Sprintf("系统错误: %+v", sum)
	Log.Println(str)

	// 这里可以设置邮件系统用于提醒错误
}
