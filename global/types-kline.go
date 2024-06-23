package global

import (
	"strings"

	"github.com/handy-golang/go-tools/m_time"
)

// 系统支持的 交易所
var ExchangeOpt = []string{
	"binance", // 币安
	"okx",     // 欧意
}

// 系统支持的 Bar
type KlineBarType struct {
	Interval int64  // 每一条数据之间间隔的毫秒数
	Okx      string // Okx 的参数值
	Binance  string // Binance 的参数值
}

var KlineBarOpt = map[string]KlineBarType{
	"1m": {
		Interval: m_time.UnixTimeInt64.Minute * 1,
		Okx:      "1m",
		Binance:  "1m",
	},
	"5m": {
		Interval: m_time.UnixTimeInt64.Minute * 5,
		Okx:      "5m",
		Binance:  "5m",
	},
	"15m": {
		Interval: m_time.UnixTimeInt64.Minute * 15,
		Okx:      "15m",
		Binance:  "15m",
	},
	"30m": {
		Interval: m_time.UnixTimeInt64.Minute * 30,
		Okx:      "30m",
		Binance:  "30m",
	},
	"1h": {
		Interval: m_time.UnixTimeInt64.Hour * 1,
		Okx:      "1H",
		Binance:  "1h",
	},
}

func GetBarOpt(b string) (resData KlineBarType) {
	param := strings.ToLower(b)
	bar := KlineBarOpt[param]
	return bar
}

type GetKlineOpt struct {
	GoodsId  string   `json:"GoodsId"`  // 商品ID
	Bar      string   `json:"Bar"`      // K 线之间的间隔; 允许值: type KlineBarOpt
	Before   int64    `json:"Before"`   // 此时间之前的内容; 允许值: 13 位毫秒时间戳
	Limit    int      `json:"Limit"`    // 获取数据的总条目; 允许值: 1-500
	Exchange []string `json:"Exchange"` // 交易所名称列表; 允许值: type ExchangeOpt
}

type KlineType struct {
	GoodsId  string `json:"GoodsId"`  // 商品ID
	TimeUnix int64  `json:"TimeUnix"` // 开始时间
	TimeStr  string `json:"TimeStr"`  // 开始时间 字符串形式
	O        string `json:"O"`        // 开盘价
	H        string `json:"H"`        // 最高价
	L        string `json:"L"`        // 最低价
	C        string `json:"C"`        // 收盘价格
	V        string `json:"V"`        // 成交量(BTC数量)  V * C = Q
	Q        string `json:"Q"`        // 成交额(USDT数量)  Q / C = V
}

type KlineSimpType [7]string // TimeUnix,O,H,L,C,V,Q
