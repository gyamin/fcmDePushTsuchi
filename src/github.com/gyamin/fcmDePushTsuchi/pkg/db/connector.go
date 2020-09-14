package db

import (
	"database/sql"
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gyamin/fcmDePushTsuchi/config"
)

func GetConnection() *sql.DB {
	var conf config.Config
	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		panic(err)
	}

	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.DB.User, conf.DB.Password, conf.DB.Server, conf.DB.Port, conf.DB.DbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
