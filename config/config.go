package config

import (
	"database/sql"
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dsn = fmt.Sprintf("%v:%v@(%v)/%v", goDotEnvVariable("username"), goDotEnvVariable("password"), goDotEnvVariable("host"), goDotEnvVariable("db_name"))
)

func MySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")
  
	if err != nil {
		log.Fatal("Error loading .env file")
		return ("Error")
	}
  
	return os.Getenv(key)
  }