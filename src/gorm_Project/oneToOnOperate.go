package main

import (
	"fmt"
	"goLearning20200930/src/gorm_Project/oneToOneModels"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	items = []oneToOneModels.Item{
		{ItemName: "chicken", Amount: 119.8},
		{ItemName: "vegetable", Amount: 30.2},
		{ItemName: "beef", Amount: 17.2},
	}
)

//func main() {

	db, err := gorm.Open("mysql", "root:abcd1234@tcp(127.0.0.1:3306)/bento?charset=utf8&loc=Local&parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for index := range items {
		db.Table("item").Create(&items[index])
	}
	order := oneToOneModels.Order{Status: "pending"}
	db.Table("order").Create(&order)
	item1 := oneToOneModels.OrderItem{OrderID: order.ID, ItemID: items[0].ID, Quantity: 1}
	item2 := oneToOneModels.OrderItem{OrderID: order.ID, ItemID: items[1].ID, Quantity: 4}
	db.Table("order_item").Create(&item1)
	db.Table("order_item").Create(&item2)

	rows, err := db.Table("order").Where("order.id=? and status = ?", order.ID, "pending").
		Joins("join order_item on order_item.order_id=order.id").
		Joins("join item on item.id=order_item.id").
		//select 代表要查詢的欄位
		Select("order.id,order.status,order_item.order_id,order_item.item_id,order_item.quantity" +
			",item.item_name,item.amount").Rows()

	if err != nil {
		log.Panic(err)
	}

	defer rows.Close()
	newOrder := &oneToOneModels.Order{}
	newOrder.OrderItems = make([]oneToOneModels.OrderItem, 0)

	for rows.Next() {
		orderItem := oneToOneModels.OrderItem{}
		item := oneToOneModels.Item{}
		//scan 代表select 參數對應
		err = rows.Scan(&newOrder.ID, &newOrder.Status, &orderItem.OrderID,
			&orderItem.ItemID, &orderItem.Quantity, &item.ItemName, &item.Amount)

		if err != nil {
			fmt.Print(err)
		}
		orderItem.Item = item
		newOrder.OrderItems = append(newOrder.OrderItems, orderItem)
	}
	fmt.Print(newOrder)
}
