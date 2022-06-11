package db

import (
	_ "embed"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

//go:embed schema.sql
var schema string

const db_env = "prod"

// NewDB returns mysql driver based *sqlx.DB
func NewDB(e *echo.Echo) *sqlx.DB {
	dbDriver := "mysql"

	dsn := generateDsn()

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

// 本番環境なら環境変数からDB情報を取得する
// TODO: Getenvして空ならデフォルト値を、
// 値が入っていたらその値を返す関数を作って呼び出す
func generateDsn() string {

	// ローカル環境でのDB接続情報
	dbName := "go_db"
	dbUser := "go_user"
	dbPass := "password"
	dbAddress := "db"
	fmt.Println(dbName, dbUser, dbPass, dbAddress)

	var dsn string

	// 本番環境に接続する場合
	// if os.Getenv("DB_ENV") == db_env {
	// 	dbName = os.Getenv("MYSQL_DATABASE")
	// 	dbUser = os.Getenv("MYSQL_USER")
	// 	dbPass = os.Getenv("MYSQL_PASSWORD")
	// }
	// DSNの設定
	dsn = dbUser + ":" + dbPass + "@tcp(" + dbAddress + ":3306)/" + dbName + "?parseTime=true&autocommit=0&sql_mode=%27TRADITIONAL,NO_AUTO_VALUE_ON_ZERO,ONLY_FULL_GROUP_BY%27"

	return dsn
}
