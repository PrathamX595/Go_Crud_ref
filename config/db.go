package config

import (
	"context"
	"crud/utils"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var connectionStr string = GetEnv("MONGO")

const dbName = "user"
const colName = "goCRUD"

var Col *mongo.Collection

func Db() {
	options := options.Client().ApplyURI(connectionStr)
	client, err := mongo.Connect(context.TODO(), options)
	utils.CheckErr(err)
	fmt.Println("mongo connection success")
	Col = client.Database(dbName).Collection(colName)
	fmt.Println("collection instance is ready, ", Col)
}
