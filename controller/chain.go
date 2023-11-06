package controller

import (
	"github.com/labstack/echo"
	"github.com/leon123858/go-aid/modal"
	"github.com/leon123858/go-aid/repository/rpc"
	"net/http"
)

func GenerateChainGetController(chain *our_chain_rpc.Bitcoind, which string) echo.HandlerFunc {
	switch which {
	case "getUnspent":
		return func(ctx echo.Context) error {
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
		}
	case "getBalance":
		return func(ctx echo.Context) error {
			address := ctx.QueryParam("address")
			balance, err := chain.GetBalance(address, 1)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
					"result": "fail",
					"error":  err.Error(),
				})
			}
			return ctx.JSON(http.StatusOK, map[string]interface{}{
				"result": "success",
				"data":   balance,
			})
		}
	case "getPrivateKey":
		return func(ctx echo.Context) error {
			address := ctx.QueryParam("address")
			privateKey, err := chain.DumpPrivKey(address)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
					"result": "fail",
					"error":  err.Error(),
				})
			}
			return ctx.JSON(http.StatusOK, map[string]interface{}{
				"result": "success",
				"data":   privateKey,
			})
		}
	case "getTransaction":
		return func(ctx echo.Context) error {
			txid := ctx.QueryParam("txid")
			tx, err := chain.GetTransaction(txid)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
					"result": "fail",
					"error":  err.Error(),
				})
			}
			return ctx.JSON(http.StatusOK, map[string]interface{}{
				"result": "success",
				"data":   tx,
			})
		}
	default:
		return nil
	}
}

func GenerateChainPostController(chain *our_chain_rpc.Bitcoind, which string) echo.HandlerFunc {
	switch which {
	case "generateBlock":
		return func(ctx echo.Context) error {
			blockIds, err := chain.GenerateBlock(1)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
					"result": "fail",
					"error":  err.Error(),
				})
			}
			return ctx.JSON(http.StatusOK, map[string]interface{}{
				"result": "success",
				"data":   blockIds,
			})
		}
	case "dumpContractMessage":
		return func(ctx echo.Context) error {
			req := new(model.ContractRequest)
			err := ctx.Bind(&req)
			if err != nil {
				return err
			}
			message, err := chain.DumpContractMessage(req.Address, req.Arguments)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
					"result": "fail",
					"error":  err.Error(),
				})
			}
			return ctx.JSON(http.StatusOK, map[string]interface{}{
				"result": "success",
				"data":   message,
			})
		}
	case "createRawTransaction":
		return func(ctx echo.Context) error {
			req := new(model.RawTransactionRequest)
			err := ctx.Bind(&req)
			if err != nil {
				return err
			}
			result, err := chain.CreateRawTransaction(req.Inputs, req.Outputs, req.Contract)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
					"result": "fail",
					"error":  err.Error(),
				})
			}
			return ctx.JSON(http.StatusOK, map[string]interface{}{
				"result": "success",
				"data":   result,
			})
		}
	case "signRawTransaction":
		return func(ctx echo.Context) error {
			req := new(model.SignRequest)
			err := ctx.Bind(&req)
			if err != nil {
				return err
			}
			result, err := chain.SignRawTransaction(req.RawTransaction, req.PrivateKey)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
					"result": "fail",
					"error":  err.Error(),
				})
			}
			return ctx.JSON(http.StatusOK, map[string]interface{}{
				"result": "success",
				"data":   result,
			})
		}
	case "sendRawTransaction":
		return func(ctx echo.Context) error {
			req := new(model.SendRequest)
			err := ctx.Bind(&req)
			if err != nil {
				return err
			}
			result, err := chain.SendRawTransaction(req.RawTransaction)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
					"result": "fail",
					"error":  err.Error(),
				})
			}
			return ctx.JSON(http.StatusOK, map[string]interface{}{
				"result": "success",
				"data":   result,
			})
		}
	default:
		return nil
	}
}
