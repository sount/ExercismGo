package listops

// predFunc function of int returns bool
type predFunc func(int) bool

//IntList list of integers
type IntList []int

//binFunc takes to integers and returns one int
type binFunc func(int, int) int

//unaryFunc takes an inter and returns an int
type unaryFunc func(int) int

//Filter takes a list and a function f, then  filters all element for which f(e) == false
func (list IntList) Filter(p predFunc) IntList {
	result := []int{}
	for i := range list {
		if p(list[i]) {
			result = append(result, list[i])
		}
	}
	return result
}

//Length returns the length of a list
func (list IntList) Length() int {
	var count int
	for range list {
		count++
	}
	return count
}

// Map applies a function to a list
func (list IntList) Map(u unaryFunc) IntList {
	new_list := []int{}
	for i := range list {
		new_list = append(new_list, u(list[i]))
	}
	return new_list
}

// Reverse a list
func (list IntList) Reverse() IntList {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	return list
}

// Append a list to a list
func (list IntList) Append(appending_list IntList) IntList {
	return append(list, appending_list...)

}

//Concat many lists
func (list IntList) Concat(liste []IntList) IntList {
	for i := range liste {
		list = append(list, liste[i]...)
	}
	return list
}

// Foldl very confusing functions
func (list IntList) Foldl(b binFunc, acc int) int {

	for i := range list {

		acc = b(acc, list[i])
	}
	return acc
}

//Foldr same like Foldl
func (list IntList) Foldr(b binFunc, acc int) int {
	for i := len(list) - 1; i >= 0; i-- {

		acc = b(list[i], acc)
	}
	return acc
}
