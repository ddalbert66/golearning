package main

import (
	"fmt"
	"goLearning20200930/src/gorm_Project/dao/stockDao"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//type StockOral Stock

//Stock 櫃買市場 資料庫
//依照股票代號 寫入收盤價
//URL 日期要記得改
func main() {
	url := "https://www.tpex.org.tw/web/stock/aftertrading/daily_close_quotes/stk_quote_result.php?l=zh-tw&o=htm&d=110/07/22&s=0,asc,0"
	urlarray := strings.Split(url, "&")
	dateStr := strings.Split(urlarray[2], "=")[1]

	c := colly.NewCollector()

	count := 0

	// if QueryStockByDate2Str(dateStr) > 0 {
	// 	fmt.Print("已經擁有")
	// } else {
	var stock stockDao.Stock
	// 當Visit訪問網頁後，在網頁響應(Response)之後、發現這是HTML格式 執行的事情
	c.OnHTML("td", func(e *colly.HTMLElement) { // 每找到一個符合 goquerySelector字樣的結果，便會進這個OnHTML一次
		//fmt.Println(e.Text)

		switch count % 19 {

		case 1:
			if len(e.Text) == 4 {
				stock = stockDao.QueryStockByCodeAndDate(e.Text, dateStr)
			}
			//fmt.Printf("排 行:%s", e.Text)
			//fmt.Print(stock)
			//stock.StockCode = e.Text
			//fmt.Printf("股票代號:%s", e.Text)
		case 2:
			//stock.StockName = e.Text
			//fmt.Printf("股票名稱:%s", e.Text)

			//stockService.QueryStockCloserPrice(dateStr, stock.StockCode)
		case 3:
			var err error
			stock.ClosingPrice, err = strconv.ParseFloat(e.Text, 64)
			if err == nil {
				Update(stock)
			} else {
				stock.ClosingPrice = 0
			}
			//fmt.Printf("收盤價:%s \n", e.Text)

		}

		count++

	})

	//defer db().Close()
	c.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	c.Visit(url)
	// }
	fmt.Print("end day:", dateStr)

}

func Update(stock stockDao.Stock) {
	stock.Update()
}

// func praseStockData(v interface{}, dateStr string) {
// 	var jsonObj2 []interface{}
// 	b, err := json.Marshal(v)
// 	if err != nil {
// 		fmt.Println("---")
// 		panic(err)
// 	}
// 	//var result string
// 	//fmt.Println(b)

// 	json.Unmarshal(b, &jsonObj2)
// 	//fmt.Println(jsonObj2)

// 	for _, vs := range jsonObj2 {
// 		var jsonObj3 []interface{}
// 		b2, err := json.Marshal(vs)
// 		if err != nil {
// 			fmt.Println("---")
// 			panic(err)
// 		}
// 		//var result string
// 		//fmt.Println(b)

// 		json.Unmarshal(b2, &jsonObj3)
// 		var stock stockDao.Stock
// 		stock.DateStr = dateStr

// 		for i, v2 := range jsonObj3 {

// 			switch i % 18 {
// 			case 0:
// 				//fmt.Print(v2)
// 				//fmt.Print(" ")
// 				//stock.StockCode = v2.(string)
// 				stock = stockDao.QueryStockByCodeAndDate(v2.(string), dateStr)
// 				if stock.ID == 0 {

// 				} else {
// 					fmt.Print(v2.(string))
// 					fmt.Println("查得到")
// 				}
// 			case 1:
// 				//fmt.Print(v2)
// 				//fmt.Print(" ")
// 				//stock.StockName = v2.(string)
// 			case 8:
// 				//fmt.Print(v2)
// 				//fmt.Print(" ")
// 				stock.ClosingPrice, _ = strconv.ParseFloat(v2.(string), 64)
// 			case 15:
// 				//fmt.Println(v2)
// 				stock.PEratio, _ = strconv.ParseFloat(v2.(string), 64)
// 				if i > 0 {
// 					//Update(stock)
// 				}
// 			}
// 		}

// 	}
// }
