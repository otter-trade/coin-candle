package okx

import (
	"coin-candle/global"
	"fmt"

	"github.com/handy-golang/go-tools/m_fetch"
	"github.com/handy-golang/go-tools/m_json"
	jsoniter "github.com/json-iterator/go"
)

type OkxTickerReqType struct {
	Code string                 `json:"code"`
	Msg  string                 `json:"msg"`
	Data []global.OkxTickerType `json:"data"`
}

func GetTicker() (resData []global.OkxTickerType, resErr error) {
	resData = nil
	resErr = nil

	fetchData, err := m_fetch.NewHttp(m_fetch.HttpOpt{
		Origin: BaseUrl,
		Path:   "/api/v5/market/tickers",
		DataMap: map[string]any{
			"instType": "SPOT",
		},
		ProxyURLs: global.Path.ProxyURLs,
	}).Get()

	if err != nil {
		resErr = err
		return
	}

	var result OkxTickerReqType
	jsoniter.Unmarshal(fetchData, &result)
	if result.Code != "0" {
		resErr = fmt.Errorf("错误:结果返回不正确 %+v", m_json.ToStr(fetchData))
		return
	}

	if len(result.Data) < 2 {
		resErr = fmt.Errorf("错误:数据长度不正确: %+v", result)
		return
	}

	resData = result.Data

	return
}
