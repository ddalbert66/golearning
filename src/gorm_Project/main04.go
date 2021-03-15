package main

import (
	"goLearning20200930/src/gorm_Project/stockService"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//名次	股票名稱	收盤價	漲跌	漲跌幅	成交量	週轉率
//1	   8021  尖點	 32.20	+ 2.60	 +8.78%	 38,222	 26.88%
func main() {
	stockService.FuncgetGoodStock("4529")

}
