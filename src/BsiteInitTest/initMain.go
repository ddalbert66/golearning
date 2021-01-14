package main

import (
	"goLearning20200930/src/BsiteInitTest/initTest"
	"fmt"
)

func main(){
	initTest.Run()
	initTest.RunB()
}

func init() {
	fmt.Println("mainInit")
}