package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	/*
		Seek(offset int64,whence int) (int64,error),設置指針光標的位置
			第一個參數:偏移量
			第二個參數:如何設置
			 0:seekstart,表示相對於文件開始
			 1:seekcurrent,表示相對於當前位置的偏移量
			 2:seekend,表示相對於末尾

			 //Seek whence values
			 const {
				 SeekStart = 0
				 SeekCurrent = 1
				 SeeKend = 2
			 }
	*/
	fileName := "D:/GitBranch/GOlangLearn/golearning/src/BsiteFileTest/test.txt"
	file, error := os.OpenFile(fileName, os.O_RDONLY, os.ModePerm)

	if error != nil {
		fmt.Println(error)
	}

	defer file.Close()
	bs := make([]byte, 10, 10)
	file.Read(bs)
	fmt.Println(string(bs[:]))

	file.Seek(10, io.SeekStart)
	file.Read(bs)
	fmt.Println(string(bs[:]))

}
