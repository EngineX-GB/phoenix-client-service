package datasource

import "database/sql"

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/db_phoenix?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}
