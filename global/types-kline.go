package global

/*
所有和 Kline 相关的 type 定义都放在这里
*/

import (
	"strings"

	"github.com/handy-golang/go-tools/m_time"
)

// 系统支持的 Bar
type KlineBarType struct {
	Interval int64  // 每一条数据之间间隔的毫秒数
	DirName  string // 数据存储目录的值
	Okx      string // Okx 的参数值
	Binance  string // Binance 的参数值
}

var KlineBarOpt = map[string]KlineBarType{
	"1m": {
		Interval: m_time.UnixTimeInt64.Minute * 1,
		DirName:  "1m",
		Okx:      "1m",
		Binance:  "1m",
	},
	"5m": {
		Interval: m_time.UnixTimeInt64.Minute * 5,
		DirName:  "5m",
		Okx:      "5m",
		Binance:  "5m",
	},
	"15m": {
		Interval: m_time.UnixTimeInt64.Minute * 15,
		DirName:  "15m",
		Okx:      "15m",
		Binance:  "15m",
	},
	"30m": {
		Interval: m_time.UnixTimeInt64.Minute * 30,
		DirName:  "30m",
		Okx:      "30m",
		Binance:  "30m",
	},
	"1h": {
		Interval: m_time.UnixTimeInt64.Hour * 1,
		DirName:  "1h",
		Okx:      "1H",
		Binance:  "1h",
	},
}

func GetBarOpt(b string) (resData KlineBarType) {
	param := strings.ToLower(b)
	bar := KlineBarOpt[param]
	return bar
}

// 系统支持的 交易所
var ExchangeOpt = []string{
	"binance", // 0 币安 顺序不可打乱
	"okx",     // 1 欧意 顺序不可打乱
}

type GetKlineOpt struct {
	GoodsId  string   `json:"GoodsId"`  // 商品ID , 必传
	Bar      string   `json:"Bar"`      // K 线之间的间隔; 允许值: global.KlineBarOpt
	EndTime  int64    `json:"EndTime"`  // K 线的结束时间; 允许值: 13 位毫秒时间戳, 若时间无效，则为当前时间。
	Limit    int      `json:"Limit"`    // 获取数据的总条目; 允许值: 1~KlineMaxLimit 缺省值 KlineLimitDefault
	Exchange []string `json:"Exchange"` // 交易所名称列表; 允许值: global.ExchangeOpt , 缺省值 okx
}

type KlineType struct {
	GoodsId  string `json:"GoodsId"`  // 商品ID
	Exchange string `json:"Exchange"` // 交易所
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

type KlineExchangeMap map[string][]KlineSimpType

const SendEndTimeFix = 10 // 发出 请求时 K线 时间戳修订 10 毫秒，考虑到网络延迟以及交易所不同标准的修订问题，不影响实盘实时数据的正常值

// 如果小于  2018-01-11 22:00:00 这个时间，则交易所数据就不全了， 这里将合法的最小时间定为 2018-03-01
var TimeOldest = m_time.TimeParse(m_time.LaySP_ss, "2018-03-01 00:00:00")

const KlineMaxLimit = 500 // 请求 K线时 最大的 Limit

const KlineLimitDefault = 10 // 请求 K线时 缺省的 Limit

const ExchangeKlineLimit = 100 // 交易所拿取K线的最大值

// 在 20 秒内，不重复请求，意思就是说价格允许 20 秒的延迟。可修改为 0 但是请求交易所的频率会很高
var KlineRequestInterval = m_time.UnixTimeInt64.Seconds * 20

// 文件名计算基准时间
var FileNameBaseTime = m_time.TimeParse(m_time.LaySP_ss, "2024-05-20 00:00:00")
