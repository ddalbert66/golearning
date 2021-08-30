package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("D:/tests.txt")
	catchErr(err)

	defer func() {
		fmt.Println("file close")
		if f != nil {
			f.Close()
		}
	}()

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	catchErr(err)

	fmt.Printf("%b bytes: %s\n", n1, string(b1))
}

func catchErr(err error) {
	if err != nil {
		panic(err)
	}
}
