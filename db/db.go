package db

import (
	_ "embed"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

//go:embed schema.sql
var schema string

// NewDB returns mysql driver based *sqlx.DB
func NewDB(e *echo.Echo) *sqlx.DB {

	// ローカル環境でのDB接続情報
	dbDriver := "mysql"
	dbName := "go-db"
	dbUser := "go-user"
	dbPass := "password"
	dbAddress := "db"

	// 本番環境なら環境変数からDB情報を取得する
	// if os.Getenv("DB_ENV") == "prod" {
	// 	dbName = os.Getenv("MYSQL_DATABASE")
	// 	dbUser = os.Getenv("MYSQL_USER")
	// 	dbPass = os.Getenv("MYSQL_PASSWORD")
	// }

	// DSNの設定
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbAddress + ":3306)/" + dbName + "?parseTime=true&autocommit=0&sql_mode=%27TRADITIONAL,NO_AUTO_VALUE_ON_ZERO,ONLY_FULL_GROUP_BY%27"

	// DBオープン
	db, err := sqlx.Open(dbDriver, dsn)
	if err != nil {
		e.Logger.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("db connection succeeded")

	return db
}
