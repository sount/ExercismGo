
package leap

// IsLeapYear returns if given year is a leap year
func IsLeapYear(x int) bool {
	val := float64(x) / 4
	if val != float64(int(val)){
		return false
	}
	val2 := float64(x) / 100
	if val2 == float64(int(val2)){
		val3 := float64(x) / 400
		if val3 == float64(int(val3)) {
			return true
		}
		return false
	}
	return true
}
