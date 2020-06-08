package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*
对数据库操作需要使用到database/sql，但是数据库连接驱动官方没有提供，需要使用第三方驱动包
*/
func init() {
	open, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(&open)
}
func main() {

}
