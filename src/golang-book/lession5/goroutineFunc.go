package main

import (
	"fmt"
)

func forRun() {
	for i := 0; i <= 10; i++ {
		fmt.Print(i, ",")
	}
	fmt.Println()
}

func main() {
	go forRun()

	input := ""
	fmt.Scanln(&input)
}
