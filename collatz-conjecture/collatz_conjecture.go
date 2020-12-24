package collatzconjecture

import "fmt"

// CollatzConjecture gets an integer and return the number of steps to solve the 3n+1 problem
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("error number is <=0")
	}

	for i := 0; ; i++ {
		if n == 1 {
			return i, nil
		}
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}

	}

}
