package stockDao

import (
	"goLearning20200930/src/gorm_Project/dao"
	"goLearning20200930/src/gorm_Project/stockModels"

	"github.com/jinzhu/gorm"
)

type StockV2 stockModels.StockV2

func db() *gorm.DB {
	return dao.GetDB()
}

func (s StockV2) TableName() string {
	return "Stock_v2"
}

func QueryStockByCode(code string) (s []Stock) {
	db().Order("date_str").Find(&s, "stock_code = ?", code)
	return
}
