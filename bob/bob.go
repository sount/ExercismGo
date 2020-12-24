package bob

import "strings"
// NoLetter checks if the string contains only letters and if yes, it returns true otherwise false
func NoLetter(s string) bool {
	var k int
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			k++
		}
	}
	if k == len(s) {
		return true
	}
	return false

}


// Hey is returning bobs answers to various strings
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	if len(remark) == 0 {
		return "Fine. Be that way!"
	}

	is_question := string(remark[len(remark)-1]) == "?"
	is_upper := strings.ToUpper(remark) == remark
	no_letters := NoLetter(remark)

	if no_letters {
		if is_question {
			return "Sure."
		}
		return "Whatever."

	}
	if is_question && !is_upper {
		return "Sure."
	} else if !is_question && is_upper {
		return "Whoa, chill out!"
	} else if is_question && is_upper {
		return "Calm down, I know what I'm doing!"
	}
	return "Whatever."

}
