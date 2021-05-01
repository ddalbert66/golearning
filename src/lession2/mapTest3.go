package main

import "fmt"

func main() {

	map1 := make(map[int]string, 3)

	map1[0] = "01"
	map1[1] = "02"
	_, ok := map1[2]
	fmt.Println(ok)
	for k, v := range map1 {

		fmt.Println("k :", k, "v :", v)
	}
}
