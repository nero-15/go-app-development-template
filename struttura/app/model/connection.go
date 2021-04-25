package models

import (
	"database/sql"
	"fmt"
	"log"
	"struttura/config"

	_ "github.com/go-sql-driver/mysql"
)

var DbConnection *sql.DB

func init() {
	var err error
	db, _ := fmt.Printf("%v:%v@/%v", config.Config.DbHost, config.Config.DbPassword, config.Config.DbName)
	DbConnection, err = sql.Open(config.Config.SQLDriver, string(db))
	if err != nil {
		log.Fatalln(err)
	}

}
