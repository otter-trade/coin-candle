package mock_trade

// 读取当前时间点的持仓情况
type ReadPositionOpt struct {
	StrategyID     string
	MockName       string
	Timestamp      int64 // 读取任意时间点的持仓情况(13位毫秒时间戳)，0 或 空 则为当前时间。
	ReportGenerate bool  // 是否生成报告，缺省值 false
}

func ReadPosition(opt ReadPositionOpt) {
	// 读取当前时间持仓情况，包括收益，策略情况评估等。
}
