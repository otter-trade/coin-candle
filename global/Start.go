package global

import (
	"time"

	"github.com/handy-golang/go-tools/m_cycle"
	"github.com/handy-golang/go-tools/m_json"
)

type Opt struct {
	LogPath  string // 日志文件存放目录，缺省值：./logs
	DataPath string // 数据文件存放目录，缺省值：./data
}

func Start(opt Opt) {
	// 初始化项 运行所需要的各种 目录
	DirInit(opt)

	// 初始化日志系统
	LogInit()
	// 设定日志文件的定时清理
	m_cycle.New(m_cycle.Opt{
		Func:      ClearLog,
		SleepTime: time.Hour * 24,
	}).Start()

	Log.Println(
		`系统初始化完成`,
		m_json.Format(Dir),
	)
}
