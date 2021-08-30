package main

import (
	"encoding/json"
	"fmt"
)

type Teacher struct {
	Name    string `json:"-"`          //==>在使用json編碼時,這個編碼不參與
	Age     int    `json:"age,string"` //==>在json編碼時,將age轉乘string類型
	Address string
}

func main() {
	teacher1 := Teacher{
		Name:    "teacher1",
		Age:     19,
		Address: "彰化",
	}

	jsonString, _ := json.Marshal(&teacher1)

	fmt.Println(string(jsonString))
}
