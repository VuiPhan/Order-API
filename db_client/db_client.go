package db_client

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DBClient *sqlx.DB

func InitialConnection() {
	db, err := sqlx.Open("mysql", "root:password@tcp(localhost:3306)/OnlineShop?parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	DBClient = db
}
