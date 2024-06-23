package okx

import (
	"coin-candle/global"
	"fmt"

	"github.com/handy-golang/go-tools/m_fetch"
	"github.com/handy-golang/go-tools/m_json"
	jsoniter "github.com/json-iterator/go"
)

// 详情请参考  https://www.okx.com/docs-v5/zh/#public-data-rest-api-get-instruments
type OkxInstType struct {
	BaseCcy      string `json:"baseCcy"`      // 交易货币币种，如 BTC-USDT 中的 BTC ，仅适用于币币/币币杠杆
	CtMult       string `json:"ctMult"`       // 合约乘数
	CtType       string `json:"ctType"`       // 合约类型  linear：正向合约  inverse：反向合约
	CtVal        string `json:"ctVal"`        // 合约面值 如: 一个合约 0.01 个 BTC
	CtValCcy     string `json:"ctValCcy"`     // 合约面值计价币种 如 BTC
	ExpTime      string `json:"expTime"`      // 下架时间
	InstFamily   string `json:"instFamily"`   // 交易品种，如 BTC-USD，仅适用于杠杆/交割/永续/期权
	InstID       string `json:"instId"`       // BTC-USDT   BTC-USDT-SWAP
	InstType     string `json:"instType"`     // 产品类型  SPOT 或者 SWAP
	Lever        string `json:"lever"`        // 支持的最大杠杆倍率
	ListTime     string `json:"listTime"`     // 上架时间
	LotSz        string `json:"lotSz"`        // 最小下单数量 合约为 张数，现货则为 交易品数量 如 BTC 数量
	MaxIcebergSz string `json:"maxIcebergSz"` // 冰山委托的单笔最大委托数量
	MaxLmtSz     string `json:"maxLmtSz"`     // 限价单最大委托数量  合约的数量单位是张，现货的数量单位是USDT
	MaxMktSz     string `json:"maxMktSz"`     // 市价单的单笔最大委托数量  合约的数量单位是张，现货的数量单位是USDT
	MaxStopSz    string `json:"maxStopSz"`    // 止盈止损市价委托的单笔最大委托数量
	MaxTriggerSz string `json:"maxTriggerSz"` // 计划委托委托的单笔最大委托数量
	MaxTwapSz    string `json:"maxTwapSz"`    // 时间加权单的单笔最大委托数量
	MinSz        string `json:"minSz"`        // 最小下单数量 合约的数量单位是张，现货的数量单位是交易货币
	QuoteCcy     string `json:"quoteCcy"`     // 计价货币币种，如 BTC-USDT 中的USDT ，仅适用于币币交易
	SettleCcy    string `json:"settleCcy"`    // 结算币种 ，如 USDT
	State        string `json:"state"`        // 产品状态 live：交易中 ，其余状态将被过滤
	TickSz       string `json:"tickSz"`       // 下单价格精度 如 0.0001
	Uly          string `json:"uly"`          // 标的指数，如 BTC-USD
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
