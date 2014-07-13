package vm

import "testing"

func TestPushInt64Unsafe(t *testing.T) {
	s := NewStack(100)

	s.PushInt64(4)
	s.PushInt64(3)
	s.AddInt64()

	if x := s.PopInt64(); x != 7 {
		t.Fatal(x)
	}
}

func BenchmarkAddInt64Unsafe(b *testing.B) {
	s := NewStack(100)
	x := NewStack(100)
	s.PushInt64(4)
	s.PushInt64(2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.CopyInto(x)
		x.AddInt64()
	}
}
