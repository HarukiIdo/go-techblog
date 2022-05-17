package db

import (
	_ "embed"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

//go:embed schema.sql
var schema string

// DBと接続
func ConnectDB(e *echo.Echo) *sqlx.DB {
	dsn := os.Getenv("DSN")
	driver := os.Getenv("DRIVER")
	log.Println(dsn)
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("db connection succeeded")

	// テーブルを作成する
	if _, err := db.Exec(schema); err != nil {
		return nil
	}
	return db
}
