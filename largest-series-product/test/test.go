package main

import (
	"strconv"
)


func insertionsort(items []int) []int{
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] < items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}

	return items
}


func main() {
	number := "1233"
	span := 2
	var integer_slice []int
	for i:= 0 ; i <len(number);i++ {

		temp,err:=strconv.Atoi(string(number[i]))

		if err != nil{
			//return 0, errors.New("NaN")
		}
		integer_slice = append(integer_slice,temp)
		print(integer_slice)
	}

	sorted_slice := insertionsort(integer_slice)
	print(sorted_slice)
	var product = 1
	for i:=0; i<span;i++{
		product *= sorted_slice[i]
	}
	//return product,nil
}
