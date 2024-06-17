package core

import (
	"errors"
)

type VM struct {
	Memory    []byte
	Registers map[string]int
	GasLimit  int
	GasUsed   int
	Contracts map[string]string
}

func NewVM() *VM {
	return &VM{
		Memory:    make([]byte, 1024),
		Registers: make(map[string]int),
		GasLimit:  1000000,
		GasUsed:   0,
		Contracts: make(map[string]string),
	}
}
func (vm *VM) LoadContract(name, code string) error {
	if name == "" || code == "" {
		return errors.New("invalid contract name or code")
	}
	vm.Contracts[name] = code
	return nil
}

func (vm *VM) Execute(contractName string) (string, error) {
	contract, exists := vm.Contracts[contractName]
	if !exists {
		return "", errors.New("contract not found")
	}

	// Hier könnte die Logik zur Ausführung des Smart Contracts implementiert werden
	result, err := vm.runContract(contract)
	if err != nil {
		return "", err
	}
	return result, nil
}
