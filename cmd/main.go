package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	ourChain "github.com/leon123858/go-aid/service/rpc"
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

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORS())

	getGroup := e.Group("/get")
	getGroup.GET("/balance", controller.GenerateChainGetController(chain, "getBalance"))       // just used for node owner
	getGroup.GET("/privatekey", controller.GenerateChainGetController(chain, "getPrivateKey")) // just used for node owner
	getGroup.GET("/transaction", controller.GenerateChainGetController(chain, "getTransaction"))
	getGroup.GET("/utxo", controller.GenerateChainGetController(chain, "getUnspent"))

	getGroup.POST("/contractmessage", controller.GenerateChainPostController(chain, "dumpContractMessage"))

	blockGroup := e.Group("/block")
	blockGroup.POST("/generate", controller.GenerateChainPostController(chain, "generateBlock")) // just used for test

	rawTransactionGroup := e.Group("/rawtransaction")
	rawTransactionGroup.POST("/create", controller.GenerateChainPostController(chain, "createRawTransaction"))
	rawTransactionGroup.POST("/send", controller.GenerateChainPostController(chain, "sendRawTransaction"))
	rawTransactionGroup.POST("/sign", controller.GenerateChainPostController(chain, "signRawTransaction"))

	e.Logger.Fatal(e.Start(":8080"))
}
