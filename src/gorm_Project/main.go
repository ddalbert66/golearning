package main

import (
	"goLearning20200930/src/gorm_Project/dao"
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

func Insert(s Stock) uint {
	result := db().Create(&s)
	if result.Error != nil {
		panic(result.Error)
	}
	return s.ID
}

//排 行	股票代號	股票名稱	總成交股數	發行股數	週轉率(%)
//1	   3540	        曜越        16,904,464	66,613,044	25.37
func main() {
	url := "https://www.tpex.org.tw/web/stock/aftertrading/daily_turnover/trn_result.php?l=zh-tw&t=D&d=110/02/26&s=0,asc,1&o=htm"
	urlarray := strings.Split(url, "&")
	dateStr := strings.Split(urlarray[2], "=")[1]

	c := colly.NewCollector()

	count := 0

	var stock Stock
	// 當Visit訪問網頁後，在網頁響應(Response)之後、發現這是HTML格式 執行的事情
	c.OnHTML("td", func(e *colly.HTMLElement) { // 每找到一個符合 goquerySelector字樣的結果，便會進這個OnHTML一次

		if count < 600 {
			switch count % 6 {

			case 1:

				//fmt.Printf("排 行:%s", e.Text)
			case 2:
				stock.StockCode = e.Text
				//fmt.Printf("股票代號:%s", e.Text)
			case 3:
				stock.StockName = e.Text
				//fmt.Printf("股票名稱:%s", e.Text)
			case 4:
				stock.StockVolume, _ = strconv.Atoi(strings.Replace(e.Text, ",", "", -1))
				//fmt.Printf("總成交股數:%s", e.Text)
			case 5:
				stock.IssuedShares, _ = strconv.Atoi(strings.Replace(e.Text, ",", "", -1))
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
		}

	})

	c.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	c.Visit(url)

}
