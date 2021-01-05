package reverse

// Reverse takes a string and returns the inverse string
func Reverse(str string) string {

	runestr := []rune(str)
	for i, j := 0, len(runestr)-1; i < j; i, j = i+1, j-1 {
		runestr[i], runestr[j] = runestr[j], runestr[i]
	}

	return string(runestr)
}
