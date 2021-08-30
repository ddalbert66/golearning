package main

import "fmt"

func main() {
	valStar := new(int)
	changePointValue(valStar)
	fmt.Println(valStar)
	fmt.Println(*valStar)
	val := 1
	fmt.Println(val)
}

func changeValue(val int) {
	val = 5
	fmt.Println("print in func:", val)
}

func changePointValue(val *int) {
	*val = 5
}
