package scrabble

import "strings"

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Score(word string) int {

	one_point := []string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"}
	two_point := []string{"D", "G"}
	three_point := []string{"B", "C", "M", "P"}
	four_point := []string{"F", "H", "V", "W", "Y"}
	five_point := []string{"K"}
	eight_point := []string{"J", "X"}
	ten_point := []string{"Q", "Z"}
	word = strings.ToUpper(word)
	var score int
	for i := 0; i < len(word); i++ {
		if contains(one_point, string(word[i])) {
			score += 1
		}
		if contains(two_point, string(word[i])) {
			score += 2
		}
		if contains(three_point, string(word[i])) {
			score += 3
		}
		if contains(four_point, string(word[i])) {
			score += 4
		}
		if contains(five_point, string(word[i])) {
			score += 5
		}
		if contains(eight_point, string(word[i])) {
			score += 8
		}
		if contains(ten_point, string(word[i])) {
			score += 10
		}

	}
	return score
}
