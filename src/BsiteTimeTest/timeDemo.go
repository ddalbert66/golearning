package main

import (
	"fmt"
	"math/rand"
	"time"
)

const shortForm = "2006年01月02日"

func main() {

	t1 := time.Now()
	fmt.Printf("%T\n", t1)
	fmt.Println(t1)

	//2006年1月2日下午3點4分5秒
	s1 := time.Now().Format("2006年01月02日 15:04:05")
	fmt.Println(s1)

	/*
	 * time.Prase("模板",str )
	 */
	s2 := "2010年3月4日"
	t2, err := time.Parse("2006年1月2日", s2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t2)

	//當前時間取出指定內容
	year, month, day := t1.Date() //年 月 日
	fmt.Println(year, month, day)

	hour, min, sec := t1.Clock()
	fmt.Println(hour, min, sec)

	//時間戳 指定日期 距離1970年1月1日0點0分0秒 的時間差距: 秒 那秒
	t4 := time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC)
	timeStamp1 := t4.Unix()
	fmt.Println(timeStamp1)
	timeStamp2 := t1.Unix()
	fmt.Println(timeStamp2)
	timeStamp3 := t4.Unix()
	fmt.Println(timeStamp3)
	timeStamp4 := t1.Unix()
	fmt.Println(timeStamp4)

	//睡眠
	//time.Sleep(3 * time.Second) //讓當前的程序進入睡眠狀態
	fmt.Println("main.... over....")

	//睡眠[1-10]的隨機秒數
	//rand.Seed 為了放入不同的參數讓他去產生隨機變數
	rand.Seed(time.Now().UnixNano())
	//取 0~9 +1
	randNum := rand.Intn(10) + 1
	fmt.Println(randNum)
	randNum = rand.Intn(10) + 1
	fmt.Println(randNum)
	//time.Sleep(time.Duration(randNum) * time.Second)
	fmt.Println("睡醒了")

}
