package util

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// DB ip
	DB_IP = "mongodb+srv://engineers:MiS2UvU6dEVIsgky@cluster0.gyixj.gcp.mongodb.net/?retryWrites=true&w=majority"
)

var mgoCli *mongo.Client

func initEngine() {
	var err error
	clientOptions := options.Client().ApplyURI(DB_IP)

	mgoCli, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
}
func GetMgoCli() *mongo.Client {
	if mgoCli == nil {
		initEngine()
	}
	return mgoCli
}
