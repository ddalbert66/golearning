package main

import (
	"fmt"
	"goLearning20200930/src/gorm_Project/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	db, err := gorm.Open("mysql", "root:abcd1234@tcp(127.0.0.1:3306)/bento?charset=utf8&loc=Local&parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Table("store").Create(&models.Store{
		Name:        "好吃餐廳",
		Phone_no:    123456789,
		Region:      "1",
		Create_user: "sheldon",
		Update_user: "",
	})

	var store models.Store
	db.Table("store").First(&store, "name=?", "好吃餐廳")

	fmt.Print(store)
}
