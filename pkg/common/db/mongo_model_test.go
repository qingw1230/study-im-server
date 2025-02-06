package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/qingw1230/study-im-server/pkg/common/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMongoInsert(t *testing.T) {
	uri := fmt.Sprintf("mongodb://%s/%s?maxPoolSize=%d", config.Config.Mongo.DBAddress[0], "testDB", config.Config.Mongo.DBMaxPoolSize)
	clientOpts := options.Client().ApplyURI(uri)
	clientOpts.Auth = &options.Credential{
		Username:   config.Config.Mongo.DBUserName,
		Password:   config.Config.Mongo.DBPassword,
		AuthSource: "admin",
	}
	mongoClient, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		t.Fatal(err.Error())
	}

	coll := mongoClient.Database("testDB").Collection("testColl")
	result, err := coll.InsertOne(
		context.TODO(),
		bson.D{
			{Key: "item", Value: "canvas"},
			{Key: "size", Value: bson.D{
				{Key: "h", Value: 28},
				{Key: "w", Value: 35.5},
			}},
		})
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(result)
}
