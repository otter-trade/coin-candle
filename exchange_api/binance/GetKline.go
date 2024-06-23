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

	// resData[0] 最新  resData[last] 最旧
	m_file.Write(global.Path.Binance.Dir+"/kline.json", m_json.JsonFormat(resData))

	return
}
