package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/axgle/mahonia"
)

func main() {

	var enc mahonia.Decoder
	enc = mahonia.NewDecoder("UTF-8")

	filePath := "D:/GitBranch/GOlangLearn/golearning/src/BsiteFileTest/test.txt"
	fileInfo, error := os.Stat(filePath)

	if error != nil {
		fmt.Println(error)
		return
	} else {
		fmt.Println(fileInfo)
	}

	//文件名
	fmt.Println(fileInfo.Name())
	//文件大小
	fmt.Println(fileInfo.Size())
	//是否為目錄
	fmt.Println(fileInfo.IsDir())
	//修改時間
	fmt.Println(fileInfo.ModTime())
	//權限
	fmt.Println(fileInfo.Mode())

	//讀取文件
	//1.打開文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("err")
		return
	}
	//3.關閉文件
	defer file.Close()

	//2.讀取
	//bs := make([]byte,4,4)

	//第一次讀取
	// n := -1
	// for {
	// 	n,error = file.Read(bs)
	// 	if n==0|| error == io.EOF {
	// 		fmt.Println("到了文件的末尾")
	// 		break
	// 	}
	// 	fmt.Println(string(bs[:n]))
	// }

	//创建一个 *Reader , 是带缓冲的, 默认缓冲区为4096个字节
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF {                  //io.EOF表示文件的末尾
			break
		}
		//输出内容
		fmt.Println("UTF-8 to GBK:", enc.ConvertString(str))
	}

}
