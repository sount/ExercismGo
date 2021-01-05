package dna

import "fmt"

// Histogram the return type like a dict
type Histogram map[rune]int

// DNA contains the information to check
type DNA string

// Counts gets a string and counts the occurrences of the different letters
func (d DNA) Counts() (Histogram, error) {
	var h Histogram
	h = Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}
	for i := range d {

		if string(d[i]) == "A" {
			h['A']++
		} else if string(d[i]) == "C" {
			h['C']++
		} else if string(d[i]) == "G" {
			h['G']++
		} else if string(d[i]) == "T" {
			h['T']++
		} else {
			return h, fmt.Errorf("unexpected caracter %s in input %s", string(d[i]), d)
		}
	}

	return h, nil
}
