package main

import (
	"fmt"
	"math/big"
)

// OpFunc defines the type for EVM instruction functions.
// Each function takes a pointer to the EVM instance and performs its specific operation.
type OpFunc func(vm *EVM)

// JumpTable is a mapping of opcode values (0-255) to their corresponding instruction functions.
// This allows for efficient dispatching of instructions during execution.
type JumpTable [256]OpFunc

func NewJumpTable() *JumpTable {
	table := &JumpTable{}

	// Fill the jump table with default functions that panic for unknown opcodes.
	for i := range 256 {
		table[i] = func(vm *EVM) { panic(fmt.Sprintf("Unknown Opcode: 0x%x", i)) }
	}

	// Define specific functions for known opcodes.
	table[0x00] = opStop
	table[0x01] = opAdd
	table[0x02] = opMul
	table[0x03] = opSub
	table[0x60] = opPush1
	table[0x51] = opMload
	table[0x52] = opMstore
	table[0x10] = opLt
	table[0x56] = opJump
	table[0x57] = opJumpI
	table[0x5B] = opJumpDest

	return table
}

var GasTable [256]uint64

func init() {
	for i := range 256 {
		GasTable[i] = 1
	}

	GasTable[ADD] = GasQuickStep
	GasTable[SUB] = GasQuickStep
	GasTable[PUSH1] = GasQuickStep
	GasTable[MUL] = GasSlowStep
	GasTable[MSTORE] = GasMemoryStep
	GasTable[MLOAD] = GasMemoryStep
	GasTable[JUMP] = GasQuickStep
	GasTable[SSTORE] = GasStorageStep
	GasTable[STOP] = 0 // Free to stop
}

func opStop(vm *EVM) {}

func opAdd(vm *EVM) {
	a := vm.stack.Pop()
	b := vm.stack.Pop()
	vm.stack.Push(new(big.Int).Add(a, b))
}

func opMul(vm *EVM) {
	a := vm.stack.Pop()
	b := vm.stack.Pop()
	vm.stack.Push(new(big.Int).Mul(a, b))
}

func opSub(vm *EVM) {
	a := vm.stack.Pop()
	b := vm.stack.Pop()
	vm.stack.Push(new(big.Int).Sub(b, a)) // Note: SUB is b - a
}

func opPush1(vm *EVM) {
	// Read next byte as a value and push it onto the stack
	val := big.NewInt(int64(vm.code[vm.pc]))
	vm.pc++
	vm.stack.Push(val)
}

func opMstore(vm *EVM) {
	offset := vm.stack.Pop().Uint64()
	value := vm.stack.Pop()
	// Exapnd memory if needed (EVM words are 32 bytes)
	vm.ensureMemory(offset, 32)
	valBytes := value.FillBytes(make([]byte, 32))
	copy(vm.memory[offset:], valBytes)
}

func opMload(vm *EVM) {
	offset := vm.stack.Pop().Uint64()
	vm.ensureMemory(offset, 32)
	// Read 32 bytes from memory and push it onto the stack
	valBytes := vm.memory[offset : offset+32]
	val := new(big.Int).SetBytes(valBytes)
	vm.stack.Push(val)
}

func opLt(vm *EVM) {
	a := vm.stack.Pop()
	b := vm.stack.Pop()
	// If a < b, push 1 (true), else push 0 (false)
	if a.Cmp(b) < 0 {
		vm.stack.Push(big.NewInt(1)) // true
	} else {
		vm.stack.Push(big.NewInt(0)) // false
	}
}

func opJumpI(vm *EVM) {
	dest := vm.stack.Pop().Uint64()
	condition := vm.stack.Pop() // 0 = False, Anything else = True

	if condition.Sign() != 0 {
		// Condition true, perform jump
		if dest >= uint64(len(vm.code)) || vm.code[dest] != byte(JUMPDEST) {
			panic(fmt.Sprintf("invalid jump destination at %d", dest))
		}
		vm.pc = int(dest)
	} else {
		vm.pc++ // Condition false, just keep waiting
	}
}

func opJump(vm *EVM) {
	dest := vm.stack.Pop().Uint64()

	if dest >= uint64(len(vm.code)) || vm.code[dest] != byte(JUMPDEST) {
		panic(fmt.Sprintf("Invalid jump destination at %d", dest))
	}

	vm.pc = int(dest)
}
func opJumpDest(vm *EVM) {
	// Just a marker for valid jump destinations, no action needed
	fmt.Printf("JUMPDEST: Valid jump destination at %d\n", vm.pc-1)
}
