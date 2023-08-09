package driver

import (
	"fmt"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	mysqlConfig "go-module/config/mysql"
)

type MySqlDB struct {
	SQL *sql.DB
}

var MySQL = &MySqlDB{}

func Connect() (*MySqlDB) {
	addString := mysqlConfig.HOST_URL + ":" + mysqlConfig.PORT
	// Specify connection properties.
	cfg := mysql.Config{
		User:   mysqlConfig.USERNAME,
		Passwd: mysqlConfig.PASSWORD,
		Net:    "tcp",
		Addr:   addString,
		DBName: mysqlConfig.DB_NAME,
	}

	// Get a database handle.
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	// connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
	// 						host, port, user, password, dbname)
	// 						fmt.Println(connStr)
	// db, err := sql.Open("mysql", connStr) // database driver name, connection string
	// if err != nil {
	// 	// return nil, err
	// 	panic(err)
	// }
	
	fmt.Println("Connection ok")
	MySQL.SQL = db
	return MySQL
}