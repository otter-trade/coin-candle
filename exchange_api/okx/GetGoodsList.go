package okx

import (
	"coin-candle/global"
	"fmt"

	"github.com/handy-golang/go-tools/m_fetch"
	"github.com/handy-golang/go-tools/m_json"
	jsoniter "github.com/json-iterator/go"
)

func GetGoodsList() {
	GetGoodsList_SPOT()
	GetGoodsList_SWAP()
}

type OkxInstType struct {
	Alias        string `json:"alias"`
	BaseCcy      string `json:"baseCcy"`
	Category     string `json:"category"`
	CtMult       string `json:"ctMult"`
	CtType       string `json:"ctType"`
	CtVal        string `json:"ctVal"`
	CtValCcy     string `json:"ctValCcy"`
	ExpTime      string `json:"expTime"`
	InstFamily   string `json:"instFamily"`
	InstID       string `json:"instId"`
	InstType     string `json:"instType"`
	Lever        string `json:"lever"`
	ListTime     string `json:"listTime"`
	LotSz        string `json:"lotSz"`
	MaxIcebergSz string `json:"maxIcebergSz"`
	MaxLmtAmt    string `json:"maxLmtAmt"`
	MaxLmtSz     string `json:"maxLmtSz"`
	MaxMktAmt    string `json:"maxMktAmt"`
	MaxMktSz     string `json:"maxMktSz"`
	MaxStopSz    string `json:"maxStopSz"`
	MaxTriggerSz string `json:"maxTriggerSz"`
	MaxTwapSz    string `json:"maxTwapSz"`
	MinSz        string `json:"minSz"`
	OptType      string `json:"optType"`
	QuoteCcy     string `json:"quoteCcy"`
	SettleCcy    string `json:"settleCcy"`
	State        string `json:"state"`
	Stk          string `json:"stk"`
	TickSz       string `json:"tickSz"`
	Uly          string `json:"uly"`
}

type OkxInstrumentsType struct {
	Code string        `json:"code"`
	Data []OkxInstType `json:"data"`
	Msg  string        `json:"msg"`
}

func GetGoodsList_SPOT() (resData []OkxInstType, resErr error) {

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
func GetGoodsList_SWAP() (resData []OkxInstType, resErr error) {
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
	//将请求结果写入目录
	return
}
