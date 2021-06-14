package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		name := context.Query("name")
		fmt.Println(name)
		context.Writer.Write([]byte("hello" + name))
	})
	engine.POST("/login", func(context *gin.Context) {
		fmt.Print(context.FullPath())
		name, ok := context.GetPostForm("name")
		if ok {
			fmt.Println(name)
		}
		pwd, ok := context.GetPostForm("pwd")
		if ok {
			fmt.Println(pwd)
		}
		context.Writer.Write([]byte("name=" + name + ";pwd=" + pwd + "正在登陆..."))
	})
	engine.DELETE("/user/:id", func(context *gin.Context) {
		user_id := context.Param("id")
		context.Writer.Write([]byte("user_id=" + user_id + "的用户被删除了"))
	})
	engine.Run()
}
