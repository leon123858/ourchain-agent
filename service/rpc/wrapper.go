package our_chain_rpc

import (
	"encoding/json"
	"errors"
)

func (b *Bitcoind) GetContractGeneralInterface(address string) (generalInterface ContractGeneralInterface, err error) {
	result, err := b.DumpContractMessage(address, []string{"get"})
	if err != nil {
		return
	}
	if result == "" || result == "null" || result == "undefined" {
		err = errors.New("contract get error")
		return
	}
	err = json.Unmarshal([]byte(result), &generalInterface)
	if generalInterface.Protocol == "" || generalInterface.Version == "" {
		err = errors.New("contract serialization error")
	}
	return
}
