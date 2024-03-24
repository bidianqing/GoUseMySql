package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dataSourceName := "root:root@tcp(127.0.0.1:3306)/demo?charset=utf8mb4"
	// Open函数可能只是验证其参数格式是否正确，实际上并不创建与数据库的连接。
	// 返回的DB对象可以安全地被多个goroutine并发使用，并且维护其自己的空闲连接池。因此，Open函数应该仅被调用一次，很少需要关闭这个DB对象。
	db, err := sql.Open("mysql", dataSourceName)
	failOnError(err)

	// 如果要检查数据源的名称是否真实有效，应该调用Ping方法
	if err := db.Ping(); err != nil {
		failOnError(err)
	}
	defer db.Close()

	// 查询一条数据
	var user User
	row := db.QueryRow("select * from tb_user where id = 1")
	err = row.Scan(&user.Id, &user.Name)
	failOnError(err)
	fmt.Printf("%v \n", user)

	// 查询记录总数
	var count int
	rowCount := db.QueryRow("select count(*) from tb_user")
	err = rowCount.Scan(&count)
	failOnError(err)
	fmt.Println(count)

	// 查询多条记录
	users := make([]User, 0, count)
	rows, err := db.Query("select * from tb_user")
	failOnError(err)
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Name)
		failOnError(err)
		users = append(users, user)
	}
	rows.Close()
	fmt.Println(users)

	// 插入多条数据
	// INSERT INTO test(n1, n2, n3) VALUES(v1, v2, v3),(v4, v5, v6),(v7, v8, v9);
	// https://stackoverflow.com/a/21108084
	insertRows := []User{
		{Name: "hi"},
		{Name: "hello"},
	}
	var insertSql strings.Builder
	insertSql.WriteString("insert into tb_user(Name) values")
	args := []interface{}{}
	for index, row := range insertRows {
		if index == len(insertRows)-1 {
			insertSql.WriteString("(?)")
		} else {
			insertSql.WriteString("(?),")
		}
		args = append(args, row.Name)
	}
	_, err = db.Exec(insertSql.String(), args...)
	failOnError(err)
}

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

type User struct {
	Id   int
	Name string
}
