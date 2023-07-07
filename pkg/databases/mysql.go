package databases

import (
	"database/sql"
	"fmt"
	"go_cleanarc/configs"
	"go_cleanarc/pkg/utils"

	_ "github.com/go-sql-driver/mysql"
)

// DB is a global var for connect DB
var DB *sql.DB

func NewDBConnection(cfg *configs.DBConfig) (*sql.DB, error) {
	mariaDbUrl, err := utils.ConnectionUrlBuilder("mariaDB", cfg)

	if err != nil {
		return nil, err
	}

	DB, err = sql.Open("mysql", mariaDbUrl)

	fmt.Println(mariaDbUrl)

	if err != nil {
		panic(fmt.Errorf("error in sql.Open: %w", err))
	}
	return DB, nil
}
