/*
 * Copyright (c) 2018. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */

package db

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/**
数据原配置
*/
type dbConfig struct {
	Dsn string
}

func TestDb() {
	dbw := dbConfig{
		Dsn: "admin:admin@tcp(127.0.0.1:3306)/test",
	}
	db, err := sql.Open("mysql",
		dbw.Dsn)
	if err != nil {
		panic(err)
		return
	}
	rows, err := db.Query("select version()")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var version string
		if err := rows.Scan(&version); err == nil {
			fmt.Println(version)
		}
	}
	defer db.Close()
}
