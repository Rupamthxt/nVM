package main

import (
	"fmt"
)

type EVM struct {
	pc           int        // Program Counter
	stack        *Stack     // The Stack
	memory       []byte     // Volatile Memory (RAM)
	code         []byte     // The Bytecode
	instructions *JumpTable // Jump Table for opcode handlers
}

func NewEVM(code []byte) *EVM {
	return &EVM{
		pc:           0,
		stack:        NewStack(),
		memory:       make([]byte, 0),
		code:         code,
		instructions: NewJumpTable(),
	}
}

// Run executes the EVM bytecode. It fetches, decodes, and executes instructions until it encounters a STOP opcode or reaches the end of the code.
func (vm *EVM) Run() {
	for vm.pc < len(vm.code) {
		// Fetch
		op := OpCode(vm.code[vm.pc])
		vm.pc++ // Move past the opcode byte

		// Decode & Execute
		fmt.Printf("PC: %d | OpCode: 0x%x | Stack Top: %v\n", vm.pc-1, op, vm.stack.Peek())

		operation := vm.instructions[op]

		operation(vm) // Execute the operation

		if op != JUMP && op != JUMPI && op != PUSH1 && op != JUMPDEST {
			vm.pc++ // Move to the next instruction (except for JUMP/JUMPI which manage PC themselves, and PUSH1 which advances PC by 2)
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
