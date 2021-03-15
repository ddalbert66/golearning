package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// 一年零一個月一天之後
	// 一段時間之後
	fmt.Println(now.Add(time.Duration(24) * time.Hour))

}
