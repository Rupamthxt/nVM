package main

type OpCode byte

const (
	STOP OpCode = 0x00
	ADD  OpCode = 0x01
	MUL  OpCode = 0x02
	SUB  OpCode = 0x03

	// Push Operations
	PUSH1 OpCode = 0x60
	PUSH2 OpCode = 0x61

	// Memory Operations
	MLOAD  OpCode = 0x51
	MSTORE OpCode = 0x52

	LT  OpCode = 0x10 // Less Than
	GT  OpCode = 0x11 // Greater Than
	EQ  OpCode = 0x14 // Equal To
	AND OpCode = 0x16 // Bitwise AND
	OR  OpCode = 0x17 // Bitwise OR
	XOR OpCode = 0x18 // Bitwise XOR

	JUMPDEST OpCode = 0x5B // Mark a valid jump target
	JUMP     OpCode = 0x56 // Jump to address
	JUMPI    OpCode = 0x57 // Conditional jump
)
