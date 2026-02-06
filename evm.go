package main

import (
	"fmt"
	"math/big"
)

type EVM struct {
	pc     int    // Program Counter
	stack  *Stack // The Stack
	memory []byte // Volatile Memory (RAM)
	code   []byte // The Bytecode
}

func NewEVM(code []byte) *EVM {
	return &EVM{
		pc:     0,
		stack:  NewStack(),
		memory: make([]byte, 0),
		code:   code,
	}
}

// Run executes the EVM bytecode. It fetches, decodes, and executes instructions until it encounters a STOP opcode or reaches the end of the code.
func (vm *EVM) Run() {
	for vm.pc < len(vm.code) {
		// Fetch
		op := OpCode(vm.code[vm.pc])
		vm.pc++ // Advance PC

		// Decode & Execute
		fmt.Printf("PC: %d | OpCode: 0x%x | Stack Top: %v\n", vm.pc-1, op, vm.stack.Peek())

		switch op {
		case STOP:
			return
		case PUSH1:
			// Read next byte as a value and push it onto the stack
			val := big.NewInt(int64(vm.code[vm.pc]))
			vm.pc++
			vm.stack.Push(val)
		case ADD:
			// Pop two values from the stack, add them, and push the result back onto the stack
			a := vm.stack.Pop()
			b := vm.stack.Pop()
			// a + b
			result := new(big.Int).Add(a, b)
			vm.stack.Push(result)
		case MUL:
			a := vm.stack.Pop()
			b := vm.stack.Pop()
			result := new(big.Int).Mul(a, b)
			vm.stack.Push(result)
		case SUB:
			a := vm.stack.Pop()
			b := vm.stack.Pop()
			result := new(big.Int).Sub(b, a) // Note: SUB is b - a
			vm.stack.Push(result)
		case MSTORE:
			offset := vm.stack.Pop().Uint64()
			value := vm.stack.Pop()
			// Exapnd memory if needed (EVM words are 32 bytes)
			vm.ensureMemory(offset, 32)

			valBytes := value.FillBytes(make([]byte, 32))

			copy(vm.memory[offset:], valBytes)
			fmt.Printf("MSTORE: Wrote %v to address %d\n", value, offset)

		case MLOAD:
			offset := vm.stack.Pop().Uint64()
			vm.ensureMemory(offset, 32)

			// Read 32 bytes from memory and push it onto the stack
			valBytes := vm.memory[offset : offset+32]
			val := new(big.Int).SetBytes(valBytes)
			vm.stack.Push(val)
			fmt.Printf("MLOAD: Read %v from address %d\n", val, offset)
		default:
			panic(fmt.Sprintf("Unknown opcode: 0x%x", op))
		}
	}
}

// ensureMemory checks if the memory is large enough to accommodate the specified offset and size. If not, it resizes the memory accordingly.
func (vm *EVM) ensureMemory(offset, size uint64) {
	requiredSize := offset + size
	if uint64(len(vm.memory)) < requiredSize {
		// Create a new slice with the requirted size
		newMem := make([]byte, requiredSize)
		// Copy old data over
		copy(newMem, vm.memory)
		vm.memory = newMem
	}
}
