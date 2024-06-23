package binance

import (
	"coin-candle/global"
	"fmt"

	"github.com/handy-golang/go-tools/m_fetch"
	"github.com/handy-golang/go-tools/m_json"
	jsoniter "github.com/json-iterator/go"
)

type SymbolType struct {
	Symbol                     string   `json:"symbol"`
	Status                     string   `json:"status"`
	BaseAsset                  string   `json:"baseAsset"`
	BaseAssetPrecision         int      `json:"baseAssetPrecision"`
	QuoteAsset                 string   `json:"quoteAsset"`
	QuotePrecision             int      `json:"quotePrecision"`
	QuoteAssetPrecision        int      `json:"quoteAssetPrecision"`
	BaseCommissionPrecision    int      `json:"baseCommissionPrecision"`
	QuoteCommissionPrecision   int      `json:"quoteCommissionPrecision"`
	OrderTypes                 []string `json:"orderTypes"`
	IcebergAllowed             bool     `json:"icebergAllowed"`
	OcoAllowed                 bool     `json:"ocoAllowed"`
	OtoAllowed                 bool     `json:"otoAllowed"`
	QuoteOrderQtyMarketAllowed bool     `json:"quoteOrderQtyMarketAllowed"`
	AllowTrailingStop          bool     `json:"allowTrailingStop"`
	CancelReplaceAllowed       bool     `json:"cancelReplaceAllowed"`
	IsSpotTradingAllowed       bool     `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed     bool     `json:"isMarginTradingAllowed"`
	Filters                    []struct {
		FilterType            string `json:"filterType"`
		MinPrice              string `json:"minPrice,omitempty"`
		MaxPrice              string `json:"maxPrice,omitempty"`
		TickSize              string `json:"tickSize,omitempty"`
		MinQty                string `json:"minQty,omitempty"`
		MaxQty                string `json:"maxQty,omitempty"`
		StepSize              string `json:"stepSize,omitempty"`
		Limit                 int    `json:"limit,omitempty"`
		MinTrailingAboveDelta int    `json:"minTrailingAboveDelta,omitempty"`
		MaxTrailingAboveDelta int    `json:"maxTrailingAboveDelta,omitempty"`
		MinTrailingBelowDelta int    `json:"minTrailingBelowDelta,omitempty"`
		MaxTrailingBelowDelta int    `json:"maxTrailingBelowDelta,omitempty"`
		BidMultiplierUp       string `json:"bidMultiplierUp,omitempty"`
		BidMultiplierDown     string `json:"bidMultiplierDown,omitempty"`
		AskMultiplierUp       string `json:"askMultiplierUp,omitempty"`
		AskMultiplierDown     string `json:"askMultiplierDown,omitempty"`
		AvgPriceMins          int    `json:"avgPriceMins,omitempty"`
		MinNotional           string `json:"minNotional,omitempty"`
		ApplyMinToMarket      bool   `json:"applyMinToMarket,omitempty"`
		MaxNotional           string `json:"maxNotional,omitempty"`
		ApplyMaxToMarket      bool   `json:"applyMaxToMarket,omitempty"`
		MaxNumOrders          int    `json:"maxNumOrders,omitempty"`
		MaxNumAlgoOrders      int    `json:"maxNumAlgoOrders,omitempty"`
	} `json:"filters"`
	Permissions                     []interface{} `json:"permissions"`
	PermissionSets                  [][]string    `json:"permissionSets"`
	DefaultSelfTradePreventionMode  string        `json:"defaultSelfTradePreventionMode"`
	AllowedSelfTradePreventionModes []string      `json:"allowedSelfTradePreventionModes"`
}

type BinanceExchangeInfoType struct {
	Timezone   string `json:"timezone"`
	ServerTime int64  `json:"serverTime"`
	RateLimits []struct {
		RateLimitType string `json:"rateLimitType"`
		Interval      string `json:"interval"`
		IntervalNum   int    `json:"intervalNum"`
		Limit         int    `json:"limit"`
	} `json:"rateLimits"`
	ExchangeFilters []interface{} `json:"exchangeFilters"`
	Symbols         []SymbolType  `json:"symbols"`
}

func GetGoodsList() (resData []SymbolType, resErr error) {

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
