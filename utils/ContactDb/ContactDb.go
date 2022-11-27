package ContactDb

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var DB *sqlx.DB

func ContactDb() (db *sqlx.DB) {
	if DB != nil {
		return
	}

	config := ReadYamlConfig()
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/sine", config.Mysql.UserName, config.Mysql.Pwd, config.Mysql.Host, config.Mysql.Port)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Println(err)
		return
	}
	DB = db

	return DB
}
