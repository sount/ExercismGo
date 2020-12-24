package main

import (
	"strconv"
	"strings"
)

//IsValidISBN checks if a given string is a valid ISBN
func IsValidISBN(isbn string) bool {
	var sum int
	isbn = strings.Replace(isbn, "-", "", -1)

	if len(isbn) != 10 {
		return false
	}
	for i := 0; i < len(isbn); i++ {

		if string(isbn[i]) == "X" && i==9 {
			sum += 10 * (10 - i)
		} else {
			temp, err := strconv.Atoi(string(isbn[i]))
			if err != nil {
				return false
			}
			sum += temp * (10 - i)
		}
	}
	if sum%11 == 0 {
		return true
	}
	return false
}

func main() {

	IsValidISBN("3-598-P1508-8")
}
