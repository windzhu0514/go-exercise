package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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

type LineInfo struct {
	SiteId         int    `db:"site_id"`
	Departure      string `db:"dpt_city_name,omitempty"`
	DptEnName      string `db:"dpt_city_py,omitempty"`
	DptCode        string `db:"dpt_city_code,omitempty"`
	DptStationName string `db:"dpt_station_name,omitempty"`
	DptStationCode string `db:"dpt_station_code,omitempty"`
	DestName       string `db:"dest_name,omitempty"`
	DestEnName     string `db:"dest_py,omitempty"`
	DestCode       string `db:"dest_code,omitempty"`
	price          float64
}

func main() {
	count := 0
	var pageId, step = 0, 2 //控制数据库分页
	db, err := sqlx.Connect("mysql", "root:123456@tcp(127.0.0.1:3306)/busticket?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		sqlStr := fmt.Sprintf(`
SELECT site_id, dpt_city_name, dpt_city_py, dpt_station_name, dpt_station_code, dest_name, dest_py, dest_code
FROM site_bus_lines
WHERE site_id = %d
AND dpt_city_name='%s' AND disabled='0' LIMIT %d,%d`, 20, "合肥", pageId, step)

		i := 0

		rows, err := db.Queryx(sqlStr)
		if err != nil {
			fmt.Println(err)
			return
		}

		for rows.Next() {
			i++

			var lineInfo LineInfo
			err := rows.StructScan(&lineInfo)
			if err != nil {
				rows.Close()
				fmt.Println(err)
				return
			}

			fmt.Println(lineInfo)
		}
		if i < 2 {
			break
		}
		rows.Close()

		fmt.Println(rows.Err())
		fmt.Println(rows.NextResultSet())

		count += i

		pageId = pageId + step
	}

	fmt.Println("over")
}
