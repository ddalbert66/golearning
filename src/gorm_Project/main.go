package main

import (
	"fmt"
	"goLearning20200930/src/gorm_Project/dao"
	"goLearning20200930/src/gorm_Project/dao/stockDao"
	"goLearning20200930/src/gorm_Project/stockModels"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Stock stockModels.Stock

func db() *gorm.DB {
	return dao.GetDB()
}

func Insert(s Stock) {
	result := db().Create(&s)
	if result.Error != nil {
		panic(result.Error)
	}
}

func QueryStockByDate2Str(dateStr string) int {

	return len(stockDao.QueryStockByDateStr(dateStr))
}

func QueryStockByDateStr(dateStr string) (s []Stock) {
	db().Find(&s, "date_str = ?", dateStr)
	return
}

//Stock 櫃買市場 資料庫
//寫入周轉排行 但是沒有收盤價
//URL 日期要記得改
func main() {
	url := "https://www.tpex.org.tw/web/stock/aftertrading/daily_turnover/trn_result.php?l=zh-tw&t=D&d=110/05/04&s=0,asc,1&o=htm"

	urlarray := strings.Split(url, "&")
	dateStr := strings.Split(urlarray[2], "=")[1]

	c := colly.NewCollector()

	count := 0

	if QueryStockByDate2Str(dateStr) > 0 {
		fmt.Print("已經擁有")
	} else {
		var stock Stock
		// 當Visit訪問網頁後，在網頁響應(Response)之後、發現這是HTML格式 執行的事情
		c.OnHTML("td", func(e *colly.HTMLElement) { // 每找到一個符合 goquerySelector字樣的結果，便會進這個OnHTML一次

			switch count % 6 {

			case 1:

				//fmt.Printf("排 行:%s", e.Text)
			case 2:
				stock.StockCode = e.Text
				//fmt.Printf("股票代號:%s", e.Text)

				//stockService.QueryStockCloserPrice(dateStr, stock.StockCode)
			case 3:
				stock.StockName = e.Text
				//fmt.Printf("股票名稱:%s", e.Text)
			case 4:
				stock.StockVolume, _ = strconv.Atoi(strings.Replace(e.Text, ",", "", -1))
				stock.StockVolume = stock.StockVolume / 1000
				//fmt.Printf("總成交股數:%s", e.Text)
			case 5:
				stock.IssuedShares, _ = strconv.Atoi(strings.Replace(e.Text, ",", "", -1))
				stock.IssuedShares = stock.IssuedShares / 1000
				//fmt.Printf("發行股數:%s", e.Text)
			case 0:
				stock.TurnoverRate, _ = strconv.ParseFloat(e.Text, 64)
				//fmt.Printf("週轉率:%s \n", e.Text)
				stock.DateStr = dateStr
				//fmt.Print(stock)
				if count != 0 {
					Insert(stock)
				}
			}

			count++

		})

		defer db().Close()
		c.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
			r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
		})

		c.Visit(url)
	}

}
