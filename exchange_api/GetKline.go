package exchange_api

import (
	"coin-candle/exchange_api/binance"
	"coin-candle/exchange_api/okx"
	"coin-candle/global"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

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

	// Limit 缺省值 10 条不能再多了
	Limit := global.KlineLimitDefault
	if opt.Limit > 0 && opt.Limit < global.KlineMaxLimit {
		Limit = opt.Limit
	} else {
		resErr = fmt.Errorf("参数 Limit 必须为 1~%+v 的正整数", global.KlineMaxLimit)
		return
	}

	// EndTime 缺省值
	now := m_time.GetUnixInt64()
	EndTime := now
	// 时间 传入的时间戳 必须大于最早时间才有效否则重置为当 now
	if opt.EndTime > global.TimeOldest {
		EndTime = opt.EndTime
	}
	// 计算起始时间  //  结束时间 - 时间间隔 * 条数
	StartTime := EndTime - BarObj.Interval*int64(Limit)

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

	SendParamList := GetKlineFilePath(GetKlineFilePathOpt{
		Limit:     Limit,
		StartTime: StartTime,
		EndTime:   EndTime,
		BarObj:    BarObj,
		Goods:     GoodsDetail,
		Exchange:  Exchange,
	})

	// 在这里进行K线的整合并返回
	var KlineMap = map[string][]global.KlineSimpType{}

	for _, item := range SendParamList {
		kline, err := SendKlineRequest(item)
		if err != nil {
			resErr = err
			return
		}
		// 将 kline 串起来
		KlineMap[item.Exchange] = append(kline, KlineMap[item.Exchange]...)
	}

	// var ExchangeMapList = []global.KlineExchangeMap{}
	for key, kline := range KlineMap {
		var list = []global.KlineSimpType{}
		for _, item := range kline {
			timeUnix, _ := strconv.ParseInt(item[0], 10, 64)
			// 过滤，挑选出符合规则的数据
			if timeUnix >= StartTime && timeUnix <= EndTime {
				list = append(list, item)
			}
		}

		fmt.Println(key, len(list))
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
			global.Path.DataPath, // 数据目录
			os.PathSeparator,
			exchange, // 按照交易所名称分开存储
			os.PathSeparator,
			opt.Goods.GoodsId, // 统一采用 GoodsId 作为目录
			os.PathSeparator,
			opt.BarObj.DirName, // 时间间隔
		)

		// 获得 file name 的时间颗粒度
		Before_Time := FindFileTime(opt)

		if Before_Time == 0 {
			global.LogErr("没有找到文件时间戳列表")
			return
		}

		// 计算最多遍历多少次 MaxLoop = Limit / 100  + 2 （前后时间拢余都算上）
		var MaxLoop = global.KlineMaxLimit/global.ExchangeKlineLimit + 2

		// 计算要发送的请求列表
		SendKlineRequestOptList := []SendKlineRequestOpt{}
		for i := 0; i < MaxLoop; i++ {
			var fileInterval = opt.BarObj.Interval * global.ExchangeKlineLimit
			var timeUnix = Before_Time - fileInterval*int64(i) // 最初的时间 挨个递减 条
			year_month := m_time.MsToTime(opt.EndTime, "0").Format("2006-01")

			var SendKlineRequestOpt = SendKlineRequestOpt{
				GoodsId:  opt.Goods.GoodsId,
				Exchange: exchange,
				EndTime:  timeUnix,   // 发出请求的时间
				BarObj:   opt.BarObj, // 发出请求的时间间隔
				StoreFilePath: m_str.Join( // 请求来的数据应当存放的目录
					Dir,
					os.PathSeparator,
					year_month, // 年月
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
			if timeUnix-fileInterval < opt.StartTime {
				break
			}
		}
		resData = append(resData, SendKlineRequestOptList...)
	}

	return
}

type SendKlineRequestOpt struct {
	GoodsId        string              `json:"GoodsId"`
	Exchange       string              `json:"Exchange"`   // 交易所
	Okx_instId     string              `json:"Okx_instId"` // 和 Binance_symbol 二选一
	Binance_symbol string              `json:"Binance_symbol"`
	BarObj         global.KlineBarType `json:"Bar"`
	EndTime        int64               `json:"EndTime"`
	StoreFilePath  string              `json:"StoreFilePath"`
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
			if len(kline) == global.ExchangeKlineLimit { //数据解析成功并返回
				resData = kline
				return
			} else {
				global.LogErr("数据不完整，将重新获取并写入", opt.StoreFilePath)
			}
		}
	}

	// 去交易所请求
	var fetchData []global.KlineSimpType
	var err error
	if len(opt.Okx_instId) > 2 {
		fetchData, err = okx.GetKline(okx.GetKlineOpt{
			Okx_instId: opt.Okx_instId,
			Bar:        opt.BarObj.Okx,
			EndTime:    opt.EndTime,
		})

	}

	if len(opt.Binance_symbol) > 2 {
		fetchData, err = binance.GetKline(binance.GetKlineOpt{
			Binance_symbol: opt.Binance_symbol,
			Bar:            opt.BarObj.Binance,
			EndTime:        opt.EndTime,
		})
	}

	if err != nil {
		resErr = err
		return
	}

	lastTime, err := strconv.ParseInt(fetchData[len(fetchData)-1][0], 10, 64)
	if err != nil {
		resErr = err
		return
	}

	// 如果存在未来时间,则进行摘除
	if opt.EndTime > lastTime {
		diffLimit := (opt.EndTime - lastTime) / opt.BarObj.Interval
		resData = fetchData[diffLimit:]
	} else {
		resData = fetchData
	}
	m_file.WriteByte(opt.StoreFilePath, m_json.ToJson(resData))
	return
}

// 获取目录时间戳列表 并排序
func FindFileTime(opt GetKlineFilePathOpt) (resData int64) {

	// 文件之间的时间间隔
	var fileInterval = opt.BarObj.Interval * global.ExchangeKlineLimit

	// 计算用户和基准时间之间的差值
	var maxFileTime int64
	if opt.EndTime > global.FileNameBaseTime {
		var diffLimit = (opt.EndTime-global.FileNameBaseTime)/fileInterval + 1
		maxFileTime = global.FileNameBaseTime + (diffLimit * fileInterval) //加文件间隔
	} else {
		var diffLimit = (global.FileNameBaseTime-opt.EndTime)/fileInterval + 1
		maxFileTime = global.FileNameBaseTime - (diffLimit * fileInterval) //减文件间隔
	}

	return maxFileTime
}
