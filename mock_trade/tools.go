package mock_trade

import (
	"fmt"
	"os"
	"regexp"

	"github.com/handy-golang/go-tools/m_str"
	"github.com/otter-trade/coin-candle/global"
)

type MockPathType struct {
	MockDataDir           string
	MockDataFullDir       string
	ConfigPath            string
	ConfigFullPath        string
	PositionIndexPath     string
	PositionIndexFullPath string
}

func CheckMockName(opt global.FindMockServeOpt) (resData MockPathType, resErr error) {
	resData = MockPathType{}
	resErr = nil

	// StrategyID 不能为空
	if len(opt.StrategyID) < 1 {
		resErr = fmt.Errorf("StrategyID不能为空")
		return
	}
	// MockName 必须为2-24位字母数字下划线和中文
	isMockNameReg := IsMockNameReg(opt.MockName)
	if !isMockNameReg {
		resErr = fmt.Errorf("MockName必须为2-24位字母数字下划线和中文")
		return
	}

	MockDataDir := m_str.Join(
		opt.StrategyID,
		os.PathSeparator,
		opt.MockName,
		os.PathSeparator,
	)

	MockDataFullDir := m_str.Join(
		global.Path.MockTradeDir,
		os.PathSeparator,
		opt.StrategyID,
		os.PathSeparator,
		opt.MockName,
		os.PathSeparator,
	)

	ConfigPath := m_str.Join(
		MockDataDir,
		"config.json",
	)
	ConfigFullPath := m_str.Join(
		MockDataFullDir,
		"config.json",
	)

	PositionIndexPath := m_str.Join(
		MockDataDir,
		"position_index.json",
	)
	PositionIndexFullPath := m_str.Join(
		MockDataFullDir,
		"position_index.json",
	)

	resData = MockPathType{
		MockDataDir:           MockDataDir,
		MockDataFullDir:       MockDataFullDir,
		ConfigPath:            ConfigPath,
		ConfigFullPath:        ConfigFullPath,
		PositionIndexPath:     PositionIndexPath,
		PositionIndexFullPath: PositionIndexFullPath,
	}

	return
}

// 特殊字符的检查
func IsDescReg(str string) bool {
	pattern := "[<>/|{}\\[\\]\\\\:;\"\\`\\*\\s\\'\\\"]"
	reg := regexp.MustCompile(pattern)
	return !reg.MatchString(str)
}

// 2-24位字母数字下划线和中文
func IsMockNameReg(str string) bool {
	pattern := "^[a-zA-Z0-9_\u4e00-\u9fa5]{2,24}$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(str)
}
