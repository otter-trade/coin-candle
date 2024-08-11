package global

/*
初始化目录
*/

import (
	"os"

	"github.com/handy-golang/go-tools/m_path"
	"github.com/handy-golang/go-tools/m_str"
)

type ExchangeDir struct {
	Dir string
}

var Path struct {
	Home          string      // Home 根目录
	App           string      // APP 根目录
	DataPath      string      // 数据文件存放目录
	LogPath       string      // 日志文件存放目录
	ProxyURLs     []string    // 代理地址
	Binance       ExchangeDir // 币安数据目录
	Okx           ExchangeDir // 欧意数据目录
	GoodsListFile string      // 商品列表文件
	MockTradeDir  string      // 虚拟持仓数据目录
}

// 初始化日志目录 ，且必须为有效目录
func init_Path(opt SysInitOpt) {
	Path.Home = m_path.GetHomePath()
	Path.App = m_path.GetPwd()

	if len(opt.ProxyURLs) > 0 {
		Path.ProxyURLs = opt.ProxyURLs
	}

	// 初始化日志目录
	if len(opt.LogPath) > 1 {
		Path.LogPath = opt.LogPath
	} else {
		Path.LogPath = m_str.Join(
			Path.App,
			os.PathSeparator,
			"logs",
		)
	}
	// 不存在则新建
	isLogoPath := m_path.IsExist(Path.LogPath)
	if !isLogoPath {
		os.MkdirAll(Path.LogPath, os.ModePerm)
	}

	// 初始化数据目录
	if len(opt.DataPath) > 1 {
		Path.DataPath = opt.DataPath
	} else {
		Path.DataPath = m_str.Join(
			Path.App,
			os.PathSeparator,
			"data",
		)
	}
	// 不存在则新建
	isDataPath := m_path.IsExist(Path.DataPath)
	if !isDataPath {
		os.MkdirAll(Path.DataPath, os.ModePerm)
	}

	// 初始化 币安数据 目录
	Path.Binance.Dir = m_str.Join(
		Path.DataPath,
		os.PathSeparator,
		ExchangeOpt[0],
	)
	if !m_path.IsExist(Path.Binance.Dir) {
		os.MkdirAll(Path.Binance.Dir, os.ModePerm)
	}

	// 初始化 欧意数据目录
	Path.Okx.Dir = m_str.Join(
		Path.DataPath,
		os.PathSeparator,
		ExchangeOpt[1],
	)
	if !m_path.IsExist(Path.Okx.Dir) {
		os.MkdirAll(Path.Okx.Dir, os.ModePerm)
	}

	// 商品列表文件路径
	Path.GoodsListFile = m_str.Join(
		Path.DataPath,
		os.PathSeparator,
		"goods-list.json",
	)

	// 初始化虚拟持仓数据目录
	Path.MockTradeDir = m_str.Join(
		Path.DataPath,
		os.PathSeparator,
		"mock-trade",
	)

	if !m_path.IsExist(Path.MockTradeDir) {
		os.MkdirAll(Path.MockTradeDir, os.ModePerm)
	}
}
