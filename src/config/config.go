package config

import(
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

var (
	db *sql.DB
)

func Connect(){
	d, err := sql.Open("mysql", "test:deez@tcp(127.0.0.1:3306)/takehome")
	if err != nil {
        panic(err.Error())
    }
	db = d	
}

func GetDB() *sql.DB{
	return db
}