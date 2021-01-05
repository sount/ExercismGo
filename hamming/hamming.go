package hamming

import (
	"fmt"
)

// Distance takes two string and returns  the hamming distance as an integer  and an error
func Distance(a, b string) (int, error) {
	ar, br := []rune(a), []rune(b)
	if len(ar) != len(br) {
		return 0, fmt.Errorf("length Mismatch for %s and %s", a, b)
	}

	var distance int
	for i := range ar {
		if ar[i] != br[i] {
			distance++
		}
	}
	return distance, nil
}
