package main

import (
	"fmt"
	"goLearning20200930/src/gorm_Project/dao"
	"goLearning20200930/src/gorm_Project/dao/stockDao"

	"strings"

	"github.com/jinzhu/gorm"
)

//type Stock stockModels.Stock

func db() *gorm.DB {
	return dao.GetDB()
}

/*
*	周轉率是否有持續增加
*
 */
func main() {

	//stockArray := stockDao.QueryStockByName("勤凱")
	stockArray := stockDao.QueryStock()
	count := 0
	//計算周轉這幾天排行 有沒有越來越高
	countStock := 5

	var needStockCode string
	var stockName string
	//var lastRate float64
	m := make(map[string][]float64)

	for _, st := range stockArray {

		if len(stockName) == 0 {
			stockName = st.StockCode
		} else if strings.Compare(stockName, st.StockCode) == 0 && len(stockName) == 4 {
			m[stockName] = append(m[stockName], st.TurnoverRate)
			//fmt.Printf("stockName : %s ,TurnoverRate :% f \n", stockName, st.TurnoverRate)
			if len(m[stockName]) > 24 {
				fmt.Printf("count :%d, x[24]:%f,x[23]:%f ,StockCode :% s \n", count, m[stockName][24], m[stockName][23], st.StockCode)
				if count >= 20 && checkStockRule(m, stockName, countStock) {
					fmt.Printf("needStockCode : %s ,StockCode :% s \n", needStockCode, st.StockCode)
					if !strings.Contains(needStockCode, st.StockCode) {

						count = 0
						//stockService.FuncgetGoodStock(st.StockCode)
						needStockCode = needStockCode + ":" + st.StockCode
					}
				}
			}

			count++
		} else if stockName != st.StockCode {
			stockName = st.StockCode
			count = 0
		}
	}

	defer db().Close()

}

//確認幾天內 周轉排行都是越來越大
func checkStockRule(m map[string][]float64, stockName string, num int) bool {
	rule := true
	count := 0
	for i := 1; i < len(m[stockName]); i++ {

		oldCount := len(m[stockName]) - (i + 1)
		newCount := len(m[stockName]) - i
		fmt.Println("新的:", stockName, ":", m[stockName][newCount], ",舊的:", stockName, ":", m[stockName][oldCount])
		if m[stockName][newCount] > m[stockName][oldCount] {
			fmt.Println(rule, count)
			count++
			if rule == false {
				return rule
			} else if num == count && rule == true {
				return rule
			}
		} else {
			rule = false
		}

	}
	return rule
}
