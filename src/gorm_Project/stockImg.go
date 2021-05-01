package main

import (
	"fmt"
	"goLearning20200930/src/gorm_Project/dao"
	"goLearning20200930/src/gorm_Project/dao/stockDao"
	"goLearning20200930/src/gorm_Project/stockService"

	"strings"

	"github.com/jinzhu/gorm"
)

//type Stock stockModels.Stock

func db() *gorm.DB {
	return dao.GetDB()
}

func main() {

	//stockArray := stockDao.QueryStockByName("勤凱")
	stockArray := stockDao.QueryStock()
	count := 0

	var needStockCode string
	var stockName string
	//var lastRate float64
	m := make(map[string][]float64)

	for _, st := range stockArray {

		if len(stockName) == 0 {
			stockName = st.StockCode
		} else if strings.Compare(stockName, st.StockCode) == 0 && len(stockName) == 4 {
			m[stockName] = append(m[stockName], st.TurnoverRate)
			fmt.Printf("stockName : %s ,TurnoverRate :% f \n", stockName, st.TurnoverRate)
			if len(m[stockName]) > 24 {
				fmt.Printf("count :%d, x[24]:%f,x[23]:%f ,StockCode :% s \n", count, m[stockName][15], m[stockName][14], st.StockCode)
				if count == 24 && m[stockName][24] > m[stockName][23] {
					fmt.Printf("needStockCode : %s ,StockCode :% s \n", needStockCode, st.StockCode)
					if !strings.Contains(needStockCode, st.StockCode) {

						count = 0
						stockService.FuncgetGoodStock(st.StockCode)
						needStockCode = needStockCode + ":" + st.StockCode
					}
				}
			}

			count++
		} else if stockName != st.StockCode {
			stockName = st.StockCode
			count = 0
		}
	}

	defer db().Close()
	// p, _ := plot.New()

	// p.Title.Text = "勤凱"
	// p.X.Label.Text = "date"
	// p.Y.Label.Text = "rate"

	// var plotterXys plotter.XYs
	// for i, st := range stockArray {
	// 	x := plotter.XY{float64(i), st.TurnoverRate}
	// 	plotterXys = append(plotterXys, x)
	// }

	// plotutil.AddLinePoints(p, plotterXys)

	// p.Save(4*vg.Inch, 4*vg.Inch, "price.png")
}
