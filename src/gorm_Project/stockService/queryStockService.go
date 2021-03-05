package stockService

import (
	"fmt"
	"goLearning20200930/src/gorm_Project/stockModels"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type StockV2 stockModels.StockV2

// func db() *gorm.DB {
// 	return dao.GetDB()
// }

func Insert(s StockV2) uint {
	result := db().Create(&s)
	if result.Error != nil {
		panic(result.Error)
	}
	return s.ID
}

//名次	股票名稱	收盤價	漲跌	漲跌幅	成交量	週轉率
//1	   8021  尖點	 32.20	+ 2.60	 +8.78%	 38,222	 26.88%
func QueryStockCloserPrice(dateStr string, stockNo string) {
	year, _ := strconv.Atoi(strings.Split(dateStr, "/")[0])
	fmt.Print(year)
	year = year + 1911
	mon := strings.Split(dateStr, "/")[1]
	date := strings.Split(dateStr, "/")[2]

	newDateStr := strconv.Itoa(year) + mon + date

	url := "https://www.twse.com.tw/exchangeReport/STOCK_DAY?response=json&date=" + newDateStr + "&stockNo=" + stockNo
	fmt.Println(url)
	c := colly.NewCollector()

	// 當Visit訪問網頁後，在網頁響應(Response)之後、發現這是HTML格式 執行的事情
	c.OnHTML("pre", func(e *colly.HTMLElement) { // 每找到一個符合 goquerySelector字樣的結果，便會進這個OnHTML一次
		fmt.Print(e.Text)
	})

	c.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	c.Visit(url)

}
