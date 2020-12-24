package lsproduct

import (
	"errors"
	"strconv"
)
// LargestSeriesProduct gets a string containing numbers and a range, returns the highest product of a substring
func LargestSeriesProduct(number string, span int) (int, error) {
	if span > len(number) || span < 0 {
		return 0, errors.New("Wrong Input")
	}
	var integerSlice []int
	for i := 0; i < len(number); i++ {

		temp, err := strconv.Atoi(string(number[i]))

		if err != nil {
			return 0, errors.New("NaN")
		}
		integerSlice = append(integerSlice, temp)
	}

	var highest = 0
	for i := 0; i <= len(integerSlice)-span; i++ {
		var product = 1
		for j := 0; j < span; j++ {
			product *= integerSlice[i+j]
		}
		if product > highest {
			highest = product
		}
	}

	return highest, nil
}
