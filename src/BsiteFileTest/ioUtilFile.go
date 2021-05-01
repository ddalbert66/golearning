package main

import (
	"io/ioutil"
	"os"
)

func main() {
	fileName := "D:/Test/test.txt"

	fileValue := "安安你好"

	ioutil.WriteFile(fileName, []byte(fileValue), os.ModePerm)
}
