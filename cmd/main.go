package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	ourChain "github.com/leon123858/go-aid/service/rpc"
	"github.com/leon123858/go-aid/service/sqlite"
	"log"

	"github.com/leon123858/go-aid/controller"
	"github.com/leon123858/go-aid/utils"
)

func main() {
	utils.LoadConfig("./config.toml")
	chain, err := ourChain.New(
		utils.OurChainConfigInstance.ServerHost,
		utils.OurChainConfigInstance.ServerPort,
		utils.OurChainConfigInstance.User,
		utils.OurChainConfigInstance.Passwd,
		utils.OurChainConfigInstance.UseSsl)
	if err != nil {
		log.Fatal(err)
	}
	sqliteClient := sqlite.Client{}
	if sqlite.New(&sqliteClient) != nil {
		log.Fatal("sqlite init failed")
	}
	repositoryDTO := controller.RepositoryDTO{Chain: chain, Database: &sqliteClient}

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORS())

	getGroup := e.Group("/get")
	getGroup.GET("/balance", controller.GenerateChainGetController(repositoryDTO, "getBalance"))       // just used for node owner
	getGroup.GET("/privatekey", controller.GenerateChainGetController(repositoryDTO, "getPrivateKey")) // just used for node owner
	getGroup.GET("/transaction", controller.GenerateChainGetController(repositoryDTO, "getTransaction"))
	getGroup.GET("/utxo", controller.GenerateChainGetController(repositoryDTO, "getUnspent"))

	getGroup.POST("/contractmessage", controller.GenerateChainPostController(repositoryDTO, "dumpContractMessage"))

	blockGroup := e.Group("/block")
	blockGroup.POST("/generate", controller.GenerateChainPostController(repositoryDTO, "generateBlock")) // just used for test

	rawTransactionGroup := e.Group("/rawtransaction")
	rawTransactionGroup.POST("/create", controller.GenerateChainPostController(repositoryDTO, "createRawTransaction"))
	rawTransactionGroup.POST("/send", controller.GenerateChainPostController(repositoryDTO, "sendRawTransaction"))
	rawTransactionGroup.POST("/sign", controller.GenerateChainPostController(repositoryDTO, "signRawTransaction"))

	e.Logger.Fatal(e.Start(":8080"))
}
