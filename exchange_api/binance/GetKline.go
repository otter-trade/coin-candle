package binance

import (
	"coin-candle/global"
	"fmt"

	"github.com/handy-golang/go-tools/m_fetch"
	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_json"
	"github.com/handy-golang/go-tools/m_str"
	"github.com/handy-golang/go-tools/m_time"
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

	resData, err := m_fetch.NewHttp(m_fetch.HttpOpt{
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
	m_file.Write(global.Path.Binance.Dir+"/kline.json", m_json.JsonFormat(resData))

	/*
		[
			1687559940000,
			"30608.01000000",
			"30628.51000000",
			"30606.36000000",
			"30622.30000000",
			"30.34908000",
			1687559999999,
			"929230.51663220",
			870,
			"15.72761000",
			"481499.12090220",
			"0"
		],
	*/

	return
}
