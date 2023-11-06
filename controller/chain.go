package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/leon123858/go-aid/modal"
	"github.com/leon123858/go-aid/service/rpc"
	"net/http"
)

func customResponseHandler(ctx echo.Context, err error, data interface{}) error {
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"result": "fail",
			"error":  err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"result": "success",
		"data":   data,
	})
}

func GenerateChainGetController(chain *our_chain_rpc.Bitcoind, which string) echo.HandlerFunc {
	switch which {
	case "getUnspent":
		return func(ctx echo.Context) error {
			list, err := chain.ListUnspent()
			return customResponseHandler(ctx, err, list)
		}
	case "getBalance":
		return func(ctx echo.Context) error {
			address := ctx.QueryParam("address")
			balance, err := chain.GetBalance(address, 1)
			return customResponseHandler(ctx, err, balance)
		}
	case "getPrivateKey":
		return func(ctx echo.Context) error {
			address := ctx.QueryParam("address")
			privateKey, err := chain.DumpPrivKey(address)
			return customResponseHandler(ctx, err, privateKey)
		}
	case "getTransaction":
		return func(ctx echo.Context) error {
			txid := ctx.QueryParam("txid")
			tx, err := chain.GetTransaction(txid)
			return customResponseHandler(ctx, err, tx)
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
			return customResponseHandler(ctx, err, blockIds)
		}
	case "dumpContractMessage":
		return func(ctx echo.Context) error {
			req := new(model.ContractRequest)
			err := ctx.Bind(&req)
			if err != nil {
				return err
			}
			message, err := chain.DumpContractMessage(req.Address, req.Arguments)
			return customResponseHandler(ctx, err, message)
		}
	case "createRawTransaction":
		return func(ctx echo.Context) error {
			req := new(model.RawTransactionRequest)
			err := ctx.Bind(&req)
			if err != nil {
				return err
			}
			result, err := chain.CreateRawTransaction(req.Inputs, req.Outputs, req.Contract)
			return customResponseHandler(ctx, err, result)
		}
	case "signRawTransaction":
		return func(ctx echo.Context) error {
			req := new(model.SignRequest)
			err := ctx.Bind(&req)
			if err != nil {
				return err
			}
			result, err := chain.SignRawTransaction(req.RawTransaction, req.PrivateKey)
			return customResponseHandler(ctx, err, result)
		}
	case "sendRawTransaction":
		return func(ctx echo.Context) error {
			req := new(model.SendRequest)
			err := ctx.Bind(&req)
			if err != nil {
				return err
			}
			result, err := chain.SendRawTransaction(req.RawTransaction)
			return customResponseHandler(ctx, err, result)
		}
	default:
		return nil
	}
}
