package vm

import (
	"fmt"
)

// OpCodes
const (
	// HALT immediately stops the vm
	HALT = iota
	// LOAD pops a pointer from the stack and
	// copies that word onto the stack.
	LOAD
	// STORE pops a pointer and a word from the stack
	// and stores the word at that address
	STORE
	// LOADI takes 1 argument and pushes it onto the stack
	LOADI
	// RET takes 1 argument n. it pops n values from the stack and returns them
	RET
	// ADD pops two values and pushes the result to the stack
	ADD
)

const (
	OPCODE_SHIFT = 56
	OPCODE_MASK  = 0xff00000000000000
	ARG_MASK     = 0x00ffffffffffffff
)

func splitOp(op uint64) (opcode uint64, arg uint64) {
	opcode = (op & OPCODE_MASK) >> OPCODE_SHIFT
	arg = op & ARG_MASK
	return
}

func makeOp(opcode, arg uint64) uint64 {
	if arg&ARG_MASK != arg {
		panic(fmt.Sprintf("arg %v does not fit in 59 bits", arg))
	}
	return (opcode << OPCODE_SHIFT) | (arg & ARG_MASK)
}
