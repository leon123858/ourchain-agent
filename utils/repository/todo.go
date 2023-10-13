package repository

import (
	"context"
	"fmt"

	model "github.com/leon123858/go-aid/utils/modal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateTodo(db *mongo.Database, todo model.Todo) (string, error) {
	iResult, err := db.Collection("todo").InsertOne(context.TODO(), model.Todo{
		Aid:       todo.Aid,
		Title:     todo.Title,
		Completed: todo.Completed,
	})
	if err != nil {
		return "", err
	}
	id := iResult.InsertedID.(primitive.ObjectID)
	return id.Hex(), nil
}

func GetTodoList(db *mongo.Database, aid string) (*[]model.Todo, error) {
	var todoList []model.Todo
	cur, err := db.Collection("todo").Find(context.TODO(), bson.M{
		"aid": aid,
	})
	if err != nil {
		return nil, err
	}
	for cur.Next(context.Background()) {
		var todo model.Todo
		err := cur.Decode(&todo)
		if err != nil {
			fmt.Print(err)
			return nil, err
		}
		todoList = append(todoList, todo)
	}
	return &todoList, nil
}

func CheckTodo(db *mongo.Database, _id string, isChecked bool) error {
	// check todo by id
	objId, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return err
	}
	uResult, err := db.Collection("todo").UpdateOne(context.TODO(),
		bson.M{"_id": objId},
		bson.M{"$set": bson.M{"completed": isChecked}})
	if err != nil {
		return err
	}
	if uResult.ModifiedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func DeleteTodoById(db *mongo.Database, _id string) error {
	objId, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return err
	}
	dResult, err := db.Collection("todo").DeleteOne(context.TODO(), bson.M{
		"_id": objId,
	})
	if err != nil {
		return err
	}
	if dResult.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
