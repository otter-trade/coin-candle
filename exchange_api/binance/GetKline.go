package binance

import (
	"coin-candle/global"
	"fmt"

	"github.com/handy-golang/go-tools/m_fetch"
	"github.com/handy-golang/go-tools/m_json"
	"github.com/handy-golang/go-tools/m_str"
	"github.com/handy-golang/go-tools/m_time"
	jsoniter "github.com/json-iterator/go"
)

type GetKlineOpt struct {
	Binance_symbol string `json:"Binance_symbol"`
	Bar            string `json:"Bar"`
	EndTime        int64  `json:"EndTime"` // 毫秒
}

/*
time := m_time.TimeParse(m_time.LaySP_ss, "2018-01-11 22:00:00")

	binanceKline, err := binance.GetKline(binance.GetKlineOpt{
		Binance_symbol: "BTCUSDT",
		Bar:            "1m",
		EndTime:        time,
	})
*/
type BinanceKlineType [12]any

func GetKline(opt GetKlineOpt) (resData []global.KlineSimpType, resErr error) {
	resData = nil
	resErr = nil

	if len(opt.Binance_symbol) < 2 {
		resErr = fmt.Errorf("参数 Binance_symbol 不正确")
		return
	}

	BarObj := global.GetBarOpt(opt.Bar)
	if BarObj.Interval < m_time.UnixTimeInt64.Minute {
		resErr = fmt.Errorf("参数 Bar 错误")
		return
	}

	// limit 固定为 global.ExchangeKlineLimit
	limit := global.ExchangeKlineLimit

	// 当前时间
	now := m_time.GetUnixInt64()
	EndTime := now
	// 时间 传入的时间戳 必须大于最早时间才有效否则重置为当 now
	if opt.EndTime > global.TimeOldest {
		EndTime = opt.EndTime
	}

	var DataMap = map[string]any{
		"symbol":   opt.Binance_symbol,
		"interval": BarObj.Binance,
		"endTime":  m_str.ToStr(EndTime + global.SendEndTimeFix), // 需要修正请求时间戳
		"limit":    limit,
	}

	fetchData, err := m_fetch.NewHttp(m_fetch.HttpOpt{
		Origin:    BaseUrl,
		Path:      "/api/v3/klines",
		DataMap:   DataMap,
		ProxyURLs: global.Path.ProxyURLs,
	}).Get()

	if err != nil {
		resErr = err
		return
	}

	// 时间 小 -> 大
	/*
		[
			1687559940000,   开始时间 0
			"30608.01000000",  开盘价 1
			"30628.51000000",  最高价 2
			"30606.36000000",  最低价 3
			"30622.30000000",  收盘价 4
			"30.34908000",  成交量 5
			1687559999999,   k线收盘时间   6
			"929230.51663220",  成交额 7
			870,  成交笔数 8
			"15.72761000",  主动买入成交量 9
			"481499.12090220",   主动买入成交额 10
			"0"
		],
	*/

	var listStr []BinanceKlineType
	jsoniter.Unmarshal(fetchData, &listStr)

	if len(listStr) != limit {
		resErr = fmt.Errorf("错误:K线长度不正确: %+v", m_json.ToStr(fetchData))
		return
	}

	// var Kline = []global.KlineType{}
	var KlineSimp = []global.KlineSimpType{}
	for _, item := range listStr {
		time_unix_str := m_json.ToStr(item[0])
		// var time = m_time.TimeGet(time_unix_str)

		// kItem := global.KlineType{
		// 	GoodsId:  opt.Binance_symbol,
		// 	TimeUnix: time.TimeUnix,
		// 	TimeStr:  time.TimeStr,
		// 	O:        m_str.ToStr(item[1]),
		// 	H:        m_str.ToStr(item[2]),
		// 	L:        m_str.ToStr(item[3]),
		// 	C:        m_str.ToStr(item[4]),
		// 	V:        m_str.ToStr(item[5]),
		// 	Q:        m_str.ToStr(item[7]),
		// }
		// Kline = append(Kline, kItem)

		KlineSimp = append(KlineSimp, global.KlineSimpType{
			time_unix_str,
			m_str.ToStr(item[1]),
			m_str.ToStr(item[2]),
			m_str.ToStr(item[3]),
			m_str.ToStr(item[4]),
			m_str.ToStr(item[5]),
			m_str.ToStr(item[7]),
		})
	}

	resData = KlineSimp
	// m_file.Write(global.Path.Binance.Dir+"/kline-Format.json", m_json.ToStr(Kline))
	// m_file.WriteByte(global.Path.Binance.Dir+"/kline-Simp-byte.json", m_json.ToJson(KlineSimp))
	return
}
