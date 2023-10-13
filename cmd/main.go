package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	model "github.com/leon123858/go-aid/utils/modal"
	util "github.com/leon123858/go-aid/utils/mongo"
	"github.com/leon123858/go-aid/utils/repository"
)

func main() {
	db := util.GetMgoCli().Database("todo")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.POST("/create/todo", func(ctx echo.Context) error {
		todo := new(model.Todo)
		ctx.Bind(&todo)
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
	})
	e.GET("/get/todo", func(ctx echo.Context) error {
		aid := ctx.QueryParam("aid")
		println(aid)
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
	})
	e.POST("/update/todo", func(ctx echo.Context) error {
		todo := new(model.Todo)
		ctx.Bind(&todo)
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
	})
	e.POST("/delete/todo", func(ctx echo.Context) error {
		todo := new(model.Todo)
		ctx.Bind(&todo)
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
	})

	e.Logger.Fatal(e.Start(":8080"))
}
