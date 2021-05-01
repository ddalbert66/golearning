package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "D:/GitBranch/GOlangLearn/golearning/src/BsiteFileTest/test.txt"

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	bs := []byte{65, 66, 67, 68, 69, 70}
	file.Write(bs)
}
