package db

import (
	"context"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/qingw1230/study-im-server/pkg/common/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB DataBases

type DataBases struct {
	MySQLDB     mysqlDB
	redisPool   *redis.Pool
	mongoClient *mongo.Client
}

// key 将 address name 连接形成键
func key(dbAddress, dbName string) string {
	return dbAddress + "_" + dbName
}

func init() {
	initMySQLDB()

	var mongoClient *mongo.Client
	uri := fmt.Sprintf("mongodb://%s:%s@%s/%s/?maxPoolSize=%d",
		config.Config.Mongo.DBUserName,
		config.Config.Mongo.DBPassword,
		config.Config.Mongo.DBAddress[0],
		config.Config.Mongo.DBDatabase,
		config.Config.Mongo.DBMaxPoolSize)
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err.Error())
	}
	DB.mongoClient = mongoClient

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
