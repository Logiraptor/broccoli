package vm

type VM struct {
	stack *Stack
}

func NewVM() *VM {
	return &VM{
		stack: NewStack(100),
	}
}

func (v *VM) Exec(source []uint64) {
	var (
		pc  int
		op  uint64
		arg uint64
	)

	for pc < len(source) {
		op, arg = splitOp(source[pc])
		pc++
		switch op {
		case HALT:
			break
		case LOAD:
			panic("LOAD not implemented")
		case STORE:
			panic("STORE not implemented")
		case LOADI:
			v.stack.PushInt64(int64(arg))
		case RET:
			panic("RET not implemented")
		case ADD:
			v.stack.AddInt64()
		}
	}
}
