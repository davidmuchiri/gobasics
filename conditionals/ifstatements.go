package conditionals

import (
	"fmt"
	"math"
)

//Sqrt function gets the square root of the number passed
func Sqrt(x float64) string {
	if x < 0 {
		return Sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

//Pow function
func Pow(numOne, numTwo, numThree float64) float64 {
	if numFour := math.Pow(numOne, numTwo); numFour < numThree {
		return numFour
	}
	return numThree
}

//PowDef function
func powDef(numOne, numTwo, numThree float64) float64 {
	numFour := math.Pow(numOne, numTwo)
	if numFour < numThree {
		return numFour
	}
	return numThree
}
