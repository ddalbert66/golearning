package main

import (
	"encoding/json"
	"fmt"
	"goLearning20200930/src/gorm_Project/dao"
	"goLearning20200930/src/gorm_Project/stockModels"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func mapToJson(result interface{}) string {
	// map转 json str
	jsonBytes, _ := json.Marshal(result)
	jsonStr := string(jsonBytes)
	return jsonStr
}

type StockCollect stockModels.StockCollect

func main() {
	db, err := gorm.Open("mysql", "root:abcd1234@tcp(127.0.0.1:3306)/bento?charset=utf8&loc=Local&parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//var stockCollects []stockModels.StockCollect

	dateStrS2 := "110/05/10"
	dateStrS1 := "110/05/11"

	rows, err := db.Raw(`select a.* from  (
		select s1.stock_name,s1.stock_code,s1.turnover_rate2 s1_Rate,s2.turnover_rate2 s2_Rate,floor(s1.shares_traded/1000) stock_Volume,s1.closing_price,s1.date_str,((s1.closing_price-s2.closing_price)/s1.closing_price)*100 closeing_Prince_Rate,s1.stock_type from
		 (select v2.stock_name,v2.stock_code,v2.closing_price,v2sv.stock_volume,v2.shares_traded,v2.shares_num,v2.date_str,(v2.shares_traded/v2sv.stock_volume)*100 turnover_rate2,v2sv.stock_type from bento.stock_v2 v2 join bento.stock_v2_stock_volumes v2sv  on v2.stock_code = v2sv.stock_code where v2.date_str = ?) s1 join 
		(select v2.stock_name,v2.stock_code,v2.closing_price,v2sv.stock_volume,v2.shares_traded,v2.shares_num,v2.date_str,(v2.shares_traded/v2sv.stock_volume)*100 turnover_rate2,v2sv.stock_type from bento.stock_v2 v2 join bento.stock_v2_stock_volumes v2sv  on v2.stock_code = v2sv.stock_code where v2.date_str = ? ) s2 on 
		s2.stock_code = s1.stock_code 
		where   s1.turnover_rate2>s2.turnover_rate2 and s1.turnover_rate2 >=2 and s1.turnover_rate2 < 10  and  s2.turnover_rate2 < 1 
		) a where a.closeing_Prince_Rate >0 order by closeing_Prince_Rate desc`, dateStrS1, dateStrS2).Rows()

	defer rows.Close()

	if err != nil {
		log.Panic(err)
	}

	for rows.Next() {
		var stockCollect StockCollect
		//scan 代表select 參數對應
		err = rows.Scan(&stockCollect.StockName, &stockCollect.StockCode, &stockCollect.S1Rate,
			&stockCollect.S2Rate, &stockCollect.StockVolume, &stockCollect.ClosingPrice, &stockCollect.DateStr, &stockCollect.CloseingPrinceRate, &stockCollect.StockType)
		//fmt.Printf("%f %f", stockCollect.S1Rate, stockCollect.S2Rate)

		if err != nil {
			fmt.Print(err)
		}
		InsertStockCollect(stockCollect)
	}
}

func db3() *gorm.DB {
	return dao.GetDB()
}

func InsertStockCollect(s StockCollect) uint {
	result := db3().Create(&s)
	if result.Error != nil {
		panic(result.Error)
	}
	return s.ID
}
