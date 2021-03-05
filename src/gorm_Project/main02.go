package main

import (
	"fmt"
	"goLearning20200930/src/gorm_Project/dao"
	"goLearning20200930/src/gorm_Project/stockModels"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type stockV2 stockModels.StockV2

func db() *gorm.DB {
	return dao.GetDB()
}

func Insert(s stockV2) uint {
	result := db().Create(&s)
	if result.Error != nil {
		panic(result.Error)
	}
	return s.ID
}

//名次	股票名稱	收盤價	漲跌	漲跌幅	成交量	週轉率
//1	   8021  尖點	 32.20	+ 2.60	 +8.78%	 38,222	 26.88%
func main() {
	url := "https://www.moneydj.com/Z/ZG/ZG.djhtm?a=BD"

	c := colly.NewCollector()

	count := 0

	var stock stockV2
	// 當Visit訪問網頁後，在網頁響應(Response)之後、發現這是HTML格式 執行的事情
	c.OnHTML("td", func(e *colly.HTMLElement) { // 每找到一個符合 goquerySelector字樣的結果，便會進這個OnHTML一次

		if strings.Contains(e.Attr("class"), "t3") {

			switch count % 7 {
			case 0:

				fmt.Printf("排 行:%s", e.Text)
			case 1:
				stock.StockCode = strings.Split(e.Text, "'")[1]
				fmt.Printf("股票代號:%s", stock.StockCode)

				stock.StockName = strings.Split(e.Text, "'")[3]
				fmt.Printf("股票名稱:%s", stock.StockName)
			case 2:

				stock.ClosingPrice, _ = strconv.ParseFloat(strings.TrimSpace(e.Text), 64)
				fmt.Printf("收盤價:%f", stock.ClosingPrice)
			case 3:
				stock.UpAndDown = e.Text
				fmt.Printf("漲跌:%s", stock.UpAndDown)
			case 4:
				stock.UpAndDownRate = e.Text
				fmt.Printf("漲跌幅:%s", stock.UpAndDownRate)
			case 5:
				stock.StockVolume, _ = strconv.Atoi(strings.Trim(strings.TrimSpace(e.Text), ","))
				fmt.Printf("成交量:%d", stock.StockVolume)
			case 6:
				stock.TurnoverRate, _ = strconv.ParseFloat(strings.TrimSpace(strings.Trim(e.Text, "%")), 64)
				fmt.Printf("週轉率:%f \n", stock.TurnoverRate)
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
