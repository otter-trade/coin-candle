package binance

import (
	"coin-candle/global"
	"fmt"

	"github.com/handy-golang/go-tools/m_fetch"
	"github.com/handy-golang/go-tools/m_json"
	jsoniter "github.com/json-iterator/go"
)

type BinanceExchangeInfoType struct {
	Timezone   string `json:"timezone"`
	ServerTime int64  `json:"serverTime"`
	RateLimits []struct {
		RateLimitType string `json:"rateLimitType"`
		Interval      string `json:"interval"`
		IntervalNum   int    `json:"intervalNum"`
		Limit         int    `json:"limit"`
	} `json:"rateLimits"`
	ExchangeFilters []interface{}              `json:"exchangeFilters"`
	Symbols         []global.BinanceSymbolType `json:"symbols"`
}

func GetGoodsList() (resData []global.BinanceSymbolType, resErr error) {

	resData = nil
	resErr = nil

	fetchData, err := m_fetch.NewHttp(m_fetch.HttpOpt{
		Origin:    BaseUrl,
		Path:      "/api/v3/exchangeInfo",
		ProxyURLs: global.Path.ProxyURLs,
	}).Get()

	if err != nil {
		resErr = err
		return
	}

	var result BinanceExchangeInfoType
	jsoniter.Unmarshal(fetchData, &result)

	if len(result.Symbols) < 2 {
		resErr = fmt.Errorf("错误:结果返回不正确 %+v", m_json.ToStr(result))
		return
	}

	resData = result.Symbols

	return
}
