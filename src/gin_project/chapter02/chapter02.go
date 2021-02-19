package chapter02

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id   string `form:"id"`
	Name string `form:"name"`
	Old  int    `form:"old"`
}

const (
	Suc = "success"
	Err = "error"
)

func user(ctx *gin.Context) {
	name := "sheldonTest"
	ctx.HTML(200, "index.html", name)
}

func UserStruct(ctx *gin.Context) {
	userInfo := User{
		Id:   "安安",
		Name: "sheldon",
		Old:  10,
	}

	ctx.HTML(http.StatusOK, "user/user.html", userInfo)
}

func ArrStruct(ctx *gin.Context) {
	arrTest := [3]int{2, 3, 5}

	ctx.HTML(http.StatusOK, "arrTest/arr.html", arrTest)
}

func ArrUserStruct(ctx *gin.Context) {
	arrTest := [3]User{
		{Id: "安安",
			Name: "sheldon",
			Old:  10},
		{Id: "安安1",
			Name: "sheldon1",
			Old:  11},
		{Id: "安安2",
			Name: "sheldon2",
			Old:  12},
	}

	ctx.HTML(http.StatusOK, "arrTest/arrUser.html", arrTest)
}

func MapStruct(ctx *gin.Context) {
	mapStruct := map[string]string{
		"name":  "sheldon",
		"name2": "sheldon2",
		"name3": "sheldon3",
	}

	ctx.HTML(http.StatusOK, "mapTest/map.html", mapStruct)

}

func ParamStruct(ctx *gin.Context) {
	//
	id := ctx.Param("id")
	ctx.String(http.StatusOK, "hello %s", id)

}

func GetQuery(ctx *gin.Context) {

	id := ctx.Query("id")

	ctx.String(http.StatusOK, "hello %s", id)

}

func PostPage(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "postTest/postForm.html", nil)

}

func PostParam(ctx *gin.Context) {

	id := ctx.PostForm("id")
	password := ctx.PostForm("password")

	fmt.Println("id:" + id)
	fmt.Println("password:" + password)

	ctx.String(http.StatusOK, "添加成功", nil)

}

func BindForm(ctx *gin.Context) {

	var user User
	err := ctx.ShouldBind(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)

	ctx.String(http.StatusOK, "添加成功BindForm", nil)

}

func NoSetting(ctx *gin.Context) {
	data := gin.H{
		"msg":  "undefined path!",
		"code": Err,
	}
	ctx.JSON(http.StatusNotFound, data)
}
