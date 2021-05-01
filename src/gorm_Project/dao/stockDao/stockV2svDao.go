package stockDao

import (
	"goLearning20200930/src/gorm_Project/stockModels"
)

type StockV2sv stockModels.StockV2StockVolume

func QueryStockV2svByCode(code string) (s StockV2sv) {
	db().Table("stock_v2_stock_volumes").Find(&s, "stock_code = ?", code)
	return
}
