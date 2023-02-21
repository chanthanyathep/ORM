package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n=============================\n", sql)
}

func initDB() *gorm.DB {
	dsn := "root:1234@tcp(localhost:3306)/showroom?parseTime=true"
	dial := mysql.Open(dsn)

	var err error
	db, err := gorm.Open(dial, &gorm.Config{
		//Logger: &SqlLogger{},
	})
	if err != nil {
		panic(err)
	}
	return db
}
