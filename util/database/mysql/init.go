package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

// Init init the database from config to database connection
func Init(cfg DBConfig) (*sqlx.DB, error) {
	// toDNS := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Asia%2FJakarta&charset=utf8&autocommit=false", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	var toDNS string
	switch cfg.Dialect {
	case "mysql":
		toDNS = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	case "postgres":
		toDNS = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	default:
		toDNS = ""
	}
	if toDNS == "" {
		return nil, fmt.Errorf("Database support only mysql / postgres")
	}
	db, err := sqlx.Connect(cfg.Dialect, toDNS)
	if err != nil {
		return nil, err
	}
	return db, nil
}
