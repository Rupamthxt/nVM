package main

import "fmt"

func main() {
	code := []byte{
		byte(JUMPDEST), // 0: Marker
		0x60, 0x00,     // 1: PUSH1 0 (Target)
		byte(JUMP), // 3: JUMP (Go to 0)
	}

	fmt.Println("Starting EVM...")
	vm := NewEVM(code, 20)
	vm.Run()

	if vm.stack.Peek() != nil {
		fmt.Printf("Final Stack top: 0x%x (Expected 0xAA)\n", vm.stack.Pop())
	}

	fmt.Printf("Final remaining gas: %d\n", vm.gas)
}
