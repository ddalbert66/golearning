package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(FormateMoney(-1234.0))
}

//將錢以千位數分隔
func FormateMoney(money float64) string {
	moneyStr := strconv.FormatFloat(money, 'f', -1, 64)
	fmt.Println(moneyStr)

	//全部的數字 包含小數點前後
	var result string
	//小數點前的數字
	var pMoneyStr string
	////小數點後的數字
	var moneyDotStr string
	if !strings.Contains(moneyStr, ".") {
		pMoneyStr = moneyStr
		moneyDotStr = ""
	} else {
		split := strings.Split(moneyStr, ".")
		pMoneyStr = split[0]
		moneyDotStr = "." + split[1]
	}
	var count int
	//每隔三位數 組成字串
	count = 0
	pNum := strings.Split(pMoneyStr, "")
	pMoneyStr = ""
	for i := len(pNum) - 1; i >= 0; i-- {

		if len(pNum) > i+1 && count%3 == 0 {
			pMoneyStr = pNum[i] + "," + pMoneyStr
			count = 0
		} else {
			pMoneyStr = pNum[i] + pMoneyStr
		}
		if strings.Compare(pNum[i], "-") == 1 {
			count++
		}
	}
	result = pMoneyStr + moneyDotStr
	return result
}
