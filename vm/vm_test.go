package vm

import (
	"testing"
)

func TestADD(t *testing.T) {
	vm := NewVM()

	vm.Exec([]uint64{
		makeOp(LOADI, 3),
		makeOp(LOADI, 5),
		makeOp(ADD, 0),
		makeOp(LOADI, 34),
		makeOp(ADD, 0),
	})

	val := vm.stack.PopInt64()
	if val != 42 {
		t.Fatalf("Wanted 42 got %d", val)
	}
}
