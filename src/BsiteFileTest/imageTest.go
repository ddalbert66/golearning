package main

import (
	"os"
	"fmt"
	"image"
	"github.com/axgle/mahonia"
)
func main() {

	var enc mahonia.Decoder
	enc = mahonia.NewDecoder("UTF-8")
	
	filePath:= "C:/Users/USER/OneDrive/圖片/aa.png"
	fileInfo , error := os.Stat(filePath)

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
	file,err :=os.Open(filePath)
	if err !=nil {
		fmt.Println("err")
		return
	}
	//3.關閉文件
	defer file.Close()

	f, err := os.Open(filePath)
    if err != nil {
		fmt.Println(err)
        return 
    }
	image, _, err := image.Decode(f)
	
	fmt.Println(image)
  
}