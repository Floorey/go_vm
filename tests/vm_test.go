package tests

import (
	"testing"
	"vm/core"
)

func TestVM(t *testing.T) {
	vm := core.NewVM()

	// Test loading a contract with valid opcodes
	contract := string([]byte{byte(core.OP_PRINT), byte(core.OP_ADD)})
	err := vm.LoadContract("test_contract", contract)
	if err != nil {
		t.Fatalf("Failed to load contract: %v", err)
	}

	// Test executing a contract
	vm.Registers["a"] = 5
	vm.Registers["b"] = 10
	result, err := vm.Execute("test_contract")
	if err != nil {
		t.Fatalf("Failed to execute contract: %v", err)
	}

	if result != "execution successful" {
		t.Fatalf("Unexpected result: %s", result)
	}
	if vm.Registers["a"] != 15 {
		t.Fatalf("Unexpected register value: %d", vm.Registers["a"])
	}
}

func TestOpcodes(t *testing.T) {
	vm := core.NewVM()

	// Test opcode processing
	contract := string([]byte{byte(core.OP_PRINT), byte(core.OP_ADD)})
	err := vm.LoadContract("op_test", contract)
	if err != nil {
		t.Fatalf("Failed to load contract: %v", err)
	}

	vm.Registers["a"] = 5
	vm.Registers["b"] = 10

	result, err := vm.Execute("op_test")
	if err != nil {
		t.Fatalf("Failed to execute contract: %v", err)
	}

	if result != "execution successful" {
		t.Fatalf("Unexpected result: %s", result)
	}
	if vm.Registers["a"] != 15 {
		t.Fatalf("Unexpected register value: %d", vm.Registers["a"])
	}
}

func TestGasLimit(t *testing.T) {
	vm := core.NewVM()
	vm.GasLimit = 1 // Set a low gas limit

	// Test executing a contract that exceeds the gas limit
	contract := string([]byte{byte(core.OP_PRINT), byte(core.OP_PRINT)})
	err := vm.LoadContract("gas_test", contract)
	if err != nil {
		t.Fatalf("Failed to load contract: %v", err)
	}

	_, err = vm.Execute("gas_test")
	if err == nil {
		t.Fatalf("Expected out of gas error, got none")
	}
}
