package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(FormateMoney(1234.45))
}

//將錢以千位數分隔
func FormateMoney(money float64) string {
	var result string
	moneyStr := strconv.FormatFloat(money, 'f', 2, 64)
	fmt.Println(moneyStr)
	if strings.Contains(moneyStr, ".") {
		split := strings.Split(moneyStr, ".")
		fmt.Println(split)

		if len(split) == 0 {
			fmt.Println(0)
		} else if len(split) == 2 {
			var count int
			//每隔三位數 組成字串
			count = 0
			pNum := strings.Split(split[0], "")
			fmt.Println(split[0])
			fmt.Println(pNum)
			for i := 0; i < len(pNum); i++ {
				count++
				fmt.Println(pNum[len(pNum)-1-i])
				if count == 3 {
					result = result + "," + pNum[len(pNum)-1-i]
					count = 0
				} else {
					result = result + pNum[len(pNum)-1-i]
				}
			}

			if len(split[1]) != 0 {
				if split[1] == "00" {
					//donothing
				} else {
					result = result + "." + split[1]
				}
			}

			return result
		}

	}
	return result
}
