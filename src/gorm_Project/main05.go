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

type StockV2StockVolume stockModels.StockV2StockVolume

func db4() *gorm.DB {
	return dao.GetDB()
}

func InsertStockV2SV(s StockV2StockVolume) uint {
	result := db4().Create(&s)
	if result.Error != nil {
		panic(result.Error)
	}
	return s.ID
}

//名次	股票名稱	收盤價	漲跌	漲跌幅	成交量	週轉率
//1	   8021  尖點	 32.20	+ 2.60	 +8.78%	 38,222	 26.88%
func main() {
	StockV2arr := stockDao.QueryStockByDate("110/04/29")

	var stock StockV2StockVolume

	for _, st := range StockV2arr {
		if len(st.StockCode) == 4 && !strings.HasPrefix(st.StockCode, "00") {
			stock.StockCode = st.StockCode
			stock.StockVolume, stock.StockRemark = getNumberOfIssuedShares(st.StockCode)
			InsertStockV2SV(stock)
		}

	}

}

func getNumberOfIssuedShares(stockCode string) (int64, string) {
	url := "https://www.cnyes.com/twstock/intro.aspx?code=" + stockCode
	c := colly.NewCollector()
	var result int64
	var stockRemark string
	c.OnHTML("td span[id]", func(e *colly.HTMLElement) {
		if strings.Compare(e.Attr("id"), "ctl00_ContentPlaceHolder1_Label017") == 0 {
			//strconv.Atoi(strings.Replace(e.Text, ",", "", -1))
			result, _ = strconv.ParseInt(strings.Replace(e.Text, ",", "", -1), 10, 64)
			fmt.Println(stockCode + ":" + strings.Replace(e.Text, ",", "", -1))
		}
		if strings.Compare(e.Attr("id"), "ctl00_ContentPlaceHolder1_Label010_1") == 0 {
			//strconv.Atoi(strings.Replace(e.Text, ",", "", -1))
			stockRemark = e.Text
			fmt.Println(stockCode + ":" + stockRemark)
		}

	})

	c.OnRequest(func(r *colly.Request) { // iT邦幫忙需要寫這一段 User-Agent才給爬
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	c.Visit(url)
	return result, stockRemark
}
