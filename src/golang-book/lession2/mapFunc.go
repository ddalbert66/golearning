package main

import "fmt"

func main() {
	var mapEx = make(map[string]int)
	mapEx["a"] = 1
	name, ok := mapEx["a"]
	fmt.Println("第一次", name, ok)

	if name, ok := mapEx["a"]; ok {
		fmt.Println("IF 第一次", name, ok)
	}
	delete(mapEx, "a")
	name2, ok2 := mapEx["a"]
	fmt.Println("第二次", name2, ok2)

	if name2, ok2 := mapEx["a"]; ok2 {
		fmt.Println("IF 第二次", name2, ok2)
	}

}
