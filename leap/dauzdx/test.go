package main


func IsLeapYear(x int) bool {
	val := float64(x) / 4
	if val == float64(int(val)){
		return false
	}
	val2 := float64(x) / 10
	if val2 == float64(int(val2)){
		return false
	}
	return false
}


func main (){

	IsLeapYear(1970)
}
