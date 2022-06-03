package commonstock

import (
	"fmt"
	"testing"

	"github.com/idoall/TokenExchangeCommon/commonmodels"
)

func TestOBV(t *testing.T) {
	t.Parallel()
	list := InitTestKline()
	//计算新的OBV
	stock := NewOBV(list)
	stock.Calculation()
	obvList := stock.GetPoints()

	//计算MAOBV30
	var maObvArray []*commonmodels.Kline
	for _, v := range obvList {
		maObvArray = append(maObvArray, &commonmodels.Kline{
			Close:     v.point.Value,
			KlineTime: v.point.Time,
		})
	}
	maOBV := NewMA(maObvArray, 30)
	maOBVPoints := maOBV.Calculation().GetPoints()
	for i := 0; i < len(maOBVPoints); i++ {
		fmt.Printf("[%d]Time:%s\t Value:%f	MAOBV(30):%f\n", i, obvList[i].Time.Format("2006-01-02 15:04:05"), obvList[i].Value, maOBVPoints[i].Value)
	}
}
