package fixme

import (
	"sync"
)

// Excercise (medium): write and run a test that finds the 2 race conditions and the deadlock in the following function, and fix them.

func SumOfAll(n int) (sum int) {
	var (
		mu  sync.Mutex
		wg  sync.WaitGroup
		tot int
	)
	wg.Add(n)

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			go func() {
				mu.Lock()
				defer mu.Unlock()
				defer wg.Done()
				tot += i
			}()
			continue
		}
		tot += i
	}
	wg.Wait()
	return tot
}
