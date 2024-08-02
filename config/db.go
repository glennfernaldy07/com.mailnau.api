package config

import _ "github.com/go-sql-driver/mysql"
import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/opentracing/opentracing-go/log"
	"gopkg.in/jinzhu/gorm.v1"
	"time"
)

func NewMysqlDB(cfg Config) *gorm.DB {
	driverName := "mysql"
	username := cfg.GetString(DBUsername)
	password := cfg.GetString(DBPassword)
	host := cfg.GetString(DBHost)
	port := cfg.GetString(DBPort)
	dbname := cfg.GetString(DBName)
	parsetime := cfg.GetString(DBParsetime)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=%s", username, password, host, port, dbname, parsetime)

	db, err := gorm.Open(driverName, dsn)
	db.LogMode(cfg.GetBool(EnableDatabaseLog))
	db.BlockGlobalUpdate(true)
	if err != nil {
		log.Error(errors.New(fmt.Sprintf("cannot connect to DB, err=%s", err)))
	}

	//configureConnectionPool(db.DB())
	return db
}

func configureConnectionPool(dbPool *sql.DB) {
	// Set maximum number of connections in idle connection pool.
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	dbPool.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	dbPool.SetMaxOpenConns(30)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	dbPool.SetConnMaxLifetime(time.Hour)
}
