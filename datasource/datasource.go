package datasource

import (
	"database/sql"
	"phoenix-client-service/config"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", config.GetDataSourceUsername()+":"+config.GetDataSourcePassword()+"@tcp("+config.GetDataSourceHostname()+")/"+config.GetDataSourceName()+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}
