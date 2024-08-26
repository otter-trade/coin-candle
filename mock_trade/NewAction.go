package mock_trade

import (
	"fmt"

	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_json"
	jsoniter "github.com/json-iterator/go"
	"github.com/otter-trade/coin-candle/global"
)

type MockActionObj struct {
	StrategyID      string // 策略的Id
	MockName        string // 本次回测的名称
	MockPath        MockPathType
	MockServeConfig global.MockServeConfigType
	PositionIndex   global.PositionIndexType
	NewPosition     []NewPositionType
}

type NewMockActionOpt struct {
	StrategyID string // 策略的Id
	MockName   string // 本次回测的名称
}

// #### New 一个 Action 对象  ####
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

// #### 读取并更新 MockServeConfig ####
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

// #### 本地存储一下 MockServeConfig ####
func (obj *MockActionObj) StoreMockServeConfig() (resErr error) {
	resErr = nil
	config := m_json.ToJson(obj.MockServeConfig)
	err := m_file.WriteByte(obj.MockPath.ConfigFullPath, config)
	if err != nil {
		resErr = err
		return
	}
	return
}

// #### 读取最新的 PositionIndex ####
func (obj *MockActionObj) SetPositionIndex() (resErr error) {
	resErr = nil

	if len(obj.MockPath.PositionIndexFullPath) < 20 {
		resErr = fmt.Errorf("该 PositionIndex 目录不正确")
		return
	}
	PositionIndexByte := m_file.ReadFile(obj.MockPath.PositionIndexFullPath)
	var PositionIndex global.PositionIndexType
	err := jsoniter.Unmarshal(PositionIndexByte, &PositionIndex)
	if err != nil {
		resErr = err
		return
	}

	obj.PositionIndex = PositionIndex
	return
}

// #### 把当前的 PositionIndex 进行一次本地存储 ####
func (obj *MockActionObj) StorePositionIndex() (resErr error) {
	if len(obj.MockPath.PositionIndexFullPath) < 20 {
		resErr = fmt.Errorf("该 PositionIndex 目录不正确")
		return
	}
	config := m_json.ToJson(obj.MockServeConfig)
	err := m_file.WriteByte(obj.MockPath.ConfigFullPath, config)
	if err != nil {
		resErr = err
		return
	}
	return
}
