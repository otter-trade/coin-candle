package okx

import (
	"coin-candle/global"
	"fmt"

	"github.com/handy-golang/go-tools/m_fetch"
	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_json"
	"github.com/handy-golang/go-tools/m_str"
	"github.com/handy-golang/go-tools/m_time"
)

/*

	resData, err := okx.GetKline(global.GetOkxKlineOpt{
		Okx_instId: "BTC-USDT",
		Bar:        "1m",
		Before:     m_time.GetUnixInt64() - m_time.UnixTimeInt64.Day*1, // 一天前
	})
	fmt.Println(resData, err)

*/

type GetKlineOpt struct {
	Okx_instId string `json:"Okx_instId"`
	Bar        string `json:"Bar"`
	Before     int64  `json:"Before"`
}

func GetKline(opt GetKlineOpt) (resData []byte, resErr error) {

	resData = nil
	resErr = nil

	if len(opt.Okx_instId) < 2 {
		resErr = fmt.Errorf("参数 Okx_instId 不能为空")
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

	path := "/api/v5/market/candles"
	// 当前时间 - 之前的时间 / 时间间隔 = 距离当前的历史条目
	fromNowItem := (now - before) / BarObj.Interval
	if fromNowItem > 800 { // 大于 800 条就从历史接口提取数据
		path = "/api/v5/market/history-index-candles"
	}

	var DataMap = map[string]any{
		"instId": opt.Okx_instId,
		"bar":    BarObj.Okx,
		"after":  m_str.ToStr(before),
		"limit":  limit,
	}

	resData, err := m_fetch.NewHttp(m_fetch.HttpOpt{
		Origin:    BaseUrl,
		Path:      path,
		DataMap:   DataMap,
		ProxyURLs: global.Path.ProxyURLs,
	}).Get()

	if err != nil {
		global.LogErr("okx.GetKline Err", err, m_json.Format(DataMap))
	}

	// resData[0] 最新  resData[last] 最旧
	m_file.Write(global.Path.Okx.Dir+"/kline-spot.json", m_json.JsonFormat(resData))

	return
}
