package main

import "fmt"

func main() {
	var (
		maptest   map[string]int
		sliceTest []int
	)

	maptest = make(map[string]int)
	fmt.Println("mapTest is nil ?", maptest == nil)
	fmt.Println("mapTest is nil ?", maptest == nil)
	fmt.Println("mapTest:", maptest)
	maptest["test1"] = 12
	fmt.Println(maptest["test1"])
	fmt.Println(maptest["test2"])
	v, ok := maptest["test2"]
	fmt.Println(ok)
	fmt.Println(v)
	fmt.Println("sliceTest is nil ?", sliceTest == nil)

	teststring := "test"
	fmt.Println(teststring)
	teststring = "test2"
	fmt.Println(teststring)
}
