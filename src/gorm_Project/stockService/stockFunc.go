package stockService

import (
	"encoding/json"
	"fmt"
	"goLearning20200930/src/gorm_Project/stockModels"
	"log"
	"math"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var ClosingPriceMap = make(map[string][]float64, 21)
var ClosingPriceArr = make([]float64, 0)
var StockMap = make(map[string][]Stock, 21)
var StockArr = make([]Stock, 0)

func mapToJson(result interface{}) string {
	// map转 json str
	jsonBytes, _ := json.Marshal(result)
	jsonStr := string(jsonBytes)
	return jsonStr
}

//type StockV2 stockModels.StockV2

type Stock stockModels.StockV2

func init() {
	db, err := gorm.Open("mysql", "root:abcd1234@tcp(127.0.0.1:3306)/bento?charset=utf8&loc=Local&parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//var stockCollects []stockModels.StockCollect

	// dateStrS2 := "110/04/28"
	// dateStrS1 := "110/04/29"

	rows, err := db.Raw(`SELECT Stock_Name,Stock_Code,Stock_Volume,Turnover_Rate,Closing_Price,ma5,ma10,ma20,Date_Str FROM bento.stock_v2 order by stock_code,date_str desc`).Rows()
	//rows, err := db.Raw(`SELECT Stock_Name,Stock_Code,Stock_Volume,Turnover_Rate,Closing_Price,ma5,ma10,ma20,Date_Str FROM bento.stocks order by stock_code,date_str desc`).Rows()

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
		err = rows.Scan(&stock.StockName, &stock.StockCode, &stock.StockVolume, &stock.TurnoverRate, &stock.ClosingPrice, &stock.Ma5, &stock.Ma10, &stock.Ma20, &stock.DateStr)
		writeClosingPrice(stock)
		writStock(stock)

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
func writStock(s Stock) {

	if _, ok := StockMap[s.StockCode]; ok {
		if len(StockArr) > 5 {
			return
		}
		//fmt.Println(result, ok)
		fmt.Println("股票單號:", s.StockCode, ",收盤價:", s.ClosingPrice)
		StockArr = append(StockArr, s)
		fmt.Println(StockArr)
		StockMap[s.StockCode] = StockArr
	} else {

		StockArr = make([]Stock, 0)
		StockArr = append(StockArr, s)
		StockMap[s.StockCode] = StockArr
		//fmt.Println("without key")
	}
}

func writeClosingPrice(s Stock) {

	if _, ok := ClosingPriceMap[s.StockCode]; ok {
		if len(ClosingPriceArr) > 21 {
			return
		}
		//fmt.Println(result, ok)
		fmt.Println("股票單號:", s.StockCode, ",收盤價:", s.ClosingPrice)
		ClosingPriceArr = append(ClosingPriceArr, s.ClosingPrice)
		fmt.Println(ClosingPriceArr)
		ClosingPriceMap[s.StockCode] = ClosingPriceArr
	} else {

		ClosingPriceArr = make([]float64, 0)
		ClosingPriceArr = append(ClosingPriceArr, s.ClosingPrice)
		ClosingPriceMap[s.StockCode] = ClosingPriceArr
		//fmt.Println("without key")
	}
}

func MovingAverageTangled(stockCode string) string {
	fmt.Println("movingAverageTangled stockCode:", stockCode)
	result := ""
	cp := StockMap[stockCode]

	if checkRange5Percent(cp[0]) && checkRange5Percent(cp[1]) && checkRange5Percent(cp[2]) && checkRange5Percent(cp[3]) && checkRange5Percent(cp[4]) {
		fmt.Println("movingAverageTangled !!!! stockCode:", stockCode)
		result = "均線糾結"
	}
	return result

}

func checkRange5Percent(s Stock) bool {
	fmt.Println(s.Ma5, s.Ma10, s.Ma20)
	rangCMa5 := (math.Abs(s.ClosingPrice-s.Ma5) / s.ClosingPrice) * 100
	rangCMa10 := (math.Abs(s.ClosingPrice-s.Ma10) / s.ClosingPrice) * 100
	rangCMa20 := (math.Abs(s.ClosingPrice-s.Ma20) / s.ClosingPrice) * 100

	rangMa10Ma5 := (math.Abs(s.Ma10-s.Ma5) / s.Ma10) * 100
	rangMa20Ma10 := (math.Abs(s.Ma20-s.Ma10) / s.Ma20) * 100
	rangMa20Ma5 := (math.Abs(s.Ma20-s.Ma5) / s.Ma20) * 100
	fmt.Println(rangCMa5, rangCMa10, rangCMa20, rangMa10Ma5, rangMa20Ma10, rangMa20Ma5)
	if rangCMa5 <= 0.5 && rangCMa10 <= 0.5 && rangCMa20 <= 0.5 && rangMa10Ma5 <= 0.5 && rangMa20Ma10 <= 0.5 && rangMa20Ma5 <= 0.5 {
		return true
	} else {
		return false
	}

}

func OneRedThreeBlack(stockCode string) string {
	fmt.Println("oneRedThreeBlack stockCode:", stockCode)
	result := ""
	cp := ClosingPriceMap[stockCode]
	if cp[0] > cp[1] && cp[1] > cp[2] && cp[0] < cp[3] && cp[1] < cp[3] && cp[2] < cp[3] {
		fmt.Println("oneRedThreeBlack !!!! stockCode:", stockCode, ",closingPrice:", cp)
		result = "一紅吃三黑"
	}
	return result

}

func MA5(stockCode string) float64 {
	fmt.Println("oneRedThreeBlack stockCode:", stockCode)

	var total5day float64 = 0
	cp := ClosingPriceMap[stockCode]
	fmt.Println(cp)
	for k, closingPrice := range cp {
		if k == 5 {
			break
		}
		total5day = total5day + closingPrice
	}
	fmt.Println(total5day)
	ma5 := total5day / 5
	fmt.Println("五日線:", ma5)
	return ma5

}

func MA10(stockCode string) float64 {
	fmt.Println("oneRedThreeBlack stockCode:", stockCode)

	var total10day float64 = 0
	cp := ClosingPriceMap[stockCode]
	fmt.Println(cp)
	for k, closingPrice := range cp {
		if k == 10 {
			break
		}
		total10day = total10day + closingPrice
	}
	fmt.Println(total10day)
	ma10 := total10day / 10
	fmt.Println("十日線:", ma10)
	return ma10

}

func MA20(stockCode string) float64 {
	fmt.Println("MA20 stockCode:", stockCode)

	var total20day float64 = 0
	cp := ClosingPriceMap[stockCode]
	fmt.Println(cp)
	for k, closingPrice := range cp {
		if k == 20 {
			break
		}
		total20day = total20day + closingPrice
	}
	fmt.Println(total20day)
	ma20 := total20day / 20
	fmt.Println("二十日線:", ma20)
	return ma20

}
