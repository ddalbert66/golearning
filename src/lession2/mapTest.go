package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "sheldon",
		"years":   "18",
		"sex":     "male",
		"country": "taiwan",
	}

	m2 := make(map[string]int) //m2 == empty

	var m3 map[string]int //m3 == nil

	fmt.Println(m, m2, m3)

	for k, v := range m {
		fmt.Println(k, v)
	}

	if result, ok := m["name"]; ok {
		fmt.Println(result, ok)
	} else {
		fmt.Println("without key")
	}

	fmt.Println(len(m))
}
