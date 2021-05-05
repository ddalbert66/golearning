package main

import "fmt"

func main() {
	stringcase := "t"

	switch stringcase {

	case "t":
		fmt.Println("value is t")
		fallthrough
	case " ":
		fmt.Println("value is b")

	case "a":
		fmt.Println("value is a")

	}
}
