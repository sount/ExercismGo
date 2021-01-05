package etl

import "strings"

// Transform takes a map (dict in python) containing of slices of strings and returns the integer value also as a mao
func Transform(input map[int][]string) map[string]int {

	final_map := make(map[string]int)
	for i := range input {
		for j := range input[i] {
			final_map[strings.ToLower(input[i][j])] = i
		}
	}

	return final_map
}
