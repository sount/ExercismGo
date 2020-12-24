package pythagorean

// Triplet pythagorean
type Triplet []int

// Sum gets an integer N and returns all pythagorean triplets with a+b+c=N
func Sum(input int) []Triplet {

	var triplets []Triplet
	for c := 0; c < input; c++ {
		for b := 0; b < c; b++ {
			for a := 0; a < b; a++ {
				if a+b+c == input && a*a+b*b == c*c {
					t := Triplet{a, b, c}
					triplets = append(triplets, t)
				}
			}
		}
	}
	for i, j := 0, len(triplets)-1; i < j; i, j = i+1, j-1 {
		triplets[i], triplets[j] = triplets[j], triplets[i]
	}
	return triplets
}

// Range gets two integers and returns all pythagorean triplets in this range
func Range(x, y int) []Triplet {
	var triplets []Triplet
	for c := x; c <= y; c++ {
		for b := x; b < c; b++ {
			for a := x; a < b; a++ {
				if a*a+b*b == c*c {
					t := Triplet{a, b, c}
					triplets = append(triplets, t)
				}
			}
		}
	}
	return triplets
}
