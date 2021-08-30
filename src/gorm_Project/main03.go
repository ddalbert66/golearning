package main

import (
	"encoding/json"
	"fmt"
	"goLearning20200930/src/gorm_Project/dao"
	"goLearning20200930/src/gorm_Project/stockModels"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type StockV2 stockModels.StockV2

func db3() *gorm.DB {
	return dao.GetDB()
}

func InsertV2(s StockV2) uint {
	result := db3().Create(&s)
	if result.Error != nil {
		panic(result.Error)
	}
	return s.ID
}

//寫入上市股票的 資料 不寫入周轉排行
//周轉排行 要自己查SQL的時候去找發行股票計算
func main() {
	dateStr := "110/07/22"
	year, _ := strconv.Atoi(strings.Split(dateStr, "/")[0])
	//fmt.Print(year)
	year = year + 1911
	mon := strings.Split(dateStr, "/")[1]
	date := strings.Split(dateStr, "/")[2]
	newDateStr := strconv.Itoa(year) + mon + date

	url := "https://www.twse.com.tw/exchangeReport/MI_INDEX?response=json&date=" + newDateStr + "&type=ALLBUT0999"
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}

	//defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var jsonObj map[string]interface{}
	json.Unmarshal([]byte(body), &jsonObj)

	for k, v := range jsonObj {
		if k == "data9" {
			praseStockData(v, dateStr)
			//fmt.Println(v)
		}

	}

	fmt.Println("end day:", dateStr)

}

func praseStockData(v interface{}, dateStr string) {
	var jsonObj2 []interface{}
	b, err := json.Marshal(v)
	if err != nil {
		//fmt.Println("---")
		panic(err)
	}
	//var result string
	//fmt.Println(b)

	json.Unmarshal(b, &jsonObj2)
	//fmt.Println(jsonObj2)

	for _, vs := range jsonObj2 {
		var jsonObj3 []interface{}
		b2, err := json.Marshal(vs)
		if err != nil {
			//.Println("---")
			panic(err)
		}
		//var result string
		//fmt.Println(b)

		json.Unmarshal(b2, &jsonObj3)
		var stock StockV2
		stock.DateStr = dateStr
		for i, v2 := range jsonObj3 {

			switch i % 18 {
			case 0:
				//fmt.Print(v2)
				//fmt.Print(" ")
				stock.StockCode = v2.(string)
			case 1:
				//fmt.Print(v2)
				//fmt.Print(" ")
				stock.StockName = v2.(string)
			case 2:
				//fmt.Print(v2)
				//fmt.Print(" ")
				tempValue, _ := strconv.Atoi(strings.Replace(v2.(string), ",", "", -1))
				stock.SharesTraded = tempValue
			case 3:
				//fmt.Print(v2)
				//fmt.Print(" ")
				stock.SharesNum, _ = strconv.Atoi(strings.Replace(v2.(string), ",", "", -1))
			case 8:
				//fmt.Print(v2)
				//fmt.Print(" ")
				stock.ClosingPrice, _ = strconv.ParseFloat(v2.(string), 64)
			case 15:
				//fmt.Println(v2)
				stock.PEratio, _ = strconv.ParseFloat(v2.(string), 64)
				if i > 0 {
					InsertV2(stock)
				}
			}
		}

	}
}
