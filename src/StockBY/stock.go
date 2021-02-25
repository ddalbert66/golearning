package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	stocks := []struct {
		name string
	}{
		{"尼克森"},
		{"點序"},
		{"勤凱"},
		{"富鼎"},
		{"宏致"},
		{"萬潤"},
		{"愛普"},
		{"金麗科"},
		{"晶豪科"},
		{"創意"},
		{"華邦電"},
		{"晶技"},
		{"台半"},
		{"凌陽"},
		{"九齊"},
		{"杰力"},
		{"家登"},
		{"精材"},
		{"雍智"},
		{"精星"},
		{"晶芯科"},
		{"安普新"},
		{"立積"},
	}
	c := colly.NewCollector()

	// 當Visit訪問網頁後，在網頁響應(Response)之後、發現這是HTML格式 執行的事情
	c.OnHTML("td", func(e *colly.HTMLElement) { // 每找到一個符合 goquerySelector字樣的結果，便會進這個OnHTML一次
		for _, stock := range stocks {
			if strings.Contains(e.Text, stock.name) {
				fmt.Println(e.Text)

			}
		}

	})

	c.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	c.Visit("https://www.tpex.org.tw/web/stock/aftertrading/daily_turnover/trn_result.php?l=zh-tw&t=D&d=110/02/22&s=0,asc,1&o=htm")
}
