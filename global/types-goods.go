package global

var BinanceBaseUrls = []string{
	"https://api.binance.com",         // 0
	"https://api1.binance.com",        // 1
	"https://api2.binance.com",        // 2
	"https://api3.binance.com",        // 3
	"https://api4.binance.com",        // 4
	"https://api-gcp.binance.com",     // 5
	"https://data-api.binance.vision", // 6 仅限于公开数据
}

var OkxBaseUrls = []string{
	"https://www.okx.com",
	"https://aws.okx.com",
}

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
	State        string `json:"state"`        // 产品状态 live ：交易中 ，其余状态将被过滤
	TickSz       string `json:"tickSz"`       // 下单价格精度 如 0.0001
	Uly          string `json:"uly"`          // 标的指数，如 BTC-USD
}

// 详情请参考 https://binance-docs.github.io/apidocs/spot/cn/#3f1907847c
// 币安的文档很屎 ， 什么注解都木有 ，我这里只选取有用的信息进行注解
type BinanceSymbolType struct {
	Symbol                     string   `json:"symbol"`             // 交易品Id  BTCUSDT
	Status                     string   `json:"status"`             // status  TRADING
	BaseAsset                  string   `json:"baseAsset"`          // 基础货币如 BTC
	BaseAssetPrecision         int      `json:"baseAssetPrecision"` //资产精度
	QuoteAsset                 string   `json:"quoteAsset"`         // 交易货币 USDT
	QuotePrecision             int      `json:"quotePrecision"`
	QuoteAssetPrecision        int      `json:"quoteAssetPrecision"`
	BaseCommissionPrecision    int      `json:"baseCommissionPrecision"`
	QuoteCommissionPrecision   int      `json:"quoteCommissionPrecision"`
	OrderTypes                 []string `json:"orderTypes"`     // 支持的订单类型
	IcebergAllowed             bool     `json:"icebergAllowed"` // 是否允许冰山下单
	OcoAllowed                 bool     `json:"ocoAllowed"`
	OtoAllowed                 bool     `json:"otoAllowed"`
	QuoteOrderQtyMarketAllowed bool     `json:"quoteOrderQtyMarketAllowed"`
	AllowTrailingStop          bool     `json:"allowTrailingStop"`
	CancelReplaceAllowed       bool     `json:"cancelReplaceAllowed"`
	IsSpotTradingAllowed       bool     `json:"isSpotTradingAllowed"` // 是否允许现货交易
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

// 交易品的类型定义，融合 okx 和 币安
type GoodsType struct {
	GoodsId       string            `json:"GoodsId"`       // OtterTrade 的交易品ID 以 OKX 为准 如 BTC-USDT
	State         string            `json:"State"`         // 交易品现货状态，默认；live OKX 现货，币安 现货 ，OKX 合约 有一家状态不对 则为 warning
	UpdateUnix    int64             `json:"UpdateUnix"`    // 更新时间戳
	UpdateStr     string            `json:"UpdateStr"`     // 更新时间
	QuoteCcy      string            `json:"QuoteCcy"`      // 计价货币 如 USDT
	BaseCcy       string            `json:"BaseCcy"`       // 基础货币 如 BTC
	Okx_SPOT_Info OkxInstType       `json:"Okx_SPOT_Info"` // 欧意交易所 现货 的完整产品信息
	Okx_SWAP_Info OkxInstType       `json:"Okx_SWAP_Info"` // 欧意交易所 合约 的完整产品信息
	BinanceInfo   BinanceSymbolType `json:"BinanceInfo"`   // 币安交易所的完整产品信息
}

// 设定系统的结算货币
const SystemSettleCcy = "USDT"
