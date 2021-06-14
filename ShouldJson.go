package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.POST("/add", func(context *gin.Context) {
		var person Person
		err := context.BindJSON(&person)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(person.Name)
		context.Writer.Write([]byte("新增学生：" + person.Name))
	})

	engine.Run()
}

type Person struct {
	Name string `form:name`
	Sex  string `form:sex`
	Age  string `form:age`
}
