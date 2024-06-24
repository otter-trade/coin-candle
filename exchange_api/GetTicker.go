package exchange_api

import (
	"coin-candle/exchange_api/binance"
	"coin-candle/exchange_api/okx"
	"coin-candle/global"
	"fmt"
	"os"
	"time"

	"github.com/handy-golang/go-tools/m_count"
	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_json"
	"github.com/handy-golang/go-tools/m_path"
	"github.com/handy-golang/go-tools/m_str"
	"github.com/handy-golang/go-tools/m_time"
	jsoniter "github.com/json-iterator/go"
)

func GetTickerPath() string {
	// 相当于保留每天的榜单数据
	var timeDay = time.Now().Format("2006-01-02")
	var filePath = m_str.Join(
		global.Path.DataPath,
		os.PathSeparator,
		"ticker-",
		timeDay,
		".json",
	)
	return filePath
}

func GetTickerList() (resData []global.TickerType, resErr error) {
	resData = nil
	resErr = nil

	var filePath = GetTickerPath()

	// 如果该文件不存在
	if !m_path.IsExist(filePath) {
		UpdateLocalGoodsList()
	}

	var fileData = m_file.ReadFile(filePath)
	if len(fileData) < 2 {
		resErr = fmt.Errorf("文件读取失败: %s", filePath)
		return
	}

	var TickerList []global.TickerType
	jsoniter.Unmarshal(fileData, &TickerList)
	if len(TickerList) < 10 {
		resErr = fmt.Errorf("错误:结果返回不正确 %+v", m_json.ToStr(TickerList))
		return
	}

	resData = TickerList

	return
}

// 更新本地榜单数据
func UpdateLocalTicker() {
	var BinanceTicker = GetBinanceTicker()
	m_file.WriteByte(global.Path.DataPath+"/binance-ticker.json", m_json.ToJson(BinanceTicker))
	var OkxTicker = GetOkxTicker()
	m_file.WriteByte(global.Path.DataPath+"/okx-ticker.json", m_json.ToJson(OkxTicker))

	// 将两个 ticker 合并
	tickerList := []global.TickerType{}
	for _, binance := range BinanceTicker {
		goods, _ := GetGoodsDetail(GetGoodsDetailOpt{
			Binance_Symbol: binance.Symbol,
		})
		for _, okx := range OkxTicker {
			if okx.InstID == goods.Okx_SPOT_Info.InstID {
				// 匹配到了同一款商品,将两者混合
				ticker := TickerMix(TickerMixOpt{
					GoodsInfo:     goods,
					OkxTicker:     okx,
					BinanceTicker: binance,
				})
				if len(ticker.Open24H) > 0 {
					tickerList = append(tickerList, ticker)
				}
				break
			}
		}
	}

	VolumeSortList := SortVolume(tickerList) // 按照成交量排序

	var filePath = GetTickerPath()

	m_file.WriteByte(filePath, m_json.ToJson(VolumeSortList))
	global.RunLog.Println("交易所榜单更新完成", filePath)
}

type TickerMixOpt struct {
	GoodsInfo     global.GoodsType
	OkxTicker     global.OkxTickerType
	BinanceTicker global.BinanceTickerType
}

func TickerMix(opt TickerMixOpt) (Ticker global.TickerType) {
	Ticker = global.TickerType{}

	var OKXTicker = opt.OkxTicker
	var BinanceTicker = opt.BinanceTicker
	var GoodsInfo = opt.GoodsInfo

	Okx_SPOT_Info := GoodsInfo.Okx_SPOT_Info
	Okx_SWAP_Info := GoodsInfo.Okx_SWAP_Info

	// 合约现货都得有才行
	if len(Okx_SPOT_Info.InstID) > 2 && len(Okx_SWAP_Info.InstID) > 2 {

	} else {
		return
	}

	// 合约 上架小于 36 天的不计入榜单
	diffOnLine := m_count.Sub(OKXTicker.Ts, Okx_SWAP_Info.ListTime)
	diffDay := m_count.Div(diffOnLine, m_time.UnixTime.Day) // 转化为天数
	if m_count.Le(diffDay, "36") < 0 {                      // 少于 36 天的时候
		return
	}

	Ticker.GoodsId = GoodsInfo.GoodsId
	Ticker.Okx_InstID = OKXTicker.InstID
	Ticker.Binance_Symbol = BinanceTicker.Symbol
	Ticker.BaseCcy = GoodsInfo.BaseCcy
	Ticker.State = GoodsInfo.State

	Ticker.Last = OKXTicker.Last
	Ticker.Open24H = OKXTicker.Open24H
	Ticker.High24H = OKXTicker.High24H
	Ticker.Low24H = OKXTicker.Low24H
	Ticker.OKXVol24H = OKXTicker.VolCcy24H
	Ticker.BinanceVol24H = BinanceTicker.QuoteVolume
	Ticker.U_R24 = m_count.RoseCent(OKXTicker.Last, OKXTicker.Open24H)
	Ticker.Volume = m_count.Add(OKXTicker.VolCcy24H, BinanceTicker.QuoteVolume)
	Ticker.OkxVolRose = m_count.PerCent(Ticker.OKXVol24H, Ticker.Volume)
	Ticker.BinanceVolRose = m_count.PerCent(Ticker.BinanceVol24H, Ticker.Volume)

	var timeObj = m_time.TimeGet(OKXTicker.Ts)
	Ticker.TimeUnix = timeObj.TimeUnix
	Ticker.TimeStr = timeObj.TimeStr

	return Ticker
}

func SortVolume(data []global.TickerType) []global.TickerType {
	size := len(data)
	list := make([]global.TickerType, size)
	copy(list, data)

	var swapped bool
	for i := size - 1; i > 0; i-- {
		swapped = false
		for j := 0; j < i; j++ {
			a := list[j+1].Volume
			b := list[j].Volume
			if m_count.Le(a, b) < 0 {
				list[j], list[j+1] = list[j+1], list[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	// 设置 VolIdx 并翻转
	listIDX := []global.TickerType{}
	j := 0
	for i := len(list) - 1; i > -1; i-- {
		Kdata := list[i]
		listIDX = append(listIDX, Kdata)
		j++
	}

	return listIDX
}

func GetBinanceTicker() []global.BinanceTickerType {
	binanceTicker, err := binance.GetTicker()
	if err != nil {
		global.LogErr("错误:exchange_api.GetTicker -> binance.GetTicker", err)
	}
	// 过滤数据,确认该商品存在
	var list = []global.BinanceTickerType{}
	for _, item := range binanceTicker {
		goods, err := GetGoodsDetail(GetGoodsDetailOpt{
			Binance_Symbol: item.Symbol,
		})
		if err != nil {
			continue
		}
		if len(goods.GoodsId) > 2 {
			list = append(list, item)
		}
	}

	// 按照成交量排序
	VolumeList := binance.VolumeSort(list)
	tLen := len(VolumeList)
	if tLen > 21 {
		VolumeList = VolumeList[tLen-20:] // 取出最后 20 个
	}
	TickerList := binance.Reverse(VolumeList) // 翻转数组 成交量最大的排在前面
	return TickerList
}

func GetOkxTicker() []global.OkxTickerType {
	okxTicker, err := okx.GetTicker()
	if err != nil {
		global.LogErr("错误:exchange_api.GetTicker -> okx.GetTicker", err)
	}

	// 过滤数据,确认该商品存在
	var list = []global.OkxTickerType{}
	for _, item := range okxTicker {
		goods, err := GetGoodsDetail(GetGoodsDetailOpt{
			Okx_InstID: item.InstID,
		})
		if err != nil {
			continue
		}
		if len(goods.GoodsId) > 2 {
			list = append(list, item)
		}
	}

	// 按照成交量排序
	VolumeList := okx.VolumeSort(list)
	tLen := len(VolumeList)
	if tLen > 21 {
		VolumeList = VolumeList[tLen-20:] // 取出最后 20 个
	}
	TickerList := okx.Reverse(VolumeList) // 翻转数组 成交量最大的排在前面
	return TickerList
}
