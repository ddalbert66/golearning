package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	country := "Albania"
	c := colly.NewCollector()
	// 當Visit訪問網頁後，在網頁響應(Response)之後、發現這是HTML格式 執行的事情
	c.OnHTML("a[href]", func(e *colly.HTMLElement) { // 每找到一個符合 goquerySelector字樣的結果，便會進這個OnHTML一次

		if strings.Contains(e.Attr("href"), "countries") {
			if strings.Contains(e.Text, country) {
				for k, v := range getTopNum(20, c, e.Attr("href")) {
					fmt.Printf("index:%d value:%s \n", k+1, v)
				}

			}
		}

	})

	c.Visit("https://www.alexa.com/topsites/countries")
}

func getTopNum(num int, c *colly.Collector, url string) []string {
	var count = 0
	// 當Visit訪問網頁後，在網頁響應(Response)之後、發現這是HTML格式 執行的事情
	resultTop := []string{}

	c.OnHTML(".DescriptionCell", func(e *colly.HTMLElement) { // 每找到一個符合 goquerySelector字樣的結果，便會進這個OnHTML一次
		if count < num {
			count++
			str := strings.Replace(e.Text, " ", "", -1)
			str = strings.Replace(str, "\n", "", -1)
			resultTop = append(resultTop, str)

		}
	})

	c.Visit("https://www.alexa.com/topsites/" + url)
	return resultTop
}
