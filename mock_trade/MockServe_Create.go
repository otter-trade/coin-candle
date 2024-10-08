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
	"github.com/handy-golang/go-tools/m_time"
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

	// 检查配置文件是否存在
	var config global.MockServeConfigType
	isExist := m_path.IsExist(mockPath.ConfigFullPath)
	if isExist {
		resErr = fmt.Errorf("该MockServe已存在")
		fileCont := m_file.ReadFile(mockPath.ConfigFullPath)
		err := json.Unmarshal(fileCont, &config)
		if err != nil {
			resErr = err // 存在但是解析有问题
			return
		}
		resData = config // 依然返回配置文件内容
		return
	}

	// 描述 不允许包含特殊字符串
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

	if m_count.Le(InitialAsset, "1001") < 0 {
		config.SysWarn += m_str.Join(
			"初始资产支持任何正数，但尽量使用 一千、一万、一万、十万这样的大整数;",
		)
	}

	FeeRate := global.DefaultFeeRate // 设置默认手续费
	// 当用户传递了手续费，则使用用户的手续费，但是要使用 m_count 过滤一下
	if len(opt.FeeRate) > 0 {
		FeeRate = m_count.Sub(opt.FeeRate, "0")
		// 思考了一下，虚拟持仓手续费允许为 0 和 负数，所以无需判断，但是应当给出提醒
		if m_count.Le(FeeRate, "0.1") > 0 || m_count.Le(FeeRate, "0") < 0 {
			config.SysWarn += m_str.Join(
				"手续费允许为0和负数,默认值为",
				global.DefaultFeeRate,
				"(也就是",
				m_count.Mul(global.DefaultFeeRate, "100"),
				"%) ,您的手续费超过了正常范围，请注意;",
			)
		}
	}

	config.StrategyID = opt.StrategyID
	config.MockName = opt.MockName
	config.Description = opt.Description
	config.InitialAsset = InitialAsset
	config.FeeRate = FeeRate
	config.RunMode = RunMode
	config.MockDataDir = mockPath.MockDataDir
	config.ConfigPath = mockPath.ConfigPath
	config.PositionIndexPath = mockPath.PositionIndexPath
	config.CreateTime = m_time.GetUnixInt64()

	resData = config

	// 检查是否超出最大条目
	StrategyDir := m_str.Join(
		global.Path.MockTradeDir,
		os.PathSeparator,
		opt.StrategyID,
	)
	isExist = m_path.IsExist(StrategyDir)
	if isExist {
		files, err := os.ReadDir(StrategyDir)
		if err != nil {
			resErr = err
			return
		}

		if len(files) > global.MaxMockServeCount {
			resErr = fmt.Errorf("超出最大条目,该条 MockServe 将不会写入磁盘。")
			return
		}
	}

	m_file.WriteByte(mockPath.ConfigFullPath, m_json.ToJson(config))
	m_file.WriteByte(mockPath.PositionIndexFullPath, m_json.ToJson(global.PositionIndexType{})) // 建立索引文件

	return
}
