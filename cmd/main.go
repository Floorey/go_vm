package main

import (
	"fmt"
)

func main() {
	fmt.Println("Blockchain VM starting...")

	vm := NewVM()

	result, err := vm.Execute("sample_contract")
	if err != nil {
		fmt.Println("Error executing contract:", err)
		return
	}
	fmt.Println("Execution result:", result)
}
