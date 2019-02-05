package myfunc

import "testing"

// Run with `go test -bench=. ./ -v`

func testTheAnswer(tb testing.TB) {
	if got := theAnswer(); got != 42 {
		tb.Errorf("got %d, want 42", got)
	}
}
func TestTheAnswer(t *testing.T) {
	testTheAnswer(t)
}
func BenchmarkTheAnswer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testTheAnswer(b)
	}
}

func TestMyBar(t *testing.T) {
	tests := []struct {
		name, in, want string
	}{{
		name: "foo",
		in:   "this is a foo",
		want: "this is a bar",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myBar(tt.in); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestHazard(t *testing.T) {
	if got := hazard(); got != 1 {
		t.Errorf("got %v, want 1", got)
	}
}
