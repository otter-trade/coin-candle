package okx

import (
	"fmt"

	"github.com/otter-trade/coin-candle/global"

	"github.com/handy-golang/go-tools/m_count"
	"github.com/handy-golang/go-tools/m_fetch"
	"github.com/handy-golang/go-tools/m_json"
	jsoniter "github.com/json-iterator/go"
)

type OkxTickerReqType struct {
	Code string                 `json:"code"`
	Msg  string                 `json:"msg"`
	Data []global.OkxTickerType `json:"data"`
}

func GetTicker() (resData []global.OkxTickerType, resErr error) {
	resData = nil
	resErr = nil

	fetchData, err := m_fetch.NewHttp(m_fetch.HttpOpt{
		Origin: BaseUrl,
		Path:   "/api/v5/market/tickers",
		DataMap: map[string]any{
			"instType": "SPOT",
		},
		ProxyURLs: global.Path.ProxyURLs,
	}).Get()

	if err != nil {
		resErr = err
		return
	}

	var result OkxTickerReqType
	jsoniter.Unmarshal(fetchData, &result)
	if result.Code != "0" {
		resErr = fmt.Errorf("错误:结果返回不正确 %+v", m_json.ToStr(fetchData))
		return
	}

	if len(result.Data) < 2 {
		resErr = fmt.Errorf("错误:数据长度不正确: %+v", result)
		return
	}

	resData = result.Data

	return
}

// 成交量排序
func VolumeSort(arr []global.OkxTickerType) []global.OkxTickerType {
	size := len(arr)
	list := make([]global.OkxTickerType, size)
	copy(list, arr)

	var swapped bool
	for i := size - 1; i > 0; i-- {
		swapped = false
		for j := 0; j < i; j++ {
			a := list[j+1].VolCcy24H
			b := list[j].VolCcy24H
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
func Reverse(arr []global.OkxTickerType) []global.OkxTickerType {
	list := make(
		[]global.OkxTickerType,
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
