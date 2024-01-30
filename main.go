package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/demo?charset=utf8mb4")
	if err != nil {
		fmt.Printf("Open failed, err:%v\n", err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("Ping failed, err:%v\n", err)
		return
	}

	// https://stackoverflow.com/a/21108084
	result, err := db.Exec(`insert into tb_menu(Name,ParentId,Icon,Path) values ('2.3', 2, null, null)`)
	if err != nil {
		fmt.Printf("Exec failed, err:%v\n", err)
		return
	}

	fmt.Println(result)
}
