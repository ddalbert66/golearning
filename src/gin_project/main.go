package main

import (
	"goLearning20200930/src/gin_project/chapter01"
	"goLearning20200930/src/gin_project/chapter02"

	"github.com/gin-gonic/gin"
)

func main() {
	//1.取得router
	// new 與 default的差別
	router := gin.Default()
	//router :=gin.New()

	router.NoMethod(chapter02.NoSetting)
	router.NoRoute(chapter02.NoSetting)
	//設定好資料夾路徑就可以找到此路徑下的所有文件
	router.LoadHTMLGlob("htmlTemplateFile/**/*")
	//2.router get 方法
	router.GET("/", chapter01.Hello)
	router.GET("/user", chapter02.UserStruct)
	router.GET("/arr", chapter02.ArrStruct)
	router.GET("/arrUser", chapter02.ArrUserStruct)
	router.GET("/map", chapter02.MapStruct)
	router.GET("/param/:id", chapter02.ParamStruct)
	router.GET("/query", chapter02.GetQuery)
	router.GET("/postPage", chapter02.PostPage)
	router.POST("/postParam", chapter02.PostParam)
	router.POST("/bindForm", chapter02.BindForm)

	//需設定好全部的html文件 如果有100個就要設定100個html文件
	//router.LoadHTMLFiles("index.html")
	//3.router post 方法
	//router.POST(relativePath string, handlers ...gin.HandlerFunc)
	//4.router 執行
	//可指定host 跟 port
	router.Run(":9000")
}
