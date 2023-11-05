package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"

	"github.com/leon123858/go-aid/controller"
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

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	getGroup := e.Group("/get")
	getGroup.GET("/utxo", controller.GenerateChainGetController(chain, "getUnspent"))
	getGroup.GET("/balance", controller.GenerateChainGetController(chain, "getBalance"))
	getGroup.GET("/privatekey", controller.GenerateChainGetController(chain, "getPrivateKey"))
	getGroup.GET("/transaction", controller.GenerateChainGetController(chain, "getTransaction"))
	getGroup.POST("/contractmessage", controller.GenerateChainPostController(chain, "dumpContractMessage"))

	blockGroup := e.Group("/block")
	blockGroup.POST("/generate", controller.GenerateChainPostController(chain, "generateBlock"))

	rawTransactionGroup := e.Group("/rawtransaction")
	rawTransactionGroup.POST("/create", controller.GenerateChainPostController(chain, "createRawTransaction"))
	rawTransactionGroup.POST("/send", controller.GenerateChainPostController(chain, "sendRawTransaction"))
	rawTransactionGroup.POST("/sign", controller.GenerateChainPostController(chain, "signRawTransaction"))

	e.Logger.Fatal(e.Start(":8080"))
}
