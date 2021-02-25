package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/urfave/cli"
)

type orderValue struct {
	top     string
	country string
}

type resultVal struct {
	topNumResult []string
	country      string
}

var ordervalue orderValue

//func main() {

	app := cli.NewApp()
	app.Name = "clawer"
	app.Usage = "clawer"
	app.Action = order
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "top",
			Usage: "get top number",
		},
		cli.StringFlag{
			Name:  "country",
			Usage: "getList by country",
		},
	}

	app.Run(os.Args)
}

func order(c *cli.Context) error {
	ordervalue = orderValue{
		top:     c.String("top"),
		country: c.String("country"),
	}

	return exec()
}

func exec() error {
	fmt.Println("top:", ordervalue.top)
	fmt.Println("country:", ordervalue.country)

	//判斷top是否有輸入值
	if len(ordervalue.top) > 0 {
		num, _ := strconv.Atoi(ordervalue.top)
		fmt.Println(getTopNum(num))
	}

	//判斷country是否有輸入值
	if len(ordervalue.country) > 0 {
		getByCountry(ordervalue.country)
	}
	return nil
}

func getByCountry(country string) {
	c := colly.NewCollector()
	// 當Visit訪問網頁後，在網頁響應(Response)之後、發現這是HTML格式 執行的事情
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		//取得attr(href)變數
		if strings.Contains(e.Attr("href"), "countries") {
			if strings.Contains(e.Text, country) {
				//印出前20 Country
				for k, v := range getTopNumForCountry(20, c, e.Attr("href")) {
					fmt.Printf("index:%d value:%s \n", k+1, v)
				}

			}
		}

	})

	c.Visit("https://www.alexa.com/topsites/countries")
}

func getTopNum(num int) []string {

	var count = 0
	c := colly.NewCollector()
	resultTop := []string{}

	// 當Visit訪問網頁後，在網頁響應(Response)之後、發現這是HTML格式 執行的事情

	c.OnHTML(".DescriptionCell", func(e *colly.HTMLElement) { // 每找到一個符合 goquerySelector字樣的結果，便會進這個OnHTML一次
		if count < num {
			count++
			str := strings.Replace(e.Text, " ", "", -1)
			str = strings.Replace(str, "\n", "", -1)
			resultTop = append(resultTop, str)

		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.alexa.com/topsites/")

	return resultTop
}

func getTopNumForCountry(num int, c *colly.Collector, url string) []string {
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
