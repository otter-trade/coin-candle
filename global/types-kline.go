package global

// 系统支持的 Bar
var KlineBarOpt = []string{
	"1m",
	"5m",
	"15m",
	"30m",
	"1h",
}

// 系统支持的 交易所
var ExchangeOpt = []string{
	"binance", // 币安
	"okx",     // 欧意
}

type GetKlineOpt struct {
	GoodsId  string   `json:"GoodsId"`  // 商品ID
	Before   int64    `json:"Before"`   // 此时间之前的内容; 允许值: 13 位毫秒时间戳
	Limit    int      `json:"Limit"`    // 获取数据的总条目; 允许值: 1-500
	Bar      string   `json:"Bar"`      // K 线之间的间隔; 允许值: type KlineBarOpt
	Exchange []string `json:"Exchange"` // 交易所名称列表; 允许值: type ExchangeOpt
}
