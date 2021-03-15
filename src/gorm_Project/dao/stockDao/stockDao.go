package stockDao

import (
	"goLearning20200930/src/gorm_Project/dao"
	"goLearning20200930/src/gorm_Project/stockModels"

	"github.com/jinzhu/gorm"
)

type Stock stockModels.Stock

func db2() *gorm.DB {
	return dao.GetDB()
}

func (s Stock) TableName() string {
	return "stocks"
}

// func QueryStockByName(name string) (s []Stock) {
// 	db().Order("date_str").Find(&s, "stock_name = ?", name)
// 	return
// }

func QueryStockByDateStr(dateStr string) (s []Stock) {
	db2().Find(&s, "date_str = ?", dateStr)
	return
}

func QueryStock() (s []Stock) {
	db2().Order("stock_code,date_str").Find(&s)
	return
}

func Insert(s Stock) {
	result := db2().Create(&s)
	if result.Error != nil {
		panic(result.Error)
	}
}

func QueryStockByCodeAndDate(code string, dateStr string) (s Stock) {
	db().First(&s, "stock_code = ? and date_str = ?", code, dateStr)
	return
}

func (s Stock) Update() {
	db().Save(&s)
}
