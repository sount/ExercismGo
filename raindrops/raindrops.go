package raindrops

import "strconv"

// Convert takes an integer value representing the raindrop and returns a string with the sounds of the raindrop
// if the integer has no sound it returns the input integer as a string
func Convert(raindrop int) string {
	var result string
	if raindrop%3 == 0 {
		result += "Pling"
	}
	if raindrop%5 == 0 {
		result += "Plang"
	}
	if raindrop%7 == 0 {
		result += "Plong"
	}
	if result == "" {
		return strconv.Itoa(raindrop)
	}
	return result
}
