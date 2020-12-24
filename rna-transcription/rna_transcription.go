package strand

//ToRNA takes a dna string an returns the RNA string
func ToRNA(dna string) string {
	if dna == "" {
		return ""
	}
	var output string
	for s := range dna {
		if string(dna[s]) == "G" {
			output += "C"
		} else if string(dna[s]) == "C" {
			output += "G"
		} else if string(dna[s]) == "T" {
			output += "A"
		} else if string(dna[s]) == "A" {
			output += "U"
		}
	}
	return output
}
