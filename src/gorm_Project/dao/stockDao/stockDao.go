package stockDao

import (
	"goLearning20200930/src/gorm_Project/stockModels"
)

type Stock stockModels.Stock

// func db() *gorm.DB {
// 	return dao.GetDB()
// }

func (s Stock) TableName() string {
	return "stocks"
}

// func QueryStockByName(name string) (s []Stock) {
// 	db().Order("date_str").Find(&s, "stock_name = ?", name)
// 	return
// }

func QueryStock() (s []Stock) {
	db().Order("stock_code,date_str").Find(&s)
	return
}
