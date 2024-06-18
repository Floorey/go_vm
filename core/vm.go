package core

import (
	"bytes"
	"errors"
	"fmt"
)

type Opcode byte

const (
	OP_PRINT Opcode = iota
	OP_ADD
	OP_SUB
	OP_MUL
	OP_DIV
	OP_TRANSFER
)

type VM struct {
	Memory         []byte            // Memory for the VM
	Registers      map[string]int    // Registers for the VM
	GasLimit       int               // Gas limit for execution
	GasUsed        int               // Gas used during execution
	Contracts      map[string]string // Stored smart contracts
	AccountManager *AccountManager   // Manage accounts and balances
}

func NewVM() *VM {
	return &VM{
		Memory:         make([]byte, 1024),                          // Default memory size
		Registers:      map[string]int{"a": 0, "b": 0, "amount": 0}, // Initialize default registers
		GasLimit:       1000000,                                     // Default gas limit
		GasUsed:        0,
		Contracts:      make(map[string]string),
		AccountManager: NewAccountManager(), // Initialize account manager
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

	result, err := vm.runContract(contract)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (vm *VM) runContract(contract string) (string, error) {
	for i := 0; i < len(contract); i++ {
		opcode := Opcode(contract[i])
		switch opcode {
		case OP_PRINT:
			fmt.Println("Print opcode executed")
			if err := vm.checkGasLimit(1); err != nil {
				return "", err
			}
		case OP_ADD:
			if _, ok := vm.Registers["a"]; !ok {
				return "", errors.New("register a not found")
			}
			if _, ok := vm.Registers["b"]; !ok {
				return "", errors.New("register b not found")
			}
			vm.Registers["a"] = vm.Registers["a"] + vm.Registers["b"]
			fmt.Printf("ADD executed: a = %d\n", vm.Registers["a"])
			if err := vm.checkGasLimit(2); err != nil {
				return "", err
			}
		case OP_SUB:
			if _, ok := vm.Registers["a"]; !ok {
				return "", errors.New("register a not found")
			}
			if _, ok := vm.Registers["b"]; !ok {
				return "", errors.New("register b not found")
			}
			vm.Registers["a"] = vm.Registers["a"] - vm.Registers["b"]
			fmt.Printf("SUB executed: a = %d\n", vm.Registers["a"])
			if err := vm.checkGasLimit(2); err != nil {
				return "", err
			}
		case OP_MUL:
			if _, ok := vm.Registers["a"]; !ok {
				return "", errors.New("register a not found")
			}
			if _, ok := vm.Registers["b"]; !ok {
				return "", errors.New("register b not found")
			}
			vm.Registers["a"] = vm.Registers["a"] * vm.Registers["b"]
			fmt.Printf("MUL executed: a = %d\n", vm.Registers["a"])
			if err := vm.checkGasLimit(3); err != nil {
				return "", err
			}
		case OP_DIV:
			if _, ok := vm.Registers["a"]; !ok {
				return "", errors.New("register a not found")
			}
			if _, ok := vm.Registers["b"]; !ok {
				return "", errors.New("register b not found")
			}
			if vm.Registers["b"] == 0 {
				return "", errors.New("division by zero")
			}
			vm.Registers["a"] = vm.Registers["a"] / vm.Registers["b"]
			fmt.Printf("DIV executed: a = %d\n", vm.Registers["a"])
			if err := vm.checkGasLimit(3); err != nil {
				return "", err
			}
		case OP_TRANSFER:
			result, err := vm.executeTransaction()
			if err != nil {
				return "", err
			}
			fmt.Println(result)
			if err := vm.checkGasLimit(10); err != nil {
				return "", err
			}
		default:
			return "", fmt.Errorf("unknown opcode: %d", opcode)
		}
	}

	return "execution successful", nil
}

func (vm *VM) executeTransaction() (string, error) {
	from := string(bytes.Trim(vm.Memory[0:32], "\x00")) // Trim null bytes
	to := string(bytes.Trim(vm.Memory[32:64], "\x00"))  // Trim null bytes
	amount := vm.Registers["amount"]
	err := vm.AccountManager.Transfer(from, to, amount)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("TRANSFER executed: %d from %s to %s", amount, from, to), nil
}

func (vm *VM) checkGasLimit(cost int) error {
	if vm.GasUsed+cost > vm.GasLimit {
		return errors.New("out of gas")
	}
	vm.GasUsed += cost
	return nil
}
