package main

import (
	"encoding/json"
	"fmt"
	"goLearning20200930/src/gorm_Project/dao"
	"goLearning20200930/src/gorm_Project/stockModels"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type StockV2 stockModels.StockV2

func db() *gorm.DB {
	return dao.GetDB()
}

func Insert(s StockV2) uint {
	result := db().Create(&s)
	if result.Error != nil {
		panic(result.Error)
	}
	return s.ID
}

//名次	股票名稱	收盤價	漲跌	漲跌幅	成交量	週轉率
//1	   8021  尖點	 32.20	+ 2.60	 +8.78%	 38,222	 26.88%
func main() {
	url := "https://www.wantgoo.com/stock/all-turnover-rates"
	var stock StockV2
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}

	var maptest []map[string]interface{}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &maptest)

	for _, v := range maptest {
		fmt.Print(v["investrueId"])

		stock.StockCode = v["investrueId"].(string)
		stock.TurnoverRate = v["value"].(float64)
		stock.DateStr = "2021/03/03"
		Insert(stock)
	}

}
