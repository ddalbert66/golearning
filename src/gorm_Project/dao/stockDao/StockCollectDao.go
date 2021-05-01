package stockDao

import (
	"goLearning20200930/src/gorm_Project/dao"
	"goLearning20200930/src/gorm_Project/stockModels"

	"github.com/jinzhu/gorm"
)

type StockCollect stockModels.StockCollect

func db5() *gorm.DB {
	return dao.GetDB()
}

func (s StockCollect) TableName() string {
	return "Stock_Collects"
}
