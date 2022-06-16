package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("start migration ....")

	_ = godotenv.Load()

	host := os.Getenv("MYSQL_HOST")
	port := "3306"
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")

	// user:password@tcp(host:port)/dbname?multiStatements=true
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", user, password, host, port, dbName)

	db, _ := sql.Open("mysql", source)
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"schema",
		"mysql",
		driver,
	)

	if err := m.Up(); err != nil {
		log.Fatalf("error up migration: %v", err)
	}

	fmt.Println("finish migrate!")
}
