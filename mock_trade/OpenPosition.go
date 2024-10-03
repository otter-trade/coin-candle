package mock_trade

import (
	"fmt"

	"github.com/handy-golang/go-tools/m_time"
	"github.com/otter-trade/coin-candle/global"
)

// 开仓
type OpenPositionOpt struct {
	OrderTime int64
}

func (obj *MockActionObj) OpenPosition(opt OpenPositionOpt) (resErr error) {
	resErr = nil

	// 检查是否有仓位...
	if len(obj.NewPosition) < 1 {
		resErr = fmt.Errorf("NewPosition 为空")
		return
	}

	var OrderTime int64
	nowTime := m_time.GetUnixInt64()

	// 回测模式，则取 用户传入的时间
	if obj.MockServeConfig.RunMode.Value == 1 {
		OrderTime = opt.OrderTime
	}

	// 时间为空 则为当前时间
	if OrderTime == 0 {
		OrderTime = nowTime
	} else {
		//  小于系统最老时间 则等于系统最老时间
		if OrderTime < global.TimeOldest {
			OrderTime = global.TimeOldest
		}
	}
	// 或者 大于当前 则重置为当前时间
	if OrderTime > nowTime {
		OrderTime = nowTime
	}

	fmt.Println("下单时间", OrderTime)

	for k, v := range obj.NewPosition {
		// 这里要开始下单了
		fmt.Println(k, v.GoodsId, v.Amount)
	}

	return
}
