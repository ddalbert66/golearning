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

	rows, err := db.Raw(`SELECT  a.*,st.stock_Type from  (
		select s1.stock_name,s1.stock_code,s1.turnover_rate s1_Rate,s2.turnover_rate s2_Rate,s1.stock_volume,s1.closing_price,s1.date_str,((s1.closing_price-s2.closing_price)/s1.closing_price)*100 closeing_Prince_Rate 
		from bento.stocks s1 left join 
		(select stock_code,turnover_rate,closing_price from bento.stocks s where date_str = ? ) s2 on 
		s2.stock_code = s1.stock_code 
		where   s1.date_str = ? and  s1.turnover_rate>s2.turnover_rate and s1.turnover_rate >=2 and s1.turnover_rate < 10  and  s2.turnover_rate < 1
		) a left join bento.stock_types  st on a.stock_code = st.stock_code where a.closeing_Prince_Rate >0  order by st.stock_Type,closeing_Prince_Rate desc`, dateStrS2, dateStrS1).Rows()

	defer rows.Close()

	if err != nil {
		log.Panic(err)
	}
	// fmt.Println(mapToJson(stockCollects))
	// for _, sc := range stockCollects {
	// 	fmt.Printf("%f %f", sc.S1Rate, sc.S2Rate)

	// }
	for rows.Next() {
		var stockCollect StockCollect
		//scan 代表select 參數對應
		err = rows.Scan(&stockCollect.StockName, &stockCollect.StockCode, &stockCollect.S1Rate,
			&stockCollect.S2Rate, &stockCollect.StockVolume, &stockCollect.ClosingPrice, &stockCollect.DateStr, &stockCollect.CloseingPrinceRate, &stockCollect.StockType)

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
