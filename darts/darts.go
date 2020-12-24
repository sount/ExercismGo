package darts

import "math"

// Score takes the x and y coordinates and returns and integer with the score
func Score(x, y float64) int {

	r := math.Sqrt(x*x + y*y)

	if r <= 1 {
		return 10
	} else if r > 1 && r <= 5 {
		return 5
	} else if r > 5 && r <= 10 {
		return 1
	}
	return 0

}
