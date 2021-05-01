package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fileName := "D:/GitBranch/GOlangLearn/golearning/src/BsiteFileTest/testStart2.txt"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	w1 := bufio.NewWriter(file)
	n, err := w1.WriteString("hello world")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	w1.Flush()
	file.Close()
	time.Sleep(time.Second)

	file2, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, os.ModePerm)
	bytes := make([]byte, 100)
	s1 := bufio.NewReader(file2)
	s2, err := s1.Read(bytes)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(s2) + "--")
	fmt.Println(string(bytes) + "-++++")
	defer file2.Close()

}
