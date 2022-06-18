package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

func Load() *sql.DB {
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USERNAME")
	pwd := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")

	instance := fmt.Sprintf("%s:%v@(%s:%v)/%s?parseTime=true", user, pwd, host, port, dbName)

	db, err := sql.Open("mysql", instance)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute)
	db.SetMaxOpenConns(32)

	if db.Ping() == nil {
		fmt.Println("database connected ...")
	}


	return db
}