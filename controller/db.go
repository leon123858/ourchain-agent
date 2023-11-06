package controller

import (
	"github.com/labstack/echo"
	"github.com/leon123858/go-aid/modal"
	"github.com/leon123858/go-aid/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

//db := util.GetMgoCli().Database("todo")

func CreateTodoController(db *mongo.Database, which string) func(ctx echo.Context) error {
	switch which {
	case "create":
		return func(ctx echo.Context) error {
			todo := new(model.Todo)
			err := ctx.Bind(&todo)
			if err != nil {
				return err
			}
			if id, err := repository.CreateTodo(db, *todo); err != nil {
				return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
					"result": "fail",
					"error":  err.Error(),
				})
			} else {
				return ctx.JSON(http.StatusOK, map[string]interface{}{
					"result": "success",
					"data":   id,
				})
			}
		}
	case "get":
		return func(ctx echo.Context) error {
			aid := ctx.QueryParam("aid")
			if todoList, err := repository.GetTodoList(db, aid); err != nil {
				return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
					"result": "fail",
					"error":  err.Error(),
				})
			} else {
				return ctx.JSON(http.StatusOK, map[string]interface{}{
					"result": "success",
					"data":   *todoList,
				})
			}
		}
	case "update":
		return func(ctx echo.Context) error {
			todo := new(model.Todo)
			err := ctx.Bind(&todo)
			if err != nil {
				return err
			}
			if err := repository.CheckTodo(db, todo.ID, todo.Completed); err != nil {
				return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
					"result": "fail",
					"error":  err.Error(),
				})
			} else {
				return ctx.JSON(http.StatusOK, map[string]interface{}{
					"result": "success",
					"data":   "",
				})
			}
		}
	case "delete":
		return func(ctx echo.Context) error {
			todo := new(model.Todo)
			err := ctx.Bind(&todo)
			if err != nil {
				return err
			}
			if err := repository.DeleteTodoById(db, todo.ID); err != nil {
				return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
					"result": "fail",
					"error":  err.Error(),
				})
			} else {
				return ctx.JSON(http.StatusOK, map[string]interface{}{
					"result": "success",
					"data":   "",
				})
			}
		}
	default:
		return nil
	}
}
