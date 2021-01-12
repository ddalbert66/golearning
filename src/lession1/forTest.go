package main

import (
	"fmt"
	"strconv"
)

//sheldon test covertToBin go doc
func convertToBin(num int) string {
	result := ""
	for ; num > 0; num /= 2 {
		fmt.Println("num/2:" + strconv.Itoa(num/2))
		isb := num % 2
		fmt.Println(isb)
		result = strconv.Itoa(isb) + result
	}
	return result
}

func main() {
	fmt.Println(convertToBin(5))
	fmt.Println(convertToBin(101))
}
