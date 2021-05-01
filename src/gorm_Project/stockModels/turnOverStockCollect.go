package stockModels

import (
	"github.com/jinzhu/gorm"
)

//名次	股票名稱	收盤價	   漲跌	      漲跌幅	成交量	週轉率
//1	   8021 尖點	32.20	 + 2.60	     +8.78%	   38,222	 26.88%
type StockCollect struct {
	// Id        int64
	StockName          string
	StockCode          string
	S1Rate             float64
	S2Rate             float64
	StockVolume        int
	ClosingPrice       float64
	DateStr            string
	CloseingPrinceRate float64
	StockType          string
	Create_user        string
	Update_user        string
	//gorm.Model 其包括字段 ID、CreatedAt、UpdatedAt、DeletedAt
	gorm.Model
}
