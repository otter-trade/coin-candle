package global

/*
所有和 市场相关定义 相关的 type 定义都放在这里
*/

type BinanceTickerType struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`        // 价格变化
	PriceChangePercent string `json:"priceChangePercent"` // 价格变化百分比
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

type OkxTickerType struct {
	InstType  string `json:"instType"`
	InstID    string `json:"instId"`
	Last      string `json:"last"`
	LastSz    string `json:"lastSz"`
	AskPx     string `json:"askPx"`
	AskSz     string `json:"askSz"`
	BidPx     string `json:"bidPx"`
	BidSz     string `json:"bidSz"`
	Open24H   string `json:"open24h"`
	High24H   string `json:"high24h"`
	Low24H    string `json:"low24h"`
	VolCcy24H string `json:"volCcy24h"`
	Vol24H    string `json:"vol24h"`
	Ts        string `json:"ts"`
	SodUtc0   string `json:"sodUtc0"`
	SodUtc8   string `json:"sodUtc8"`
}

type TickerType struct {
	GoodsId        string `json:"GoodsId"`
	Okx_InstID     string `json:"Okx_InstID"`
	Binance_Symbol string `json:"Binance_Symbol"`
	BaseCcy        string `json:"BaseCcy"`        // 币种名称
	State          string `json:"State"`          // 交易品现货状态，默认；live OKX 现货，币安 现货 ，OKX 合约 有一家状态不对 则为 warning
	Last           string `json:"Last"`           // 最新成交价
	Open24H        string `json:"Open24H"`        // 24小时开盘价
	High24H        string `json:"High24H"`        // 最高价
	Low24H         string `json:"Low24H"`         // 最低价
	OKXVol24H      string `json:"OKXVol24H"`      // OKX 24小时成交量 USDT 数量
	BinanceVol24H  string `json:"BinanceVol24H"`  // Binance 24 小时成交 USDT 数量
	U_R24          string `json:"U_R24"`          // 涨幅 = (最新价-开盘价)/开盘价
	Volume         string `json:"Volume"`         // 成交量总和
	OkxVolRose     string `json:"OkxVolRose"`     // 欧意占比总交易量
	BinanceVolRose string `json:"BinanceVolRose"` // 币安占比总交易量
	TimeUnix       int64  `json:"TimeUnix"`
	TimeStr        string `json:"TimeStr"`
}
