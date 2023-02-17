package config

import(
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"fmt"
)

var (
	db *sql.DB
)

func Connect(){
	password := goDotEnvVariable("MYSQLPASSWORD")
	d, err := sql.Open("mysql", fmt.Sprintf("test:%s@tcp(127.0.0.1:3306)/takehome", password))
	if err != nil {
        panic(err.Error())
    }
	db = d
}

func GetDB() *sql.DB{
	return db
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load("../.env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
  }