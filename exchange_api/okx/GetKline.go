package okx

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
	Okx_instId string `json:"Okx_instId"`
	Bar        string `json:"Bar"`
	EndTime    int64  `json:"EndTime"` // 毫秒
}

/*
	time := m_time.TimeParse(m_time.LaySP_ss, "2018-01-11 22:00:00")
	okxKline, err := okx.GetKline(okx.GetKlineOpt{
		Okx_instId: "BTC-USDT",
		Bar:        "1m",
		EndTime:    time,
	})
*/

type KlineReqType struct {
	Code string         `json:"Code"`
	Data []OkxKlineType `json:"Data"`
	Msg  string         `json:"Msg"`
}

type OkxKlineType [9]string

func GetKline(opt GetKlineOpt) (resData []global.KlineSimpType, resErr error) {

	resData = nil
	resErr = nil

	if len(opt.Okx_instId) < 3 {
		resErr = fmt.Errorf("参数 Okx_instId 不正确")
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

	path := "/api/v5/market/candles"
	// 当前时间 - 之前的时间 / 时间间隔 = 距离当前的历史条目
	fromNowItem := (now - EndTime) / BarObj.Interval
	if fromNowItem > 300 { // 大于 300 条就从历史接口提取数据
		path = "/api/v5/market/history-candles"
	}

	var DataMap = map[string]any{
		"instId": opt.Okx_instId,
		"bar":    BarObj.Okx,
		"after":  m_str.ToStr(EndTime + global.SendEndTimeFix), // 需要修正请求时间戳
		"limit":  limit,
	}

	fetchData, err := m_fetch.NewHttp(m_fetch.HttpOpt{
		Origin:    BaseUrl,
		Path:      path,
		DataMap:   DataMap,
		ProxyURLs: global.Path.ProxyURLs,
	}).Get()

	if err != nil {
		resErr = err
		return
	}

	/*
		开始时间 从 大 -> 小
		[
			"1687565880000",  开始时间 0
			"30633.5",  开盘价格  1
			"30637.6",  最高价格  2
			"30620",  最低价格 3
			"30626",  收盘价格 4
			"3.97909485",  交易量 以张为单位 5
			"121880.586345833",  交易量，以币为单位 6
			"121880.586345833",  交易量，以计价货币为单位  单位均是USDT  7
			"1"
		],

	*/

	var result KlineReqType
	jsoniter.Unmarshal(fetchData, &result)
	if result.Code != "0" {
		resErr = fmt.Errorf("错误:结果返回不正确 %+v", m_json.ToStr(fetchData))
		return
	}

	if len(result.Data) != limit {
		resErr = fmt.Errorf("错误:K线长度不正确: %+v", m_json.ToStr(fetchData))
		return
	}

	// var Kline = []global.KlineType{}
	var KlineSimp = []global.KlineSimpType{}
	for i := len(result.Data) - 1; i >= 0; i-- {
		item := result.Data[i]
		// var time = m_time.TimeGet(item[0])

		// kItem := global.KlineType{
		// 	GoodsId:  opt.Okx_instId,
		// 	TimeUnix: time.TimeUnix,
		// 	TimeStr:  time.TimeStr,
		// 	O:        item[1],
		// 	H:        item[2],
		// 	L:        item[3],
		// 	C:        item[4],
		// 	V:        item[5],
		// 	Q:        item[7],
		// }
		// Kline = append(Kline, kItem)

		KlineSimp = append(KlineSimp, global.KlineSimpType{
			item[0],
			item[1],
			item[2],
			item[3],
			item[4],
			item[5],
			item[7],
		})
	}
	resData = KlineSimp
	// m_file.Write(global.Path.Okx.Dir+"/kline-Format.json", m_json.Format(Kline))
	// m_file.WriteByte(global.Path.Okx.Dir+"/kline-Simp-byte.json", m_json.ToJson(KlineSimp))
	return
}
