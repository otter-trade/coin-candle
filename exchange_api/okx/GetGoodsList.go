package okx

import (
	"coin-candle/global"
	"fmt"

	"github.com/handy-golang/go-tools/m_fetch"
	"github.com/handy-golang/go-tools/m_json"
	jsoniter "github.com/json-iterator/go"
)

type OkxInstrumentsType struct {
	Code string               `json:"code"`
	Data []global.OkxInstType `json:"data"`
	Msg  string               `json:"msg"`
}

func GetGoodsList_SPOT() (resData []global.OkxInstType, resErr error) {

	resData = nil
	resErr = nil

	fetchData, err := m_fetch.NewHttp(m_fetch.HttpOpt{
		Origin: BaseUrl,
		Path:   "/api/v5/public/instruments",
		DataMap: map[string]any{
			"instType": "SPOT",
		},
		ProxyURLs: global.Path.ProxyURLs,
	}).Get()

	if err != nil {
		resErr = err
		return
	}

	var result OkxInstrumentsType
	jsoniter.Unmarshal(fetchData, &result)

	if len(result.Data) < 2 {
		resErr = fmt.Errorf("错误:结果返回不正确 %+v", m_json.ToStr(result))
		return
	}

	resData = result.Data

	return
}
func GetGoodsList_SWAP() (resData []global.OkxInstType, resErr error) {
	resData = nil
	resErr = nil

	fetchData, err := m_fetch.NewHttp(m_fetch.HttpOpt{
		Origin: BaseUrl,
		Path:   "/api/v5/public/instruments",
		DataMap: map[string]any{
			"instType": "SWAP",
		},
		ProxyURLs: global.Path.ProxyURLs,
	}).Get()

	if err != nil {
		resErr = err
		return
	}

	var result OkxInstrumentsType
	jsoniter.Unmarshal(fetchData, &result)

	if len(result.Data) < 2 {
		resErr = fmt.Errorf("错误:结果返回不正确 %+v", m_json.ToStr(result))
		return
	}

	resData = result.Data
	//将请求结果写入目录
	return
}
