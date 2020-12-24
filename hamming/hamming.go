package hamming

import (
	"fmt"
)

// Distance takes two string and returns  the hamming distance as an integer  and an error
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("length Mismatch for" + a + " and " + b)
	}
	var distance int
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
