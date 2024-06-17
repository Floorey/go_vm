package tests

import (
	"testing"
	"vm/core"
)

func TestIntegration(t *testing.T) {
	vm := core.NewVM()

	// Load a contract that includes multiple opcodes
	contract := string([]byte{
		byte(core.OP_PRINT),
		byte(core.OP_ADD),
		byte(core.OP_SUB),
		byte(core.OP_MUL),
		byte(core.OP_DIV),
	})
	err := vm.LoadContract("integration_contract", contract)
	if err != nil {
		t.Fatalf("Failed to load contract: %v", err)
	}

	// Initialize registers for the test
	vm.Registers["a"] = 20
	vm.Registers["b"] = 4

	// Execute the contract
	result, err := vm.Execute("integration_contract")
	if err != nil {
		t.Fatalf("Failed to execute contract: %v", err)
	}

	// Verify the final state of the registers and memory
	if result != "execution successful" {
		t.Fatalf("Unexpected result: %s", result)
	}

	expected := 20 + 4 - 4*4/4 // Correct the expected value sequence
	if vm.Registers["a"] != expected {
		t.Fatalf("Unexpected register value for a: %d, expected: %d", vm.Registers["a"], expected)
	}
}
