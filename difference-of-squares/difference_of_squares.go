package diffsquares

func SquareOfSum( x int)int{
	var k = 0
	for i:= 0; i<=x; i++{
		k = k+ i
	}

	return k*k
}
func SumOfSquares(x int) int{
	var k =0
	for i:=0; i<=x; i++{
		k = k +(i*i)
	}
	return k
}

func Difference(x int) int {

	return SquareOfSum(x)-SumOfSquares(x)
}
