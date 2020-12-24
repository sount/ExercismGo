package acronym

import (
	"strings"
)

func Abbreviate(s string) string {
	someString := strings.ReplaceAll(strings.ReplaceAll(s, "-", " "), "_", " ")

	words := strings.Fields(someString)
	var k string
	for i := 0; i < len(words); i++ {

		k = k + strings.ToUpper(string(words[i][0]))

	}
	return k
}
