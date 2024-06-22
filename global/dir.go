package global

import (
	"os"

	"github.com/handy-golang/go-tools/m_path"
	"github.com/handy-golang/go-tools/m_str"
)

var Dir struct {
	Home     string // Home 根目录
	App      string // APP 根目录
	LogPath  string // 日志文件存放目录
	DataPath string // 数据文件存放目录
}

// 初始化日志目录 ，且必须为有效目录
func DirInit(opt Opt) {
	Dir.Home = m_path.GetHomePath()
	Dir.App = m_path.GetPwd()

	// 初始化目录

	if len(opt.LogPath) > 1 {
		Dir.LogPath = opt.LogPath
	} else {
		Dir.LogPath = m_str.Join(
			Dir.App,
			m_str.ToStr(os.PathSeparator),
			"logs",
		)
	}

	// 初始化数据目录 ，且必须为有效目录
	if len(opt.DataPath) > 1 {
		Dir.DataPath = opt.DataPath
	} else {
		Dir.DataPath = m_str.Join(
			Dir.App,
			m_str.ToStr(os.PathSeparator),
			"data",
		)
	}

	// 目录不存在则新建
	isLogoPath := m_path.Exists(Dir.LogPath)
	if !isLogoPath {
		os.MkdirAll(Dir.LogPath, os.ModePerm)
	}

	isDataPath := m_path.Exists(Dir.DataPath)
	if !isDataPath {
		os.MkdirAll(Dir.DataPath, os.ModePerm)
	}

}
