package main

import "fmt"

func main() {
	code := []byte{
		0x60, 10,
		0x60, 20,
		byte(ADD),
		0x60, 0x00,
		byte(MSTORE),
		0x60, 0x00,
		byte(MLOAD),
	}

	fmt.Println("Starting EVM...")
	vm := NewEVM(code)
	vm.Run()

	result := vm.stack.Pop()
	fmt.Printf("Result on stack: %s\n", result.String())
	fmt.Printf("Final memory state :%v\n", vm.memory)
	fmt.Println("EVM execution completed.")
}
