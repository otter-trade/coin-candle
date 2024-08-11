package demo

import (
	"fmt"

	"github.com/handy-golang/go-tools/m_json"
	"github.com/otter-trade/coin-candle/global"
	"github.com/otter-trade/coin-candle/mock_trade"
)

/*
#### 创建一个 MockServe ####
*/
func CreateMockServe_demo() {
	resData, err := mock_trade.CreateMockServe(global.CreateMockServeOpt{
		StrategyID:   "mo7_StrategyID_001",
		MockName:     "测试_MockName_3",
		RunMode:      "1",
		InitialAsset: "1000",
	})
	if err != nil {
		fmt.Println("创建持仓失败", err)
	}
	m_json.Println(resData)
}

/*
####  查看某个 MockServe 的详情 ####
*/
func GetMockServeInfo_demo() {
	mockServeInfo, err := mock_trade.GetMockServeInfo(global.FindMockServeOpt{
		StrategyID: "mo7_StrategyID_001",
		MockName:   "测试_MockName_1",
	})
	if err != nil {
		fmt.Println("获取MockServe信息失败", err)
	}
	m_json.Println(mockServeInfo)
}

/*
####  查看 MockServe 列表 ####
*/
func GetMockServeList_demo() {
	mockServeList := mock_trade.GetMockServeList(global.FindMockServeListOpt{
		StrategyID: "mo7_StrategyID_001",
	})
	m_json.Println(mockServeList)
}

/*
#### 删除一个 MockServe ####
*/
func DeleteMockServe_demo() {
	err := mock_trade.DeleteMockServe(global.FindMockServeOpt{
		StrategyID: "mo7_StrategyID_001",
		MockName:   "测试_MockName_4",
	})
	if err != nil {
		fmt.Println("删除虚拟持仓失败", err)
	}
}

/*
#### 删除一个策略 ####
该策略下的 MockServeList 为空时才可以删除
*/
func ClearStrategy_demo() {
	err := mock_trade.ClearStrategy("mo7_StrategyID_001")
	if err != nil {
		fmt.Println("删除策略失败", err)
	}
}
