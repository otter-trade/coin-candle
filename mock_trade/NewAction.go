package mock_trade

import (
	"fmt"

	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_json"
	jsoniter "github.com/json-iterator/go"
	"github.com/otter-trade/coin-candle/global"
)

type MockActionObj struct {
	StrategyID      string                     // 策略的Id
	MockName        string                     // 本次回测的名称
	Path            MockPathType               // 相关数据存储的路径
	MockServeConfig global.MockServeConfigType // MockServe 的配置文件内容
	PositionIndex   global.PositionIndexType   // 每次变更持仓的时间戳列表
	NewPosition     []NewPositionType          // 新的持仓列表
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
	obj.Path = mockPath

	action = &obj
	// 读取并加载 config 文件
	action.ReadMockServeConfig()

	return
}

// #### 读取 MockServeConfig ####
func (obj *MockActionObj) ReadMockServeConfig() (resErr error) {
	resErr = nil

	config, err := ReadMockServeInfo(obj.Path.ConfigFullPath)
	if err != nil {
		resErr = err
		return
	}
	obj.MockServeConfig = config

	return
}

// #### 存储 MockServeConfig ####
func (obj *MockActionObj) StoreMockServeConfig() (resErr error) {
	resErr = nil
	config := m_json.ToJson(obj.MockServeConfig)
	err := m_file.WriteByte(obj.Path.ConfigFullPath, config)
	if err != nil {
		resErr = err
		return
	}
	return
}

// #### 读取 PositionIndex : 也就是持仓列表的索引 ####
func (obj *MockActionObj) ReadPositionIndex() (resErr error) {
	resErr = nil

	if len(obj.Path.PositionIndexFullPath) < 20 {
		resErr = fmt.Errorf("该 PositionIndex 目录不正确")
		return
	}
	PositionIndexByte := m_file.ReadFile(obj.Path.PositionIndexFullPath)
	var PositionIndex global.PositionIndexType
	err := jsoniter.Unmarshal(PositionIndexByte, &PositionIndex)
	if err != nil {
		resErr = err
		return
	}

	obj.PositionIndex = PositionIndex
	return
}

// #### 存储 PositionIndex ####
func (obj *MockActionObj) StorePositionIndex() (resErr error) {
	if len(obj.Path.PositionIndexFullPath) < 20 {
		resErr = fmt.Errorf("该 PositionIndex 目录不正确")
		return
	}
	config := m_json.ToJson(obj.MockServeConfig)
	err := m_file.WriteByte(obj.Path.ConfigFullPath, config)
	if err != nil {
		resErr = err
		return
	}
	return
}
