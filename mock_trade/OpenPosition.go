package mock_trade

import "fmt"

// 开仓

func (obj *MockActionObj) OpenPosition() (resErr error) {
	resErr = nil

	// 检查仓位是否存在
	if len(obj.NewPosition) < 1 {
		resErr = fmt.Errorf("NewPosition 为空")
		return
	}

	for k, v := range obj.NewPosition {
		// 这里要开始下单了
		fmt.Println(k, v.GoodsId, v.Amount)
	}

	return
}
