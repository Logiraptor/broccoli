package vm

import "testing"

func BenchmarkControlCopy(b *testing.B) {
	s := NewStack(100)
	x := NewStack(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.CopyInto(x)
	}
}

func BenchmarkAddNative(b *testing.B) {
	s := NewStack(100)
	s.PushFloat64(4.3)
	s.PushFloat64(2.7)
	x := NewStack(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.CopyInto(x)
		_ = 4.3 + 2.7
	}
}

func TestPushFloat64Unsafe(t *testing.T) {
	s := NewStack(100)

	s.PushFloat64(4.3)
	s.PushFloat64(3.5)
	s.AddFloat64()

	if x := s.PopFloat64(); x != 7.8 {
		t.Fatal(x)
	}
}

func BenchmarkAddFloat64Unsafe(b *testing.B) {
	s := NewStack(100)
	x := NewStack(100)
	s.PushFloat64(4.3)
	s.PushFloat64(2.7)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.CopyInto(x)
		x.AddFloat64()
	}
}
