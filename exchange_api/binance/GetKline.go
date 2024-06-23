package binance

import (
	"coin-candle/global"
	"fmt"

	"github.com/handy-golang/go-tools/m_fetch"
	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_json"
	"github.com/handy-golang/go-tools/m_str"
	"github.com/handy-golang/go-tools/m_time"
	jsoniter "github.com/json-iterator/go"
)

type GetKlineOpt struct {
	Binance_symbol string `json:"Binance_symbol"`
	Bar            string `json:"Bar"`
	Before         int64  `json:"Before"`
}

/*
	resData, err  := binance.GetKline(binance.GetKlineOpt{
		Binance_symbol: "BTCUSDT",
		Bar:            "1m",
		Before:         m_time.GetUnixInt64() - m_time.UnixTimeInt64.Day*365, // 一年前
	})
*/
type BinanceKlineType [12]any

func GetKline(opt GetKlineOpt) (resData []byte, resErr error) {
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

	// limit 固定为 100
	limit := 100
	// 当前时间
	now := m_time.GetUnixInt64()
	before := now
	// 时间 传入的时间戳 必须大于6年前 才有效
	if opt.Before > now-m_time.UnixTimeInt64.Day*2190 {
		before = opt.Before
	}

	var DataMap = map[string]any{
		"symbol":   opt.Binance_symbol,
		"interval": BarObj.Binance,
		"endTime":  m_str.ToStr(before),
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

	if len(listStr) < 1 {
		resErr = fmt.Errorf("错误:K线长度不正确: %+v", len(listStr))
		return
	}

	var Kline = []global.KlineType{}
	for _, item := range listStr {
		time_str := m_json.ToStr(item[0])
		var time = m_time.TimeGet(time_str)

		kItem := global.KlineType{
			GoodsId:  opt.Binance_symbol,
			TimeUnix: time.TimeUnix,
			TimeStr:  time.TimeStr,
			O:        m_str.ToStr(item[1]),
			H:        m_str.ToStr(item[2]),
			L:        m_str.ToStr(item[3]),
			C:        m_str.ToStr(item[4]),
			V:        m_str.ToStr(item[5]),
			Q:        m_str.ToStr(item[7]),
		}
		Kline = append(Kline, kItem)
	}
	m_file.Write(global.Path.Binance.Dir+"/kline-Format.json", m_json.Format(Kline))
	return
}
