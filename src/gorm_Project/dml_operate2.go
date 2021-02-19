package main

import (
	"goLearning20200930/src/gorm_Project/projectModels"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//func main() {

	db, err := gorm.Open("mysql", "root:abcd1234@tcp(127.0.0.1:3306)/bento?charset=utf8&loc=Local&parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Table("enumeration").CreateTable(&projectModels.Enumeration{})
	db.Table("enum_type").CreateTable(&projectModels.EnumType{})
}
