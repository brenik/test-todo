package config

import (
	"os"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetDatabase() (database *sql.DB, err error) {
	database, err = sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:3306)/%s",
			os.Getenv("USER"),
			os.Getenv("PASSWORD"),
			os.Getenv("HOST"),
			os.Getenv("DATABASE"),
		),
	)

	return
}