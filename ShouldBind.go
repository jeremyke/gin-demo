package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()
	engine.POST("/register", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var register Register
		err := context.ShouldBind(&register)
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		fmt.Println(register.Name + register.Phone)
		context.Writer.Write([]byte("hello," + register.Name))
	})
	engine.Run()
}

type Register struct {
	Name  string `form:"name"`
	Phone string `form:"phone"`
	Pwd   string `form:"pwd"`
}
