package main

import (
	"fmt"
	"goLearning20200930/src/gorm_Project/projectModels"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	enums = []projectModels.Enumeration{
		{Code: "1", Name: "北區", CreateUser: "sheldon", Version: 1},
		{Code: "2", Name: "中區", CreateUser: "sheldon", Version: 1},
		{Code: "3", Name: "南區", CreateUser: "sheldon", Version: 1},
		{Code: "4", Name: "東區", CreateUser: "sheldon", Version: 1},
	}
)

func main() {

	db, err := gorm.Open("mysql", "root:abcd1234@tcp(127.0.0.1:3306)/bento?charset=utf8&loc=Local&parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	enumType := projectModels.EnumType{Code: "regionEnumType", Name: "餐廳地區", CreateUser: "sheldon", Version: 1}
	db.Table("enum_type").Create(&enumType)

	fmt.Print(enumType.ID)

	for index := range enums {
		enums[index].EnumTypeID = enumType.ID
		db.Table("enumeration").Create(&enums[index])
	}

	rows, err := db.Table("enumeration").Where("enum_type.code=?", "regionEnumType").
		Joins("left join enum_type on enum_type.id=enumeration.enum_type_id").
		//select 代表要查詢的欄位
		Select("enumeration.code,enumeration.name,enum_type.name,enum_type.code").Rows()

	if err != nil {
		fmt.Print(err)
	}

	defer rows.Close()
	newEnumType := &projectModels.EnumType{}
	newEnumType.Enums = make([]projectModels.Enumeration, 0)

	for rows.Next() {
		enum := projectModels.Enumeration{}
		//scan 代表select 參數對應
		err = rows.Scan(&enum.Code, &enum.Name, &newEnumType.Code,
			&newEnumType.Name)

		if err != nil {
			fmt.Print(err)
		}
		newEnumType.Enums = append(newEnumType.Enums, enum)
	}
	fmt.Println(newEnumType)
}
