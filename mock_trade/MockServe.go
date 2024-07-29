package mock_trade

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/handy-golang/go-tools/m_count"
	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_json"
	"github.com/handy-golang/go-tools/m_path"
	"github.com/handy-golang/go-tools/m_str"
	"github.com/otter-trade/coin-candle/global"
)

/*
	模拟交易模块。
	设计了一套仓位管理模式，可以将策略逻辑和交易所 进行有效隔离。内建模拟交易好处:
	1. 有效简化策略开发流程。因为是内建的虚拟仓位，所以在编写策略时，只需要关注策略逻辑即可。
	2. 方便回测。在回测模式下，开发者可以在1分钟之内完成 一整年的交易订单的模拟，并查看结果，方便策略的修改。
	3. 回测-实盘 无缝对接。开发者在写完策略后，只需要简单修改一个参数，即可上架实盘。保证实盘效果和回测结果的无限接近。

	基本逻辑：
  注册一个虚拟持仓，并标注初始资产。
	不断地更新仓位，系统会记录下仓位变化，如果为 实盘模式 则同步给下单模块更新至交易所。
	可以读取任意时间点的仓位情况，包括当前时间点的收益。并根据持仓情况，对这个时间点的策略情况进行打分和评估。

	支持读取全部的持仓记录。

	基于本地文件系统，支持超高频率的读取和更新。
*/

// 创建 MockServe
func CreateMockServe(opt global.CreateMockServeOpt) (resData global.MockServeConfigType, resErr error) {
	resData = global.MockServeConfigType{}
	resErr = nil

	// 检查 StrategyID 和 MockName 并获取存储目录
	mockPath, err := CheckMockName(global.FindMockServeOpt{
		StrategyID: opt.StrategyID,
		MockName:   opt.MockName,
	})
	if err != nil {
		resErr = err
		return
	}

	isDesc := IsDescReg(opt.Description)
	if !isDesc {
		resErr = fmt.Errorf("Description禁止包含特殊符号")
		return
	}

	// 检查 RunMode
	RunMode, err := global.GetRunMode(opt.RunMode)
	if err != nil {
		resErr = err
		return
	}

	// 检查初始资产
	InitialAsset := m_count.Sub(opt.InitialAsset, "0")
	if m_count.Le(InitialAsset, global.DefaultInitialAsset) < 0 {
		InitialAsset = global.DefaultInitialAsset
	}

	// 检查手续费率 大于 0.5 的活阎王 和 0 都重置为 初始值
	FeeRate := m_count.Sub(opt.FeeRate, "0")
	if m_count.Le(FeeRate, "0.9") >= 0 || m_count.Le(FeeRate, "0") == 0 {
		FeeRate = global.DefaultFeeRate
	}

	// 检查是否存在
	var config global.MockServeConfigType
	isExist := m_path.IsExist(mockPath.ConfigPath)
	if isExist {
		resErr = fmt.Errorf("该MockServe已存在")
		fileCont := m_file.ReadFile(mockPath.ConfigPath)
		err := json.Unmarshal(fileCont, &config)
		if err != nil {
			resErr = err
			return
		}
		resData = config
		return
	}

	config.StrategyID = opt.StrategyID
	config.MockName = opt.MockName
	config.Description = opt.Description
	config.InitialAsset = InitialAsset
	config.FeeRate = FeeRate
	config.RunMode = RunMode
	config.DataDir = mockPath.MockDataDir
	config.DataIndex = []string{}

	resData = config

	// 检查是否超出最大条目
	StrategyDir := m_str.Join(
		global.Path.MockTradeDir,
		os.PathSeparator,
		opt.StrategyID,
	)
	files, err := os.ReadDir(StrategyDir)
	if err != nil {
		resErr = err
		return
	}
	if len(files) > global.MaxMockServeCount {
		resErr = fmt.Errorf("超出最大条目,该条 MockServe 将不会写入磁盘。")
		return
	}

	m_file.Write(mockPath.ConfigPath, m_json.ToStr(config))

	return
}

// 删除一个 MockServe
func DeleteMockServe(opt global.FindMockServeOpt) (resErr error) {
	resErr = nil
	// 检查 StrategyID 和 MockName 并获取存储目录
	mockPath, err := CheckMockName(global.FindMockServeOpt{
		StrategyID: opt.StrategyID,
		MockName:   opt.MockName,
	})
	if err != nil {
		resErr = err
		return
	}

	err = os.Remove(mockPath.MockDataDir)
	if err != nil {
		resErr = err
		return
	}

	return
}

// 获取某个策略下所有的 MockServe
func GetMockServeList(StrategyID string) (resData []global.MockServeConfigType) {
	resData = []global.MockServeConfigType{}
	if len(StrategyID) < 1 {
		return
	}

	StrategyDir := m_str.Join(
		global.Path.MockTradeDir,
		os.PathSeparator,
		StrategyID,
	)

	isExist := m_path.IsExist(StrategyDir)
	if !isExist {
		return
	}

	files, err := os.ReadDir(StrategyDir)
	if err != nil {
		return
	}

	for _, v := range files {
		if v.Type().IsDir() {
			MockName := v.Name()

			configPath := m_str.Join(
				StrategyDir,
				os.PathSeparator,
				MockName,
				os.PathSeparator,
				"config.json",
			)
			isExist := m_path.IsExist(StrategyDir)
			if !isExist {
				continue
			}

			var config global.MockServeConfigType
			fileCont := m_file.ReadFile(configPath)
			err = json.Unmarshal(fileCont, &config)
			if err != nil {
				continue
			}
			resData = append(resData, config)
		}
	}

	return
}

// 获取一个 MockServe 的详情
func GetMockServeInfo(opt global.FindMockServeOpt) (resData global.MockServeConfigType, resErr error) {
	resData = global.MockServeConfigType{}
	resErr = nil
	// 检查 StrategyID 和 MockName 并获取存储目录
	mockPath, err := CheckMockName(global.FindMockServeOpt{
		StrategyID: opt.StrategyID,
		MockName:   opt.MockName,
	})
	if err != nil {
		resErr = err
		return
	}

	var config global.MockServeConfigType
	fileCont := m_file.ReadFile(mockPath.ConfigPath)
	err = json.Unmarshal(fileCont, &config)
	if err != nil {
		resErr = err
		return
	}

	if len(config.MockName) < 1 {
		resErr = fmt.Errorf("解析失败")
		return
	}

	resData = config

	return
}

// 删除一个策略
func ClearStrategy(StrategyID string) (resErr error) {
	resErr = nil
	// 检查是否超出最大条目
	StrategyDir := m_str.Join(
		global.Path.MockTradeDir,
		os.PathSeparator,
		StrategyID,
	)
	files, err := os.ReadDir(StrategyDir)
	if err != nil {
		resErr = err
		return
	}
	if len(files) > 0 {
		resErr = fmt.Errorf("请先删除该策略下的所有 MockServe")
		return
	}
	return
}
