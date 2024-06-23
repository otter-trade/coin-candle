package binance

import (
	"coin-candle/global"
	"fmt"

	"github.com/handy-golang/go-tools/m_count"
	"github.com/handy-golang/go-tools/m_fetch"
	"github.com/handy-golang/go-tools/m_json"
	jsoniter "github.com/json-iterator/go"
)

func GetTicker() (resData []global.BinanceTickerType, resErr error) {

	resData = nil
	resErr = nil

	fetchData, err := m_fetch.NewHttp(m_fetch.HttpOpt{
		Origin:    BaseUrl,
		Path:      "/api/v3/ticker/24hr",
		ProxyURLs: global.Path.ProxyURLs,
	}).Get()
	if err != nil {
		resErr = err
		return
	}

	var result []global.BinanceTickerType
	jsoniter.Unmarshal(fetchData, &result)

	if len(result) < 2 {
		resErr = fmt.Errorf("错误:结果返回不正确 %+v", m_json.ToStr(result))
		return
	}
	resData = result
	return
}

// 成交量排序
func VolumeSort(arr []global.BinanceTickerType) []global.BinanceTickerType {
	size := len(arr)
	list := make([]global.BinanceTickerType, size)
	copy(list, arr)

	var swapped bool
	for i := size - 1; i > 0; i-- {
		swapped = false
		for j := 0; j < i; j++ {
			a := list[j+1].QuoteVolume
			b := list[j].QuoteVolume
			if m_count.Le(a, b) < 0 {
				list[j], list[j+1] = list[j+1], list[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return list
}

// 翻转数组
func Reverse(arr []global.BinanceTickerType) []global.BinanceTickerType {
	list := make(
		[]global.BinanceTickerType,
		len(arr),
		len(arr)*2,
	)

	j := 0
	for i := len(arr) - 1; i > -1; i-- {
		list[j] = arr[i]
		j++
	}

	return list
}
