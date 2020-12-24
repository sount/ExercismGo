package prime

func is_prime(prime int ) bool{

	for i:=2; i<prime;i++{
		if prime % i==0{
			return false
		}
	}
	return true
}
//Nth gets an integer n and then reutrns the nth prime
func Nth( input int) (int, bool){
	if input < 1{
		return input, false
	}
	var nth int
	for i:=2;; i++ {
		if is_prime(i){
			nth++
			if nth == input {
				return i, true
			}
		}
	}
	return input, false

}