package mock_trade

import (
	"fmt"
	"os"

	"github.com/handy-golang/go-tools/m_path"
	"github.com/handy-golang/go-tools/m_str"
	"github.com/otter-trade/coin-candle/global"
)

// 获取某个策略下所有的 MockServe
func GetMockServeList(opt global.FindMockServeListOpt) (resData []global.MockServeConfigType) {
	resData = []global.MockServeConfigType{}
	if len(opt.StrategyID) < 1 {
		return
	}

	StrategyDir := m_str.Join(
		global.Path.MockTradeDir,
		os.PathSeparator,
		opt.StrategyID,
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

			config, err := ReadMockServeInfo(configPath)
			if err != nil {
				continue
			}

			if opt.CreateTime[1] > opt.CreateTime[0] { // 创建时间的查询过滤
				if (config.CreateTime >= opt.CreateTime[0]) && (config.CreateTime <= opt.CreateTime[1]) {
					resData = append(resData, config)
				}
			} else if opt.LastPositionUpdateTime[1] > opt.LastPositionUpdateTime[0] { // 最后一次持仓更新时间查询过滤
				if (config.LastPositionUpdateTime >= opt.LastPositionUpdateTime[0]) && (config.LastPositionUpdateTime <= opt.LastPositionUpdateTime[1]) {
					resData = append(resData, config)
				}
			} else { // 没有时间过滤条件
				resData = append(resData, config)
			}

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

	config, err := ReadMockServeInfo(mockPath.ConfigFullPath)
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
