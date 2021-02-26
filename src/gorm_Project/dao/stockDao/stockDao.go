package stockDao

import (
	"goLearning20200930/src/gorm_Project/dao"
	"goLearning20200930/src/gorm_Project/stockModels"

	"github.com/jinzhu/gorm"
)

type Stock stockModels.Stock

func db() *gorm.DB {
	return dao.GetDB()
}

func (s Stock) TableName() string {
	return "stocks"
}

func QueryStockByName(name string) (s []Stock) {
	db().Find(&s, "stock_name = ?", name).Order("date_str asc")
	return
}
