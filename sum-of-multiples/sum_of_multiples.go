package summultiples



func Ints(input []int) []int {
	u := make([]int, 0, len(input))
	m := make(map[int]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}
	return u
}
func sum_up(input []int) int {
	sum := 0

	for i := range input {
		sum += input[i]
	}
	return sum
}
func find_divisors(limit int, div int) []int {
	var s []int
	for i := 0; i < limit; i++ {
		if (div * i) >= limit {
			break
		}
		s = append(s, i*div)

	}
	return s
}
func SumMultiples(limit int, divisors ...int) int {

	var sum []int
	for _, val := range divisors {

		sum = append(sum, find_divisors(limit, val)...)
	}
	temp :=Ints(sum)
	return sum_up(temp)
}

