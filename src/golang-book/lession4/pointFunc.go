package main

import "fmt"

func main() {
	val := 2
	changeValue(val)
	fmt.Println("withOut point:", val)
	changePointValue(&val)
	fmt.Println("have point:", val)
}

func changeValue(val int) {
	val = 5
	fmt.Println("print in func:", val)
}

func changePointValue(val *int) {
	*val = 5
}
