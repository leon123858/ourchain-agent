package model

import (
	chain "github.com/leon123858/go-aid/service/rpc"
)

type ContractRequest struct {
	Address   string   `json:"address"`
	Arguments []string `json:"arguments"`
}

type RawTransactionRequest struct {
	Inputs   []chain.TxInput       `json:"inputs"`
	Outputs  []chain.TxOutput      `json:"outputs"`
	Contract chain.ContractMessage `json:"contract"`
}

type SignRequest struct {
	RawTransaction string `json:"rawTransaction"`
	PrivateKey     string `json:"privateKey"`
}

type SendRequest struct {
	RawTransaction string `json:"rawTransaction"`
}
