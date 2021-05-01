package main

import (
	"bytes"
	"fmt"
	"goLearning20200930/src/gorm_Project/dao/stockDao"
	"io/ioutil"
	"strings"

	"github.com/gocolly/colly"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

// type StockV2StockVolume stockModels.StockV2StockVolume

// func db4() *gorm.DB {
// 	return dao.GetDB()
// }

// func InsertStockV2SV(s StockV2StockVolume) uint {
// 	result := db4().Create(&s)
// 	if result.Error != nil {
// 		panic(result.Error)
// 	}
// 	return s.ID
// }

//名次	股票名稱	收盤價	漲跌	漲跌幅	成交量	週轉率
//1	   8021  尖點	 32.20	+ 2.60	 +8.78%	 38,222	 26.88%
func main() {
	StockV2arr := stockDao.QueryStockByDate("110/03/17")

	//var stock StockV2StockVolume

	for _, st := range StockV2arr {
		if len(st.StockCode) == 4 && !strings.HasPrefix(st.StockCode, "00") {
			getStockType()
			//InsertStockV2SV(stock)
		}

	}

}

func getStockType() (int64, string) {
	url := "https://www.tej.com.tw/webtej/doc/uid.htm"
	c := colly.NewCollector()
	var result int64
	var stockRemark string
	c.OnHTML("td font", func(e *colly.HTMLElement) {
		result, _ := DecodeBig5([]byte(e.Text))
		fmt.Println(string(result))

		fmt.Println()
	})

	c.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	c.Visit(url)
	return result, stockRemark
}

func DecodeBig5(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, traditionalchinese.Big5.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}
