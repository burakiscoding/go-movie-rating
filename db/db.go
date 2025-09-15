package db

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

func NewSQL() (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "movie_rating"
	cfg.ParseTime = true

	return sql.Open("mysql", cfg.FormatDSN())
}
