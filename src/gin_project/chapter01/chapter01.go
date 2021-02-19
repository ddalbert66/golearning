package chapter01

import "github.com/gin-gonic/gin"

func Hello(ctx *gin.Context) {
	name := "sheldonTest"
	ctx.HTML(200, "index.html", name)
}