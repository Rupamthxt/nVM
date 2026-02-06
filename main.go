package main

import "fmt"

func main() {
	code := []byte{
		0x60, 0x01,
		0x60, 0x08,
		byte(JUMPI),
		0x60, 0xFF,
		byte(STOP),
		byte(JUMPDEST),
		0x60, 0xAA,
	}

	fmt.Println("Starting EVM...")
	vm := NewEVM(code)
	vm.Run()

	if vm.stack.Peek() != nil {
		fmt.Printf("Final Stack top: 0x%x (Expected 0xAA)\n", vm.stack.Pop())
	}
}
