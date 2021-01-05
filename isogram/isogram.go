package isogram

import "strings"

// IsIsogram takes a string (word) and return´s if it is an isogram or noct
func IsIsogram(isogram string) bool {
	isogram = strings.ToUpper(isogram)

	for i := 0; i < len(isogram); i++ {
		if string(isogram[i]) == "-" || string(isogram[i]) == " " {
			continue
		}
		for j := 0; j < i; j++ {
			if isogram[i] == isogram[j] {
				return false
			}
		}

	}
	return true

}
