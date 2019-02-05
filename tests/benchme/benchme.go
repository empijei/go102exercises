package benchme

import "bytes"

// Excercise (hard): write a benchmark for SimpleStrCat and BufferedStrCat, and see
// for how many strings the buffered solution becomes faster. Pick a size for the length
// of the strings you are concatenating.
// Hint: it is less than 15.
// Hint: if you are writing more than 30 lines you are not doing it right.

func SimpleStrCat(all ...string) string {
	var ret string
	for _, s := range all {
		ret += s
	}
	return ret
}
func BufferedStrCat(all ...string) string {
	// If you are using modern go, please use strings.Builder when concatenating strings,
	// it is always faster than summing three or more strings.
	var buf bytes.Buffer
	for _, s := range all {
		buf.Write([]byte(s))
	}
	return buf.String()
}
