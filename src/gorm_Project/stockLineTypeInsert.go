package main

import (
	"fmt"
	"goLearning20200930/src/gorm_Project/dao"
	"goLearning20200930/src/gorm_Project/dao/stockDao"
	"goLearning20200930/src/gorm_Project/stockService"

	"github.com/jinzhu/gorm"
)

//type Stock stockModels.Stock

func db() *gorm.DB {
	return dao.GetDB()
}

// func Update(stock stockDao.Stock) {
// 	stock.Update()
// }

func UpdateV2(stock stockDao.StockV2) {
	stock.Update()
}

/*
*	周轉率是否有持續增加
*
 */
func main() {
	dateStr := "110/05/11"
	stockArray2 := stockDao.QueryStockV2ByDate(dateStr)
	//stockArray2 := stockDao.QueryStockByDateStr(dateStr)
	for _, st := range stockArray2 {
		if len(st.StockCode) == 4 {
			stockv2 := stockDao.QueryStockV2ByCodeAndDate(st.StockCode, dateStr)
			fmt.Println(stockv2.StockCode)
			ma5 := stockService.MA5(stockv2.StockCode)
			stockv2.Ma5 = ma5

			ma10 := stockService.MA10(stockv2.StockCode)
			stockv2.Ma10 = ma10

			ma20 := stockService.MA20(stockv2.StockCode)
			stockv2.Ma20 = ma20
			UpdateV2(stockv2)
		}

	}

	stockArray := stockDao.QueryStockV2ByDate(dateStr)
	for _, st := range stockArray {
		if len(st.StockCode) == 4 {
			stock := stockDao.QueryStockV2ByCodeAndDate(st.StockCode, dateStr)
			fmt.Println(stock.StockCode, stock.Ma5, stock.Ma10, stock.Ma20)
			if stock.Ma5 != 0 {
				lineType := stockService.MovingAverageTangled(stock.StockCode)
				fmt.Println(lineType)
				stock.LineType = lineType
				UpdateV2(stock)
			} else {
				fmt.Println("沒有均線資料:", stock.StockCode)
			}
		}
	}

}
