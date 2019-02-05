package coverme

import "strconv"

// Excercise (easy): get to 100% test coverage on FizzBuzz with a single table test.
// While it is possible to get to 100% coverage with 3 inputs, the tests should also
// exhaust all possible branches (more than 3).

func FizzBuzz(n int) string {
	var s string
	if n%3 == 0 {
		s = "fizz"
	}
	if n%5 == 0 {
		s += "buzz"
	}
	if s == "" {
		return strconv.Itoa(n)
	}
	return s
}
