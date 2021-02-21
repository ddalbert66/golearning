package homework4

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

var ordervalue orderValue

func main() {

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

	num, _ := strconv.Atoi(ordervalue.top)
	fmt.Println(getTopCountry(num))
	return exec()
}

func exec() error {
	fmt.Println("top:", ordervalue.top)
	fmt.Println("country:", ordervalue.country)

	return nil
}

func getTopCountry(num int) []string {

	var count = 0
	c := colly.NewCollector()
	resultTop := []string{}

	// 當Visit訪問網頁後，在網頁響應(Response)之後、發現這是HTML格式 執行的事情

	c.OnHTML(".DescriptionCell", func(e *colly.HTMLElement) { // 每找到一個符合 goquerySelector字樣的結果，便會進這個OnHTML一次
		if count > num {
			str := strings.Replace(e.Text, " ", "", -1)
			str = strings.Replace(str, "\n", "", -1)
			resultTop = append(resultTop, str)

		} else {
			count++
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.alexa.com/topsites/")

	return resultTop
}
