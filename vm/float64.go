package vm

import "unsafe"

func (s *Stack) PushFloat64(f float64) {
	s.data = append(s.data, *(*uint64)(unsafe.Pointer(&f)))
}

func (s *Stack) PopFloat64() float64 {
	pivot := len(s.data) - 1
	i := *(*float64)(unsafe.Pointer(&s.data[pivot]))
	s.data = s.data[:pivot]
	return i
}

func (s *Stack) AddFloat64() {
	pivotA := len(s.data) - 1
	pivotB := len(s.data) - 2
	a := *(*float64)(unsafe.Pointer(&s.data[pivotA]))
	b := *(*float64)(unsafe.Pointer(&s.data[pivotB]))
	*(*float64)(unsafe.Pointer(&s.data[pivotB])) = a + b
	s.data = s.data[:pivotA]
}

func (s *Stack) MulFloat64() {
	pivotA := len(s.data) - 1
	pivotB := len(s.data) - 2
	a := *(*float64)(unsafe.Pointer(&s.data[pivotA]))
	b := *(*float64)(unsafe.Pointer(&s.data[pivotB]))
	*(*float64)(unsafe.Pointer(&s.data[pivotB])) = a * b
	s.data = s.data[:pivotA]
}

func (s *Stack) DivFloat64() {
	pivotA := len(s.data) - 1
	pivotB := len(s.data) - 2
	a := *(*float64)(unsafe.Pointer(&s.data[pivotA]))
	b := *(*float64)(unsafe.Pointer(&s.data[pivotB]))
	*(*float64)(unsafe.Pointer(&s.data[pivotB])) = a / b
	s.data = s.data[:pivotA]
}

func (s *Stack) SubFloat64() {
	pivotA := len(s.data) - 1
	pivotB := len(s.data) - 2
	a := *(*float64)(unsafe.Pointer(&s.data[pivotA]))
	b := *(*float64)(unsafe.Pointer(&s.data[pivotB]))
	*(*float64)(unsafe.Pointer(&s.data[pivotB])) = a - b
	s.data = s.data[:pivotA]
}
