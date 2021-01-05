package accumulate

// Accumulate takes a slice of strings and a function and returns a slice of strings
func Accumulate(slice []string, f func(string) string) []string {
	var result []string
	for i := range slice {
		result = append(result, f(slice[i]))
	}
	return result
}
