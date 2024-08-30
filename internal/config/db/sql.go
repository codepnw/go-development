package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetClientGorm() *gorm.DB {
	var db *gorm.DB
	var err error

	dns := fmt.Sprintf(
		"%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DBNAME"),
	)

	if db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		SkipDefaultTransaction: true,
	}); err != nil {
		panic(err)
	}

	if sql, err := db.DB(); err == nil {
		sql.SetMaxIdleConns(100)
		sql.SetMaxOpenConns(50)
		sql.SetConnMaxLifetime(time.Hour)
	}

	return db
}
