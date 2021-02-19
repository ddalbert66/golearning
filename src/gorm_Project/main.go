package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//func main() {

// 	db, err := gorm.Open("mysql", "root:abcd1234@tcp(127.0.0.1:3306)/bento?charset=utf8&loc=Local&parseTime=true")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	//db.CreateTable(&User_info{})//後面加了S
// 	db.Table("User_info").CreateTable(&models.User_info{})
// }
