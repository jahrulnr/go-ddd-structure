package mysqldb

import (
	"database/sql"
	"ddd-impl/config"
	"github.com/jmoiron/sqlx"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func New(cfg *config.Config) *sql.DB {
	db, err := sql.Open(cfg.Mysql.Driver, cfg.Mysql.Host)
	if err != nil {
		log.Fatalf("couldn't connect to database connection: %v", err.Error())
		return nil
	}

	db.SetMaxIdleConns(cfg.Mysql.MaxIdleConns)
	db.SetMaxOpenConns(cfg.Mysql.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(cfg.Mysql.MaxLifetimeConns))
	db.SetConnMaxIdleTime(time.Duration(cfg.Mysql.MaxIdleTime))

	return db
}

func NewSqlx(cfg *config.Config) *sqlx.DB {
	db, err := sqlx.Open(cfg.Mysql.Driver, cfg.Mysql.Host)
	if err != nil {
		log.Fatalf("couldn't connect to database connection: %v", err.Error())
		return nil
	}

	db.SetMaxIdleConns(cfg.Mysql.MaxIdleConns)
	db.SetMaxOpenConns(cfg.Mysql.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(cfg.Mysql.MaxLifetimeConns))
	db.SetConnMaxIdleTime(time.Duration(cfg.Mysql.MaxIdleTime))

	return db
}
