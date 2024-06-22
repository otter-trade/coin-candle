package global

type Opt struct {
	LogPath  string // 日志文件存放目录，缺省值：./logs
	DataPath string // 数据文件存放目录，缺省值：./data
}

func Start(opt Opt) {
	// 初始化项 运行所需要的各种 目录
	DirInit(opt)

	// 初始化日志系统
	// mCycle.New(mCycle.Opt{
	// 	Func:      LogInit,
	// 	SleepTime: time.Hour * 24,
	// }).Start()

	// // 加载 SysEnv
	// config.ServerEnvInit()

	// Log.Println(
	// 	`系统初始化完成`,
	// 	mJson.Format(config.Dir),
	// 	mJson.Format(config.SysEnv),
	// 	mJson.Format(config.AppInfo),
	// )
}
