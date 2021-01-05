package romannumerals

import (
	"fmt"
)

// ToRomanNumeral takes an integer and return the corresponding romanian number
func ToRomanNumeral(roman int) (string, error) {
	var result string
	if roman <= 0 || roman > 3000 {
		return "Wrong Input", fmt.Errorf("Wrong Input for %d", roman)
	}
	for roman != 0 {
		if roman >= 1000 {
			roman -= 1000
			result += "M"
		} else if roman >= 900 {
			roman -= 900
			result += "CM"
		} else if roman >= 500 {
			roman -= 500
			result += "D"
		} else if roman >= 400 {
			roman -= 400
			result += "CD"
		} else if roman >= 100 {
			roman -= 100
			result += "C"
		} else if roman >= 90 {
			roman -= 90
			result += "XC"
		} else if roman >= 50 {
			roman -= 50
			result += "L"
		} else if roman >= 40 {
			roman -= 40
			result += "XL"
		} else if roman >= 10 {
			roman -= 10
			result += "X"
		} else if roman >= 9 {
			roman -= 9
			result += "IX"
		} else if roman >= 5 {
			roman -= 5
			result += "V"
		} else if roman >= 4 {
			roman -= 4
			result += "IV"
		} else if roman >= 1 {
			roman--
			result += "I"
		}
	}

	return result, nil
}
