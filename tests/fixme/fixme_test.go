package fixme

import (
	"sync"
	"testing"
)

// test with `go test -race ./`

func TestSumOfAll(t *testing.T) {
	tests := []int{1, 2, 7, 15, 1000}
	summer := func(n int) int {
		return (n * (n - 1)) / 2
	}

	for _, tt := range tests {
		got := SumOfAll(tt)
		want := summer(tt)
		if got != want {
			t.Errorf("got: %d, want %d", got, want)
		}
	}
}

func SumOfAllSolution(n int) (sum int) {
	var mu sync.Mutex
	var wg sync.WaitGroup

	tot := 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			i := i
			wg.Add(1)
			go func() {
				mu.Lock()
				defer mu.Unlock()
				defer wg.Done()
				tot += i
			}()
			continue
		}
		mu.Lock()
		tot += i
		mu.Unlock()
	}
	wg.Wait()
	return tot
}
