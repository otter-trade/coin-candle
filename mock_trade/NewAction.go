package mock_trade

import (
	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_json"
	"github.com/otter-trade/coin-candle/global"
)

type NewMockActionOpt struct {
	StrategyID string // 策略的Id
	MockName   string // 本次回测的名称
}

type MockActionObj struct {
	StrategyID      string // 策略的Id
	MockName        string // 本次回测的名称
	MockPath        MockPathType
	MockServeConfig global.MockServeConfigType
}

func NewMockAction(opt NewMockActionOpt) (action *MockActionObj, resErr error) {
	action = nil
	resErr = nil
	obj := MockActionObj{}

	obj.MockName = opt.MockName
	obj.StrategyID = opt.StrategyID

	// 检查 StrategyID 和 MockName 并获取存储目录
	mockPath, err := CheckMockName(global.FindMockServeOpt{
		StrategyID: obj.StrategyID,
		MockName:   obj.MockName,
	})
	if err != nil {
		resErr = err
		return
	}
	obj.MockPath = mockPath
	action = &obj
	// 读取并加载 config 文件
	action.SetMockServeConfig()

	return
}

// 读取并更新 MockServeConfig
func (obj *MockActionObj) SetMockServeConfig() (resErr error) {
	resErr = nil

	config, err := ReadMockServeInfo(obj.MockPath.ConfigFullPath)
	if err != nil {
		resErr = err
		return
	}
	obj.MockServeConfig = config

	return
}

// 存储 MockServeConfig
func (obj *MockActionObj) WriteMockServeConfig() (resErr error) {
	resErr = nil
	config := m_json.ToJson(obj.MockServeConfig)
	m_file.WriteByte(obj.MockPath.ConfigFullPath, config)
	return
}
