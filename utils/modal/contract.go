package model

type ContractRequest struct {
	Address   string   `json:"address"`
	Arguments []string `json:"arguments"`
}
