package exchange_api

import (
	"coin-candle/exchange_api/binance"
	"coin-candle/exchange_api/okx"
	"coin-candle/global"
	"encoding/json"
	"fmt"
	"os"

	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_json"
	"github.com/handy-golang/go-tools/m_path"
	"github.com/handy-golang/go-tools/m_str"
	"github.com/handy-golang/go-tools/m_time"
)

func GetKline(opt global.GetKlineOpt) (resData []global.KlineType, resErr error) {
	resData = nil
	resErr = nil

	// 检查参数 GoodsId
	if len(opt.GoodsId) < 1 {
		resErr = fmt.Errorf("缺少 GoodsId")
		return
	}
	GoodsDetail, err := GetGoodsDetail(GetGoodsDetailOpt{
		GoodsId: opt.GoodsId,
	})
	if err != nil {
		resErr = err
		return
	}

	// 检查参数 Bar
	BarObj := global.GetBarOpt(opt.Bar)
	if BarObj.Interval < m_time.UnixTimeInt64.Minute {
		resErr = fmt.Errorf("参数 Bar 错误")
		return
	}

	// Before 缺省值
	now := m_time.GetUnixInt64()
	Before := now
	// 时间 传入的时间戳 必须大于6年前 才有效，否则等价与当前时间戳
	if opt.Before > now-m_time.UnixTimeInt64.Day*2190 {
		Before = opt.Before
	}

	// Limit 缺省值
	Limit := 10
	if opt.Limit > 1 && opt.Limit < 500 {
		Limit = opt.Limit
	} else {
		resErr = fmt.Errorf("Limit必须为1-500的正整数")
		return
	}

	// Exchange 缺省值, 过滤有效值
	var Exchange []string
	if len(opt.Exchange) < 1 {
		Exchange = []string{"okx"}
	} else {
		for _, item := range global.ExchangeOpt {
			for _, exchange := range opt.Exchange {
				if exchange == item {
					Exchange = append(Exchange, item)
					break
				}
			}
		}
	}
	if len(Exchange) < 1 {
		resErr = fmt.Errorf("参数 Exchange 无效")
		return
	}

	// 计算出起止时间  // 结束时间 - 时间间隔 * 条数
	EndTime := Before
	StartTime := EndTime - BarObj.Interval*int64(Limit)

	SendParamList := GetKlineFilePath(GetKlineFilePathOpt{
		Limit:     Limit,
		StartTime: StartTime,
		EndTime:   EndTime,
		BarObj:    BarObj,
		Goods:     GoodsDetail,
		Exchange:  Exchange,
	})

	for _, item := range SendParamList {

		m_json.Println(item)
		// kline, err := SendKlineRequest(item)
		// if err != nil {
		// 	resErr = err
		// 	return
		// }
	}

	return
}

type GetKlineFilePathOpt struct {
	Limit     int
	StartTime int64
	EndTime   int64
	BarObj    global.KlineBarType
	Goods     global.GoodsType
	Exchange  []string
}

func GetKlineFilePath(opt GetKlineFilePathOpt) (resData []SendKlineRequestOpt) {
	resData = []SendKlineRequestOpt{}
	for _, exchange := range opt.Exchange {
		// 拼接目录
		var Dir = m_str.Join(
			global.Path.DataPath,
			os.PathSeparator,
			exchange, // 统一采用 按照交易所分开填写
			os.PathSeparator,
			opt.Goods.GoodsId, // 统一采用 GoodsId 作为目录
			os.PathSeparator,
			opt.BarObj.Binance, // 统一采用小写作为目录
		)
		// 获得初始的 Before , 也就是 最大值
		var Before_original int64
		// 读取目录下的文件列表
		files, _ := os.ReadDir(Dir)
		if len(files) < 1 {
			// 目录下没有文件,则以 EndTime 作为Before_original
			Before_original = opt.EndTime
		} else {
			// 有文件则找到那个 最接近  opt.EndTime 的文件

			for _, file := range files {
				m_json.Println(file)
			}

		}

		// 计算最多遍历多少次 MaxLoop = Limit / 100（请求时的固定条目） + 2 （前后时间拢余都算上）
		var MaxLoop = 10 //  Limit 最大 500  , 所以遍历次数最大 10

		// 计算请求列表
		SendKlineRequestOptList := []SendKlineRequestOpt{}
		for i := 0; i < MaxLoop; i++ {
			var timeUnix = Before_original - opt.BarObj.Interval*int64(i)*100
			year := m_time.MsToTime(timeUnix, "0").Format("2006")
			var SendKlineRequestOpt = SendKlineRequestOpt{
				Before: timeUnix,
				Bar:    opt.BarObj.Binance, //内部会进行处理
				StoreFilePath: m_str.Join(
					Dir,
					os.PathSeparator,
					year,
					os.PathSeparator,
					timeUnix, ".json",
				),
			}
			// 币安
			if exchange == global.ExchangeOpt[0] {
				SendKlineRequestOpt.Binance_symbol = opt.Goods.BinanceInfo.Symbol
			}
			// 欧意
			if exchange == global.ExchangeOpt[1] {
				SendKlineRequestOpt.Okx_instId = opt.Goods.Okx_SPOT_Info.InstID
			}
			SendKlineRequestOptList = append(SendKlineRequestOptList, SendKlineRequestOpt)
			if timeUnix < opt.StartTime {
				break
			}
		}
		resData = append(resData, SendKlineRequestOptList...)
	}

	return
}

type SendKlineRequestOpt struct {
	Okx_instId     string `json:"Okx_instId"` // 和 Binance_symbol 二选一
	Binance_symbol string `json:"Binance_symbol"`
	Bar            string `json:"Bar"`
	Before         int64  `json:"Before"`
	StoreFilePath  string `json:"StoreFilePath"`
}

func SendKlineRequest(opt SendKlineRequestOpt) (resData []global.KlineSimpType, resErr error) {

	resData = nil
	resErr = nil

	// 先读取文件看看是否存在

	IsExist := m_path.IsExist(opt.StoreFilePath)
	if IsExist {
		// 存在该文件，则进行读取
		fileCont := m_file.ReadFile(opt.StoreFilePath)
		var kline []global.KlineSimpType
		err := json.Unmarshal(fileCont, &kline)
		if err != nil {
			global.LogErr("exchange_api.SendKlineRequest 文件解析失败,将重新获取并覆盖", opt.StoreFilePath)
		} else {
			if len(kline) == 100 { // 数据解析成功并返回
				resData = kline
				global.RunLog.Println("已从本地文件中返回数据", opt.StoreFilePath)
				return
			} else {
				global.LogErr("数据不完整，将重新获取并写入", opt.StoreFilePath)
			}
		}
	}

	if len(opt.Okx_instId) > 2 {
		fetchData, err := okx.GetKline(okx.GetKlineOpt{
			Okx_instId: opt.Okx_instId,
			Bar:        opt.Bar,
			Before:     opt.Before,
		})
		if err != nil {
			resErr = err
			return
		}
		resData = fetchData
	}

	if len(opt.Binance_symbol) > 2 {
		fetchData, err := binance.GetKline(binance.GetKlineOpt{
			Binance_symbol: opt.Binance_symbol,
			Bar:            opt.Bar,
			Before:         opt.Before,
		})
		if err != nil {
			resErr = err
			return
		}
		resData = fetchData
	}

	return
}
