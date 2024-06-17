package core

import (
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
)

// VM represents the virtual machine
type VM struct {
	Memory    []byte            // Memory for the VM
	Registers map[string]int    // Registers for the VM
	GasLimit  int               // Gas limit for execution
	GasUsed   int               // Gas used during execution
	Contracts map[string]string // Stored smart contracts
}

// NewVM creates a new VM instance with default parameters
func NewVM() *VM {
	return &VM{
		Memory:    make([]byte, 1024),             // Default memory size
		Registers: map[string]int{"a": 0, "b": 0}, // Initialize default registers
		GasLimit:  1000000,                        // Default gas limit
		GasUsed:   0,
		Contracts: make(map[string]string),
	}
}

// LoadContract loads a smart contract into the VM
func (vm *VM) LoadContract(name, code string) error {
	if name == "" || code == "" {
		return errors.New("invalid contract name or code")
	}
	vm.Contracts[name] = code
	return nil
}

// Execute executes a smart contract
func (vm *VM) Execute(contractName string) (string, error) {
	contract, exists := vm.Contracts[contractName]
	if !exists {
		return "", errors.New("contract not found")
	}

	// Here you can implement the logic to execute the smart contract
	result, err := vm.runContract(contract)
	if err != nil {
		return "", err
	}
	return result, nil
}

// Memory operations

// ReadMemory reads a value from memory at the given address
func (vm *VM) ReadMemory(address int) (byte, error) {
	if address < 0 || address >= len(vm.Memory) {
		return 0, errors.New("memory access out of bounds")
	}
	return vm.Memory[address], nil
}

// WriteMemory writes a value to memory at the given address
func (vm *VM) WriteMemory(address int, value byte) error {
	if address < 0 || address >= len(vm.Memory) {
		return errors.New("memory access out of bounds")
	}
	vm.Memory[address] = value
	return nil
}

// checkGasLimit checks if there is enough gas available
func (vm *VM) checkGasLimit(cost int) error {
	if vm.GasUsed+cost > vm.GasLimit {
		return errors.New("out of gas")
	}
	vm.GasUsed += cost
	return nil
}

// runContract executes the smart contract code
func (vm *VM) runContract(contract string) (string, error) {
	// Simple contract code processing
	for i := 0; i < len(contract); i++ {
		opcode := Opcode(contract[i])
		switch opcode {
		case OP_PRINT: // Example opcode: Print
			fmt.Println("Print opcode executed")
			if err := vm.checkGasLimit(1); err != nil {
				return "", err
			}
		case OP_ADD:
			// Safety check: Verify registers exist
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
		default:
			return "", fmt.Errorf("unknown opcode: %d", opcode)
		}
	}

	return "execution successful", nil
}
