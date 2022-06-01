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

// NewDB returns mysql driver based *sqlx.DB
func NewDB(e *echo.Echo) *sqlx.DB {
	dsn := os.Getenv("DSN")
	driver := os.Getenv("DRIVER")

	// DBオープン
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		e.Logger.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("db connection succeeded")

	return db
}
