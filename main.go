package main

import (
	"fmt"
	"github.com/idoall/TokenExchangeCommon/commonstock"
)

func main() {
	list := commonstock.InitStockKline()

	//计算新的MACD
	//stockList := commonstock.NewMACD(list).Calculation().GetPoints()
	//
	//for _, v := range stockList {
	//	fmt.Printf("Time:%s\t DIF:%f DEA:%f MACD %f\n", v.Time.Format("2006-01-02 15:04:05"), v.DIF, v.DEA, v.MACD)
	//}

	//计算新的KDJ
	stockList := commonstock.NewKDJ(list, 2).Calculation().GetPoints()
	for _, v := range stockList {
		fmt.Printf("Time:%s\t K:%.5f D:%.5f J:%.5f\n", v.Time.Format("2006-01-02 15:04:05"), v.K, v.D, v.J)
	}
}
