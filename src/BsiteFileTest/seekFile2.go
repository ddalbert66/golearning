package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	/*
		斷點續傳:
			文件傳遞 :文件複製

			複製流程如下
			1.一邊複製一邊紀錄複製的總量

	*/
	//1. 開啟三個file 一個暫存 一個接收端 一個發送端
	//2. 讀取暫存檔 確認資料讀取到哪裡 如果為空則從0開始
	fileNameStart := "D:/GitBranch/GOlangLearn/golearning/src/BsiteFileTest/testStart.txt"
	fileNameEnd := "D:/GitBranch/GOlangLearn/golearning/src/BsiteFileTest/testEnd.txt"
	fileNameTemp := "D:/GitBranch/GOlangLearn/golearning/src/BsiteFileTest/tempFile.txt"

	fileStart, err := os.OpenFile(fileNameStart, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	errorHandle(err)
	fileEnd, err := os.OpenFile(fileNameEnd, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	errorHandle(err)
	fileTemp, err := os.OpenFile(fileNameTemp, os.O_CREATE|os.O_RDWR, os.ModePerm)
	errorHandle(err)

	fileBt := make([]byte, 100, 100)
	nint, err := fileTemp.Read(fileBt)
	errorHandle(err)
	//fileTemp.Seek(offset int64, whence int)
	fmt.Println(strconv.ParseInt(string(fileBt[:nint]), 10, 64))

	count, err := strconv.ParseInt(string(fileBt[:nint]), 10, 64)
	errorHandle(err)
	fileStart.Seek(count, io.SeekStart)
	total := count
	fileStart.Read(fileBt[:count])

	fmt.Println(string(fileBt[:count]))

	fileTemp.WriteString(strconv.Itoa(int(count)))

	defer fileStart.Close()
	defer fileEnd.Close()
	defer fileTemp.Close()
}

func errorHandle(err error) {
	if err != nil {
		fmt.Println(err)
	}

}
