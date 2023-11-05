package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	model "github.com/leon123858/go-aid/utils/modal"
	util "github.com/leon123858/go-aid/utils/mongo"
	"github.com/leon123858/go-aid/utils/repository"
	ourChain "github.com/leon123858/go-aid/utils/rpc"
)

const (
	SERVER_HOST = "127.0.0.1"
	SERVER_PORT = 8332
	USER        = "test"
	PASSWD      = "test"
	USESSL      = false
	//WALLET_PASSPHRASE = "WalletPassphrase"
)

func main() {
	chain, err := ourChain.New(SERVER_HOST, SERVER_PORT, USER, PASSWD, USESSL)
	if err != nil {
		log.Fatal(err)
	}
	db := util.GetMgoCli().Database("todo")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.POST("/create/todo", func(ctx echo.Context) error {
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
	})
	e.GET("/get/todo", func(ctx echo.Context) error {
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
	})
	e.POST("/update/todo", func(ctx echo.Context) error {
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
	})
	e.POST("/delete/todo", func(ctx echo.Context) error {
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
	})

	e.GET("/get/utxo", func(ctx echo.Context) error {
		list, err := chain.ListUnspent()
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"result": "fail",
				"error":  err.Error(),
			})
		}
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"result": "success",
			"data":   list,
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
