package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	srcFile := "D:/GitBranch/GOlangLearn/golearning/src/BsiteFileTest/test.txt"
	destFile := "destFile.txt"
	total, err := CopyFile(srcFile, destFile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(total)

}

func CopyFile(srcFile, destFile string) (int, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}

	file2, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return 0, err
	}

	defer file1.Close()
	defer file2.Close()
	bs := make([]byte, 128, 128)
	total := 0
	for {
		n, error := file1.Read(bs)

		if n == 0 || error == io.EOF {
			fmt.Println("到了文件的末尾")
			break
		}
		total = +n
		file2.Write(bs[:n])
	}

	return total, err

}
