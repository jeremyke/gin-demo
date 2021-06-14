package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	constr := "root:root123@tcp(localhost:3306)/mytest"
	db, err := sql.Open("mysql", constr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	//add
	//_ , err = db.Exec("insert into user (name) values(?);","klp")
	//if err == nil {
	//	fmt.Println("插入成功")
	//}else {
	//	fmt.Println(err.Error())
	//	return
	//}
	//query
	row, query_err := db.Query("select * from user")
	if query_err != nil {
		log.Println(query_err.Error())
		return
	}
	for {
		if row.Next() {
			user := new(user)
			row_err := row.Scan(&user.Id, &user.Name)
			if row_err != nil {
				log.Fatal(row_err.Error())
				return
			}
			fmt.Println(user.Id, user.Name)
		} else {
			break
		}
	}

}

type user struct {
	Id   int
	Name string
}
