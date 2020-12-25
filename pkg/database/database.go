package database

import (
	"database/sql"
	"goblog/pkg/logger"
	"time"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Initialize() {
	initDB()
	createTables()
}

func initDB() {
	var err error

	config := mysql.Config{
		User:                 "homestead",
		Passwd:               "secret",
		Addr:                 "192.168.10.10:3306",
		Net:                  "tcp",
		DBName:               "goblog",
		AllowNativePasswords: true,
	}

	DB, err = sql.Open("mysql", config.FormatDSN())
	logger.LogError(err)

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	DB.SetConnMaxIdleTime(5 * time.Minute)

	err = DB.Ping()
	logger.LogError(err)
}

func createTables() {
	createArticlesSql := `CREATE TABLE IF NOT EXISTS articles(
		id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		title varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
		content longtext COLLATE utf8mb4_unicode_ci
	);`
	_, err := DB.Exec(createArticlesSql)
	logger.LogError(err)
}
