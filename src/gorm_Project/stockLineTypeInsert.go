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


func Update(stock stockDao.Stock) {
	stock.Update()
}


func UpdateV2(stock stockDao.StockV2) {
	stock.Update()
}


/*
*	周轉率是否有持續增加
*
 */
func main() {
	dateStr := "110/05/04"
	//stockArray := stockDao.QueryStockByName("勤凱")
	//stockArray := stockDao.QueryStock()
	// for _,st := range stockArray {
	// 	stock := stockDao.QueryStockByCodeAndDate(st.StockCode, dateStr)
	// 	fmt.Println(st.StockCode)
	// 	lineType := stockService.OneRedThreeBlack(stock.StockCode)
	// 	fmt.Println(lineType)
	// 	stock.LineType = lineType
	// 	Update(stock)
	// }

	stockArray := stockDao.QueryStockV2()
	for _,st := range stockArray {
		if len(st.StockCode) == 4{
			stockv2 := stockDao.QueryStockV2ByCodeAndDate(st.StockCode, dateStr)
			fmt.Println(st.StockCode)
			//lineType := stockService.OneRedThreeBlack(stockv2.StockCode)
			//fmt.Println(lineType)
			//stockv2.LineType = lineType
			ma5 := stockService.MA5(stockv2.StockCode)
			stockv2.MA5 = ma5
			UpdateV2(stockv2)
		}
		
	}

}
