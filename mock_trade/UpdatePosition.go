package mock_trade

import (
	"fmt"
	"os"

	"github.com/handy-golang/go-tools/m_count"
	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_json"
	"github.com/handy-golang/go-tools/m_str"
	"github.com/handy-golang/go-tools/m_time"
	jsoniter "github.com/json-iterator/go"
	"github.com/otter-trade/coin-candle/exchange_api"
	"github.com/otter-trade/coin-candle/global"
)

// 更新一次仓位状态
// 这一次仓位的更新，相当于上一次仓位的全部平仓。这个记录本质上是一个指令记录。
func UpdatePosition(opt global.UpdatePositionOpt) (resErr error) {
	resErr = nil

	// 读取 Config 信息
	MockConfig, err := GetMockServeInfo(global.FindMockServeOpt{
		StrategyID: opt.StrategyID,
		MockName:   opt.MockName,
	})
	if err != nil {
		resErr = err
		return
	}

	// 获取 UpdateTime 持仓的更新时间
	nowTime := m_time.GetUnixInt64()
	var UpdateTime int64
	// 只有 回测模式， UpdateTime 才有效
	if MockConfig.RunMode.Value == 1 {
		UpdateTime = opt.UpdateTime
	}
	// 小于系统最老时间 或者 大于当前 则重置为当前时间
	if UpdateTime < global.TimeOldest || UpdateTime > nowTime {
		UpdateTime = nowTime
	}

	// 你不能跳回到过去下单， UpdateTime 必须大于上一次更新时间
	lastUpdateTime := MockConfig.LastPositionUpdateTime
	if UpdateTime-lastUpdateTime < m_time.UnixTimeInt64.Minute {
		resErr = fmt.Errorf("UpdateTime必须大于上一次更新时间")
		return
	}

	var NewPositionList []global.NewPositionType
	// 参数过滤和检错
	for _, item := range opt.NewPosition {
		if len(item.GoodsId) > 1 {
			// 必须为有效的 GoodsId
			GoodsDetail, err := exchange_api.GetGoodsDetail(exchange_api.GetGoodsDetailOpt{
				GoodsId: item.GoodsId,
			})
			if err != nil {
				resErr = err
				return
			}
			// 参数检查
			position, err := NewPositionFuncParamCheck(item)
			if err != nil {
				resErr = fmt.Errorf("%+v,%+v", item.GoodsId, err) // 只要有一个持仓有问题，则该次持仓判定为失效
				return
			}
			// 下单金额大于 0 才有效 , 币种状态 live 才有效
			if m_count.Le(position.Amount, "0") > 0 && GoodsDetail.State == "live" {
				NewPositionList = append(NewPositionList, position)
			}
		}
	}

	mockPath, _ := CheckMockName(global.FindMockServeOpt{
		StrategyID: MockConfig.StrategyID,
		MockName:   MockConfig.MockName,
	})

	MockConfig.LastPositionUpdateTime = UpdateTime

	PositionIndexByte := m_file.ReadFile(mockPath.PositionIndexFullPath)
	var PositionIndex global.PositionIndexType
	jsoniter.Unmarshal(PositionIndexByte, &PositionIndex)
	PositionIndex = append(PositionIndex, UpdateTime)

	NewPositionListJsonPath := m_str.Join(
		mockPath.MockDataFullDir,
		os.PathSeparator,
		UpdateTime, ".json",
	)

	UpdatePositionInfo := global.UpdatePositionOpt{
		StrategyID:  MockConfig.StrategyID,
		MockName:    MockConfig.MockName,
		UpdateTime:  UpdateTime,
		NewPosition: NewPositionList,
	}

	m_file.WriteByte(NewPositionListJsonPath, m_json.ToJson(UpdatePositionInfo))
	m_file.WriteByte(mockPath.ConfigFullPath, m_json.ToJson(MockConfig))
	m_file.WriteByte(mockPath.PositionIndexFullPath, m_json.ToJson(PositionIndex))

	return
}

// 参数过滤与检查
func NewPositionFuncParamCheck(opt global.NewPositionType) (resData global.NewPositionType, resErr error) {
	resData = global.NewPositionType{}
	resErr = nil

	// 检查参数

	resData.GoodsId = opt.GoodsId

	if opt.Side == "Buy" || opt.Side == "Sell" {
		resData.Side = opt.Side
	} else {
		resErr = fmt.Errorf("Side不正确")
		return
	}

	_, err := global.GetTradeType(opt.TradeType)
	if err != nil {
		resErr = err
		return
	}
	resData.TradeType = opt.TradeType

	TradeModeValue := opt.TradeMode
	if len(TradeModeValue) > 1 {
		_, err := global.GetTradeMode(TradeModeValue)
		if err != nil {
			resErr = err
			return
		}
	} else {
		TradeModeValue = global.TradeModeList[0].Value
	}
	resData.TradeMode = TradeModeValue

	Leverage := "1"
	if resData.TradeMode == "SWAP" {
		Leverage = m_count.Sub(opt.Leverage, "0")
		Leverage = m_count.Cent(Leverage, 0)

		if m_count.Le(Leverage, "1") < 0 {
			Leverage = "1" // 最小值为 1
		}
		if m_count.Le(Leverage, global.MaxLeverage) > 0 {
			Leverage = global.MaxLeverage // 最大值
		}
	}
	resData.Leverage = Leverage

	// 买入金额
	Amount := m_count.Sub(opt.Amount, "0")
	if m_count.Le(Amount, "0") < 0 {
		Amount = "0" // 最小值为 0
	}
	resData.Amount = Amount

	return
}
