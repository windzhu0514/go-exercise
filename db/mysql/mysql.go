package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// var (
// 	username = "root"
// 	userpass = "123456"
// 	address  = "106.12.58.250:3306"
// 	dbname   = "busticket"
// )

var (
	username = "root"
	userpass = "windzhu0514"
	address  = "106.14.28.102:9306"
	dbname   = "busticket"
)

func main() {
	// 查询结果放入切片
	values := make([]sql.NullString, 25)
	valuesP := make([]interface{}, 25)
	for i := 0; i < 25; i++ {
		valuesP[i] = &values[i]
	}

	dsn := username + ":" + userpass + "@tcp(" + address + ")/" + dbname + "?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
		return
	}

	err = db.QueryRow("SELECT * FROM ticket_site_info WHERE id=?", 1).Scan(valuesP...)
	if err != nil {
		log.Println(err)
		return
	}

	for _, v := range values {
		fmt.Println(v.Value())
	}
}
