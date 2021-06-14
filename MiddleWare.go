package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.Use(requestInfo())
	engine.GET("/query", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"data": context.FullPath(),
		})
	})
	engine.Run()
}
func requestInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.FullPath()
		method := context.Request.Method
		fmt.Println("path=", path)
		fmt.Println("method=", method)
		context.Next()
		fmt.Println("请求状态码：", context.Writer.Status())
	}
}
