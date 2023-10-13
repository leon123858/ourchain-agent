package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	model "github.com/leon123858/go-aid/utils/modal"
	util "github.com/leon123858/go-aid/utils/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	db := util.GetMgoCli().Database("todo")
	aid := uuid.New().String()
	// create a todo
	iResult, err := db.Collection("todo").InsertOne(context.TODO(), model.Todo{
		Aid:       aid,
		Title:     "title",
		Completed: false,
	})
	if err != nil {
		fmt.Print(err)
		return
	}
	id := iResult.InsertedID.(primitive.ObjectID)
	fmt.Println("_id", id.Hex())

	// check a todo
	objId, err := primitive.ObjectIDFromHex(id.Hex())
	if err != nil {
		fmt.Println(err)
		return
	}
	uResult, err := db.Collection("todo").UpdateOne(context.TODO(), bson.M{"_id": objId}, bson.M{"$set": bson.M{"completed": true}})
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("uResult", uResult.ModifiedCount)

	// get all todos
	var todoList []model.Todo
	cur, err := db.Collection("todo").Find(context.TODO(), bson.M{
		"aid": aid,
	})
	if err != nil {
		fmt.Print(err)
		return
	}
	for cur.Next(context.Background()) {
		var todo model.Todo
		err := cur.Decode(&todo)
		if err != nil {
			fmt.Print(err)
			return
		}
		todoList = append(todoList, todo)
	}
	fmt.Println("todoList", todoList)
	_, err = primitive.ObjectIDFromHex(todoList[0].ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	// delete a todo
	objId, err = primitive.ObjectIDFromHex(id.Hex())
	if err != nil {
		fmt.Println(err)
		return
	}
	dResult, err := db.Collection("todo").DeleteOne(context.TODO(), bson.M{
		"_id": objId,
	})
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Println("dResult", dResult.DeletedCount)
}
