package main

import (
	"fmt"
	"os"
)

func main() {
	aa := os.Args

	for k, v := range aa {
		switch v {

		case "jjj":
			fmt.Println("Key:", k, "jjjValue", v)
		case "jjjj":
			fmt.Println("Key:", k, "jjjjValue", v)
		}

	}
}
