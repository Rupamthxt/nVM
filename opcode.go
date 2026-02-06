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
)
