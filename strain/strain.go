package strain

// Ints is a integer list with methods keep and discard
type Ints []int

// Strings is a string list with methods keep and discard
type Strings []string

// Lists is a list of integer list with methods keep and discard
type Lists [][]int

// Keep for Ints
func (Integers Ints) Keep(function func(int) bool) Ints {
	var result Ints
	for i := range Integers {
		if function(Integers[i]) {
			result = append(result, Integers[i])
		}
	}
	return result
}

// Discard for Ints
func (Integers Ints) Discard(function func(int) bool) Ints {
	var result Ints
	for i := range Integers {
		if !function(Integers[i]) {
			result = append(result, Integers[i])
		}
	}
	return result
}

// Keep for Strings
func (String Strings) Keep(function func(string2 string) bool) Strings {
	var result Strings
	for i := range String {
		if function(String[i]) {
			result = append(result, String[i])
		}
	}
	return result
}

//Discard for Strings
func (String Strings) Discard(function func(string) bool) Strings {
	var result Strings
	for i := range String {
		if !function(String[i]) {
			result = append(result, String[i])
		}
	}
	return result
}

//Keep for Lists
func (List Lists) Keep(function func([]int) bool) Lists {
	var result Lists
	for i := range List {
		if function(List[i]) {
			result = append(result, List[i])
		}
	}
	return result
}

// Discard for Lists
func (List Lists) Discard(function func([]int) bool) Lists {
	var result Lists
	for i := range List {
		if !function(List[i]) {
			result = append(result, List[i])
		}
	}
	return result
}
