package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

// Init init the database from config to database connection
func Init(cfg DBConfig) (*sqlx.DB, error) {
 	// toDNS := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Asia%2FJakarta&charset=utf8&autocommit=false", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	toDNS := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	db, err := sqlx.Connect("mysql", toDNS)
	if err != nil {
		return nil, err
	}
	return db, nil
}
