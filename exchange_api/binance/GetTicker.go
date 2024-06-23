package binance

import (
	"coin-candle/global"
	"fmt"

	"github.com/handy-golang/go-tools/m_fetch"
	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_json"
	jsoniter "github.com/json-iterator/go"
)

type BinanceTickerType []struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	AskPrice           string `json:"askPrice"`
	AskQty             string `json:"askQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstID            int    `json:"firstId"`
	LastID             int    `json:"lastId"`
	Count              int    `json:"count"`
}

func GetTicker() (resData []BinanceTickerType, resErr error) {

	resData = nil
	resErr = nil

	fetchData, err := m_fetch.NewHttp(m_fetch.HttpOpt{
		Origin:    BaseUrl,
		Path:      "/api/v3/ticker/24hr",
		ProxyURLs: global.Path.ProxyURLs,
	}).Get()
	if err != nil {
		resErr = err
		return
	}

	var result []BinanceTickerType
	jsoniter.Unmarshal(fetchData, &result)

	if len(result) < 2 {
		resErr = fmt.Errorf("错误:结果返回不正确 %+v", m_json.ToStr(result))
		return
	}
	resData = result

	//将请求结果写入目录
	m_file.WriteByte(global.Path.Binance.Dir+"/ticker-original-11.json", m_json.ToJson(resData))
	return

}
