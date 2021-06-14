package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()

	engine.Handle("GET", "/hello", func(context *gin.Context) {
		path := context.FullPath()
		fmt.Print(path)

		name := context.DefaultQuery("name", "hello")
		context.Writer.Write([]byte("hello," + name))

	})

	engine.Handle("POST", "/login", func(context *gin.Context) {
		path := context.FullPath()
		fmt.Print(path)

		name := context.PostForm("name")
		pwd := context.PostForm("pwd")
		context.Writer.Write([]byte("name=" + name + ";pwd=" + pwd + "正在登陆..."))
	})
	if err := engine.Run(":9090"); err != nil {
		log.Fatal(err.Error())
	}
}
