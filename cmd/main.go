package main

import (
	"fmt"
	"vm/core"
)

func main() {
	vm := core.NewVM()
	err := vm.LoadContract("example", string([]byte{byte(core.OP_PRINT), byte(core.OP_ADD)}))
	if err != nil {
		fmt.Printf("Failed to load contract: %v\n", err)
		return
	}

	vm.Registers["a"] = 5
	vm.Registers["b"] = 10

	result, err := vm.Execute("example")
	if err != nil {
		fmt.Printf("Failed to execute contract: %v\n", err)
		return
	}

	fmt.Printf("Execution result: %s\n", result)
}
