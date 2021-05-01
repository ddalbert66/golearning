package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileName := "D:/GitBranch/GOlangLearn/golearning/src/BsiteFileTest/testStart2.txt"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	b1 := bufio.NewReader(file)

	//1.高效讀取
	p := make([]byte, 1024)
	n1, err := b1.Read(p)
	fmt.Println(n1)
	fmt.Println(string(p[:n1]))

	// s1, err := b1.ReadString('\n')
	// fmt.Println(err)
	// fmt.Println(s1)

	// s1, err = b1.ReadString('\n')
	// fmt.Println(err)
	// fmt.Println(s1)

	// for {
	// 	s1, err := b1.ReadString('\n')
	// 	if err == io.EOF {
	// 		fmt.Println("讀取完畢")
	// 		break
	// 	}
	// 	fmt.Println(s1)
	// }

	b2 := bufio.NewReader(os.Stdin)
	s2, _ := b2.ReadString('\n')
	fmt.Println(s2)

}
