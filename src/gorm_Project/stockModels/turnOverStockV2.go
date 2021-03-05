package stockModels

import (
	"github.com/jinzhu/gorm"
)

//名次	股票名稱	收盤價	   漲跌	      漲跌幅	成交量	週轉率
//1	   8021 尖點	32.20	 + 2.60	     +8.78%	   38,222	 26.88%
type StockV2 struct {
	// Id        int64
	StockName     string
	StockCode     string
	ClosingPrice  float64
	UpAndDown     string
	UpAndDownRate string
	StockVolume   int
	TurnoverRate  float64
	DateStr       string
	Create_user   string
	Update_user   string
	//gorm.Model 其包括字段 ID、CreatedAt、UpdatedAt、DeletedAt
	gorm.Model
}
