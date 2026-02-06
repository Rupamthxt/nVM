package main

import (
	"math/big"
	"sync"
)

const StackLimit = 1024

type Stack struct {
	data []*big.Int
	mu   sync.RWMutex
}

func NewStack() *Stack {
	return &Stack{
		data: make([]*big.Int, 0, StackLimit),
	}
}

// Push inserts a new element onto the stack. If the stack exceeds the StackLimit, it panics with "stack overflow".
func (s *Stack) Push(d *big.Int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.data) >= StackLimit {
		panic("stack overflow")
	}
	s.data = append(s.data, d)
}

// Pop removes and returns the top element from the stack. If the stack is empty, it panics with "Stack underflow".
func (s *Stack) Pop() *big.Int {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.data) == 0 {
		panic("Stack underflow")
	}

	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item
}

// Peek returns the top element of the stack without removing it. If the stack is empty, it returns nil.
func (s *Stack) Peek() *big.Int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.data) == 0 {
		return nil
	}

	return s.data[len(s.data)-1]
}
