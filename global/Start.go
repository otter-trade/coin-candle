package global

import (
	"time"

	"github.com/handy-golang/go-tools/m_cycle"
	"github.com/handy-golang/go-tools/m_json"
)

type SysInitOpt struct {
	LogPath   string   // 日志文件存放目录，缺省值：./logs
	DataPath  string   // 数据文件存放目录，缺省值：./data
	ProxyURLs []string // []string{"http://127.0.0.1:10809"} 在拉取数据时使用的代理，默认不使用代理
}

func SysInit(opt SysInitOpt) {
	// 初始化项 准备好运行所需要的各种 目录
	init_Path(opt)

	// 初始化日志系统
	init_Log()
	// 设定日志文件的定时清理
	m_cycle.New(m_cycle.Opt{
		Func:      ClearLog,
		SleepTime: time.Hour * 24,
	}).Start()

	Log.Println(
		`系统初始化完成`,
		m_json.Format(Path),
	)
}
