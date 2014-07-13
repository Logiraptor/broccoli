package vm

// type Word uint64

type Stack struct {
	data []uint64
	temp uint64
}

func NewStack(size int) *Stack {
	return &Stack{
		data: make([]uint64, 0, size),
	}
}

func (s *Stack) Copy() *Stack {
	b := NewStack(cap(s.data))
	b.data = b.data[:len(s.data)]
	copy(b.data, s.data)
	return b
}

func (s *Stack) CopyInto(b *Stack) {
	if cap(b.data) < len(s.data) {
		b.data = make([]uint64, 0, len(s.data))
	}
	b.data = b.data[:len(s.data)]
	copy(b.data, s.data)
}
