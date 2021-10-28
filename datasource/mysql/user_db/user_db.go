package user_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
)

//TODO use yaml file and get the environmental variables using viper https://github.com/spf13/viper
func init() {
	var err error
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", "root", "root", "127.0.0.1:3306", "users_db")
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("successfull connected to database")
}
