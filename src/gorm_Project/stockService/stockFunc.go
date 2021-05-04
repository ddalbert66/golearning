package stockService

import (
	"encoding/json"
	"fmt"
	"goLearning20200930/src/gorm_Project/stockModels"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var ClosingPriceMap = make(map[string][]float64, 21) 

var ClosingPriceArr = make([]float64, 0) 

func mapToJson(result interface{}) string {
	// map转 json str
	jsonBytes, _ := json.Marshal(result)
	jsonStr := string(jsonBytes)
	return jsonStr
}

type Stock stockModels.Stock

func init() {
	db, err := gorm.Open("mysql", "root:abcd1234@tcp(127.0.0.1:3306)/bento?charset=utf8&loc=Local&parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//var stockCollects []stockModels.StockCollect

	// dateStrS2 := "110/04/28"
	// dateStrS1 := "110/04/29"

	rows, err := db.Raw(`SELECT Stock_Name,Stock_Code,Stock_Volume,Turnover_Rate,Closing_Price,Date_Str FROM bento.stock_v2 order by stock_code,date_str desc`).Rows()
	//rows, err := db.Raw(`SELECT Stock_Name,Stock_Code,Stock_Volume,Turnover_Rate,Closing_Price,Date_Str FROM bento.stocks order by stock_code,date_str desc`).Rows()

	defer rows.Close()

	if err != nil {
		log.Panic(err)
	}
	// fmt.Println(mapToJson(stockCollects))
	// for _, sc := range stockCollects {
	// 	fmt.Printf("%f %f", sc.S1Rate, sc.S2Rate)

	// }
	for rows.Next() {
		var stock Stock
		//scan 代表select 參數對應
		err = rows.Scan(&stock.StockName, &stock.StockCode, &stock.StockVolume, &stock.TurnoverRate, &stock.ClosingPrice, &stock.DateStr)
		writeClosingPrice(stock)
	
		if err != nil {
			fmt.Print(err)
		}
		// InsertStockCollect(stockCollect)
	}

	// for k,v := range ClosingPriceMap{
	// 	OneRedThreeBlack(k,v)
	// }
}

// func db3() *gorm.DB {
// 	return dao.GetDB()
// }

// func InsertStockCollect(s StockCollect) uint {
// 	result := db3().Create(&s)
// 	if result.Error != nil {
// 		panic(result.Error)
// 	}
// 	return s.ID
// }

func writeClosingPrice(s Stock) {


	if _, ok := ClosingPriceMap[s.StockCode]; ok {
		if len(ClosingPriceArr)>21 {
			return
		} 
		//fmt.Println(result, ok)
		fmt.Println("股票單號:",s.StockCode,",收盤價:",s.ClosingPrice)
		ClosingPriceArr = append(ClosingPriceArr,s.ClosingPrice)
		fmt.Println(ClosingPriceArr)
		ClosingPriceMap[s.StockCode] = ClosingPriceArr
	} else {
		
		ClosingPriceArr = make([]float64, 0) 
		ClosingPriceArr = append(ClosingPriceArr,s.ClosingPrice)
		ClosingPriceMap[s.StockCode] = ClosingPriceArr
		//fmt.Println("without key")
	}
}



func OneRedThreeBlack(stockCode string) string{
	fmt.Println("oneRedThreeBlack stockCode:",stockCode)
	result :=""
	cp := ClosingPriceMap[stockCode]
	if cp[0]>cp[1] && cp[1]>cp[2] && cp[0]<cp[3] && cp[1]<cp[3] && cp[2]<cp[3] {
		fmt.Println("oneRedThreeBlack !!!! stockCode:",stockCode,",closingPrice:",cp)
		result = "一紅吃三黑"
	}
	return result
	
}



func MA5(stockCode string) float64{
	fmt.Println("oneRedThreeBlack stockCode:",stockCode)

	var total5day float64 =0
	cp := ClosingPriceMap[stockCode]
	fmt.Println(cp) 
	for k,closingPrice := range cp {
		if k == 5 {
			break
		} 
		total5day = total5day + closingPrice
	}
	fmt.Println(total5day) 
	ma5 := total5day/5
	fmt.Println("五日線:",ma5)
	return ma5
	
}

func MA10(stockCode string) float64{
	fmt.Println("oneRedThreeBlack stockCode:",stockCode)

	var total10day float64 =0
	cp := ClosingPriceMap[stockCode]
	fmt.Println(cp) 
	for k,closingPrice := range cp {
		if k == 10 {
			break
		} 
		total10day = total10day + closingPrice
	}
	fmt.Println(total10day) 
	ma10 := total10day/10
	fmt.Println("十日線:",ma10)
	return ma10
	
}

func MA20(stockCode string) float64{
	fmt.Println("MA20 stockCode:",stockCode)

	var total20day float64 =0
	cp := ClosingPriceMap[stockCode]
	fmt.Println(cp) 
	for k,closingPrice := range cp {
		if k == 20 {
			break
		} 
		total20day = total20day + closingPrice
	}
	fmt.Println(total20day) 
	ma20 := total20day/20
	fmt.Println("二十日線:",ma20)
	return ma20
	
}