
package triangle

import "math"

type Kind string

const (
    // Pick values for the following identifiers used by the test program.
    NaT = "NaT"// not a triangle
    Equ = "Equ" // equilateral
    Iso = "Iso"// isosceles
    Sca = "Sca"// scalene
)


func check_inequalty(a,b,c float64) bool{
	var resp = true
	if a+b<c || a+c<b || b+c<a{
		resp = false
	}
	return resp
}

func check_nan_inf(a ... float64) bool{
	var k = false
	for i:=0; i<len(a);i++{
		if math.IsNaN(a[i]){
			k = true
		}
		if math.IsInf(a[i],1){
			k = true
		}
		if math.IsInf(a[i],-1){
			k = true
		}

	}
	return k

}
// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if !check_inequalty(a,b,c) || check_nan_inf(a,b,c) || c<=0 {
		return NaT
	}

	if a == b && c==a {

		return Equ
	}
	if a==b || a == c || b ==c{
		return Iso
	}

	return Sca


}
