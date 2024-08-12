package mock_trade

/*

仓位管理 Position Manage
简化持仓模型。
每次更新当前需要的持仓的状态，然后系统会帮助下单并计算应得收入。
更新仓位状态
这种方式的好处在于，每次更新持仓时都可以动态的调整杠杆倍率，持仓比率等

*/

import (
	"fmt"
	"os"

	"github.com/handy-golang/go-tools/m_count"
	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_json"
	"github.com/handy-golang/go-tools/m_str"
	jsoniter "github.com/json-iterator/go"
	"github.com/otter-trade/coin-candle/global"
)

/*

更新一次仓位状态
这个函数要做的事情：
检查参数
平掉上一次的仓位
读取余额


*/

func UpdatePosition(opt global.UpdatePositionOpt) (resWarn []string, resErr error) {
	resErr = nil
	resWarn = []string{}
	// 要在这里读取上一次的结余

	var NewPositionList []global.NewPositionType
	// 持仓数据的过滤和检查
	for _, item := range opt.NewPosition {
		if len(item.GoodsId) > 1 {
			position, err := NewPositionParam(item)
			if err != nil {
				resErr = fmt.Errorf("%+v,%+v", item.GoodsId, err) // 只要有一个持仓有问题，则该次持仓判定为失效
				return
			}

			if position.GoodsDetail.State != "live" {
				resWarn = append(resWarn, fmt.Sprintf("产品%+v 状态不正确:%+v; 该持仓已从本次下单中剔除。", item.GoodsId, position.GoodsDetail.State))
			}

			// 下单金额大于 0 且 币种状态 live 才有效
			if m_count.Le(position.Amount, "0") > 0 && position.GoodsDetail.State != "live" {
				effPosition := global.NewPositionType{
					GoodsId:   position.GoodsDetail.GoodsId,
					TradeType: position.TradeType,
					TradeMode: position.TradeMode,
					Leverage:  position.Leverage,
					Side:      position.Side,
					Amount:    position.Amount,
				}
				NewPositionList = append(NewPositionList, effPosition)
			}
		}
	}

	if len(NewPositionList) == 0 {
		resWarn = append(resWarn, "本次持仓列表为空，代表系统会平掉所有持仓。")
	}

	Action, err := NewMockAction(global.NewMockActionOpt{
		StrategyID: opt.StrategyID,
		MockName:   opt.MockName,
		Time:       opt.UpdateTime,
	})
	if err != nil {
		resErr = err
		return
	}

	// 更新索引
	PositionIndexByte := m_file.ReadFile(Action.MockPath.PositionIndexFullPath)
	var PositionIndex global.PositionIndexType
	err = jsoniter.Unmarshal(PositionIndexByte, &PositionIndex)
	if err != nil {
		resErr = err
		return
	}
	PositionIndex = append(PositionIndex, Action.UpdateTime)

	// 更新持仓时间
	Action.MockConfig.LastPositionUpdateTime = Action.UpdateTime

	NewPositionListJsonPath := m_str.Join(
		Action.MockPath.MockDataFullDir,
		os.PathSeparator,
		Action.UpdateTime, ".json",
	)

	UpdatePositionInfo := global.UpdatePositionOpt{
		StrategyID:  Action.MockConfig.StrategyID,
		MockName:    Action.MockConfig.MockName,
		UpdateTime:  Action.UpdateTime,
		NewPosition: NewPositionList,
	}

	m_file.WriteByte(NewPositionListJsonPath, m_json.ToJson(UpdatePositionInfo))
	m_file.WriteByte(Action.MockPath.ConfigFullPath, m_json.ToJson(Action.MockConfig))
	m_file.WriteByte(Action.MockPath.PositionIndexFullPath, m_json.ToJson(PositionIndex))

	return
}
