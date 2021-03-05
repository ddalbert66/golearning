package main

import (
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
	var lastRate float64

	for _, st := range stockArray {
		if len(stockName) == 0 {
			stockName = st.StockCode
		} else if strings.Compare(stockName, st.StockCode) == 0 {

			if st.TurnoverRate > lastRate {
				count++
			}

			if count > 8 {
				if !strings.Contains(needStockCode, st.StockCode) {

					count = 0
					stockService.FuncgetGoodStock(st.StockCode)
					needStockCode = needStockCode + ":" + st.StockCode
				}
			}
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
