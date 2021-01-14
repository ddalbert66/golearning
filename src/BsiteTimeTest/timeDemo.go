package main

import (
	"time"
	"fmt"
)

const shortForm = "2006年01月02日"

func main() {
	//1.獲取當前時間
	//模板的日期必需是固定 06-1-2-3-4-5
	t1 := time.Now()
	fmt.Printf("formate %T\n",t1);
	fmt.Println(t1)

	//2.獲取指定時間
	//func Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
	t2 :=  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(t2)

	s1 := t1.Format("2006年1月2日 15:04:05")
	fmt.Println(s1)

	str := "2009年11月02日"
	t3,err := time.Parse(shortForm,str)
	if err != nil {
		fmt.Println(err)
	}
	s3 := t3.Format("2006年1月2日 15:04:05")
	fmt.Println(s3)

	year,month,day := t1.Date()//年 月 日
	fmt.Println(year,month,day)

	hour,min,sec := t1.Clock()//時 分 秒
	fmt.Println(hour,min,sec)

}