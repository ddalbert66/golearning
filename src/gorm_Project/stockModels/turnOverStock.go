package stockModels

import (
	"github.com/jinzhu/gorm"
)

//排 行	股票代號	股票名稱	總成交股數	發行股數	週轉率(%)
//1	   3540	        曜越        16,904,464	66,613,044	25.37
type Stock struct {
	// Id        int64
	StockName    string
	StockCode    string
	StockVolume  int
	IssuedShares int
	TurnoverRate float64
	DateStr      string
	Create_user  string
	Update_user  string
	//gorm.Model 其包括字段 ID、CreatedAt、UpdatedAt、DeletedAt
	gorm.Model
}
