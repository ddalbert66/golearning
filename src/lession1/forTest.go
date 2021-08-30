package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		if len(arg) > 0 {
			s += sep + arg + "\r\n"
			sep = ""
		}

	}
	fmt.Println(s)

	fmt.Println(strings.Join(os.Args[1:], "\r\n"))
}
