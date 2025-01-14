package db

import (
	"github.com/gomodule/redigo/redis"
)

var DB DataBases

type DataBases struct {
	MySQLDB   mysqlDB
	redisPool *redis.Pool
}

// key 将 address name 连接形成键
func key(dbAddress, dbName string) string {
	return dbAddress + "_" + dbName
}

func init() {
	initMySQLDB()

	DB.redisPool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				"127.0.0.1:16379",
				redis.DialDatabase(0),
				redis.DialPassword("qin1002."),
			)
		},
	}
}
