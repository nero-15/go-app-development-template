package models

import (
	"database/sql"
	"log"
	"struttura/config"

	_ "github.com/go-sql-driver/mysql"
)

var DbConnection *sql.DB

func init() {
	var err error
	DbConnection, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

}
