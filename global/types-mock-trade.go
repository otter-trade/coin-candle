package global

var MockNamePattern = "^[a-zA-Z0-9_\u4e00-\u9fa5]{1,12}$"

// map 类型
type MapAny map[string]any

type CreatePositionOpt struct {
	StrategyID   string // 策略的Id，从 OtterTrade 用户数据中读取，不可为空
	MockName     string // 模拟交易的名称，策略开发者自定义，不可为空，1-12位中文，字母或数字。
	RunType      string // 虚拟持仓的运行模式，缺省值 1 ，可选值 RunTypeList
	InitialAsset string // 初始资产(USDT)  缺省值 10000
	FeeRate      string // 手续费率 缺省值 0.001 参考 https://www.okx.com/zh-hans/fees
}

// 运行模式可选参数
var RunTypeList = MapAny{
	"1": MapAny{
		"description": "回测模式，交易只会在本地产生订单，方便查询和以及辅助策略的修改。",
	},
	"2": map[string]any{
		"description": "社区模式，交易不仅会在本地产生订单，还会同时将订单结果同步至交易所完成下单。需要OtterTrade平台的权限。",
	},
	"3": map[string]any{
		"description": "生产模式，交易不仅会在本地产生订单，还会同时将订单结果同步至交易所完成下单。策略需要通过 OtterTrade 平台的审核。",
	},
}
