package vm

import "unsafe"

func (s *Stack) PushInt64(f int64) {
	s.data = append(s.data, *(*uint64)(unsafe.Pointer(&f)))
}

func (s *Stack) PopInt64() int64 {
	pivot := len(s.data) - 1
	i := *(*int64)(unsafe.Pointer(&s.data[pivot]))
	s.data = s.data[:pivot]
	return i
}

func (s *Stack) AddInt64() {
	pivotA := len(s.data) - 1
	pivotB := len(s.data) - 2
	a := *(*int64)(unsafe.Pointer(&s.data[pivotA]))
	b := *(*int64)(unsafe.Pointer(&s.data[pivotB]))
	*(*int64)(unsafe.Pointer(&s.data[pivotB])) = a + b
	s.data = s.data[:pivotA]
}

func (s *Stack) MulInt64() {
	pivotA := len(s.data) - 1
	pivotB := len(s.data) - 2
	a := *(*int64)(unsafe.Pointer(&s.data[pivotA]))
	b := *(*int64)(unsafe.Pointer(&s.data[pivotB]))
	*(*int64)(unsafe.Pointer(&s.data[pivotB])) = a * b
	s.data = s.data[:pivotA]
}

func (s *Stack) DivInt64() {
	pivotA := len(s.data) - 1
	pivotB := len(s.data) - 2
	a := *(*int64)(unsafe.Pointer(&s.data[pivotA]))
	b := *(*int64)(unsafe.Pointer(&s.data[pivotB]))
	*(*int64)(unsafe.Pointer(&s.data[pivotB])) = a / b
	s.data = s.data[:pivotA]
}

func (s *Stack) SubInt64() {
	pivotA := len(s.data) - 1
	pivotB := len(s.data) - 2
	a := *(*int64)(unsafe.Pointer(&s.data[pivotA]))
	b := *(*int64)(unsafe.Pointer(&s.data[pivotB]))
	*(*int64)(unsafe.Pointer(&s.data[pivotB])) = a - b
	s.data = s.data[:pivotA]
}
