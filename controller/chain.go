package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/leon123858/go-aid/service/scanner"
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

func GenerateChainGetController(dto RepositoryDTO, which string) echo.HandlerFunc {
	switch which {
	case "getUnspent":
		return func(ctx echo.Context) error {
			address := ctx.QueryParam("address")
			target := []string{address}
			if address == "" {
				target = []string{}
			}
			list, err := scanner.ListUnspent(dto.Chain, dto.Database, target, 2)
			return customResponseHandler(ctx, err, list)
		}
	case "getBalance":
		return func(ctx echo.Context) error {
			address := ctx.QueryParam("address")
			balance, err := dto.Chain.GetBalance(address, 1)
			return customResponseHandler(ctx, err, balance)
		}
	case "getPrivateKey":
		return func(ctx echo.Context) error {
			address := ctx.QueryParam("address")
			privateKey, err := dto.Chain.DumpPrivKey(address)
			return customResponseHandler(ctx, err, privateKey)
		}
	case "getTransaction":
		return func(ctx echo.Context) error {
			txid := ctx.QueryParam("txid")
			tx, err := dto.Chain.GetRawTransaction(txid)
			return customResponseHandler(ctx, err, tx)
		}
	default:
		return nil
	}
}

func GenerateChainPostController(dto RepositoryDTO, which string) echo.HandlerFunc {
	switch which {
	case "generateBlock":
		return func(ctx echo.Context) error {
			blockIds, err := dto.Chain.GenerateBlock(1)
			return customResponseHandler(ctx, err, blockIds)
		}
	case "dumpContractMessage":
		return func(ctx echo.Context) error {
			req := new(ContractRequest)
			err := ctx.Bind(&req)
			if err != nil {
				return err
			}
			message, err := dto.Chain.DumpContractMessage(req.Address, req.Arguments)
			return customResponseHandler(ctx, err, message)
		}
	case "createRawTransaction":
		return func(ctx echo.Context) error {
			req := new(RawTransactionRequest)
			err := ctx.Bind(&req)
			if err != nil {
				return err
			}
			result, err := dto.Chain.CreateRawTransaction(req.Inputs, req.Outputs, req.Contract)
			return customResponseHandler(ctx, err, result)
		}
	case "signRawTransaction":
		return func(ctx echo.Context) error {
			req := new(SignRequest)
			err := ctx.Bind(&req)
			if err != nil {
				return err
			}
			result, err := dto.Chain.SignRawTransaction(req.RawTransaction, req.PrivateKey)
			return customResponseHandler(ctx, err, result)
		}
	case "sendRawTransaction":
		return func(ctx echo.Context) error {
			req := new(SendRequest)
			err := ctx.Bind(&req)
			if err != nil {
				return err
			}
			result, err := dto.Chain.SendRawTransaction(req.RawTransaction)
			return customResponseHandler(ctx, err, result)
		}
	default:
		return nil
	}
}
