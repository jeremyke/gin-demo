package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	engine.GET("/hello_json", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg":  "成功",
			"data": context.FullPath(),
		})
	})
	engine.GET("/hello_struct", func(context *gin.Context) {
		resp := response{
			Code: 1,
			Msg:  "请求成功",
			Data: context.FullPath(),
		}
		context.JSON(200, &resp)
	})
	engine.Run()
}

type response struct {
	Code int
	Msg  string
	Data interface{}
}
