package mock_trade

type NewMockActionOpt struct {
	StrategyID string // 策略的Id
	MockName   string // 本次回测的名称
}

type MockActionObj struct {
	StrategyID string // 策略的Id
	MockName   string // 本次回测的名称
	Time       int64  // 本次回测的名称
}

func NewMockAction(opt NewMockActionOpt) (action *MockActionObj, resErr error) {
	action = nil
	resErr = nil
	obj := MockActionObj{}

	action = &obj
	return
}
