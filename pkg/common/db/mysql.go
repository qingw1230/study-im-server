package db

import (
	"fmt"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Test mysqlDB

type mysqlDB struct {
	sync.RWMutex
	dbMap map[string]*gorm.DB
}

func initMySQLDB() {
	dsn := "root:qin1002.@tcp(127.0.0.1:13306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		time.Sleep(time.Duration(10) * time.Second)
		db, err = gorm.Open("mysql", dsn)
		if err != nil {
			panic(err.Error())
		}
	}

	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s default charset utf8 COLLATE utf8_general_ci;", "studyim")
	err = db.Exec(sql).Error
	if err != nil {
		panic(err.Error())
	}
	db.Close()

	dsn = "root:qin1002.@tcp(127.0.0.1:13306)/studyim?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(
		&UserInfo{},
		&UserInfoBeauty{},
	)
	db.Set("gorm:table_options", "CHARSET=utf8mb4")
	db.Set("gorm:table_options", "collation=utf8mb4_general_ci")
}
