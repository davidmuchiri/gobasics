package functions

import (
	"fmt"
	"math"
)

//Add function takes two numbers and adds them up
func Add(x, y int) int {
	return x + y
}

// Swap function takes two strings and swaps them
func Swap(x, y string) (string, string) {
	return y, x
}

// Split takes a number and splits it into two
func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//Def is a function that uses the defer statement a defer statement runs after everything else has run
func Def() {
	defer fmt.Print(" world")
	fmt.Print("hello")
}

//DefStacking  defer function calls are pushed onto a stack. when a function returns, its deferred calls are executed in LIFO
func DefStacking() {
	fmt.Println("Counting")
	for i := 1; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")

}

/*
* passing functions as callbalks
 */
// Functions can be passed around just like other values
// Function values may be used as function arguments and return values

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// PassedFunc is a function used to illustrate the concept of callbacks in go
func PassedFunc() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))

}

/*
* Functions as closures
Go function may be closures. A closure is a function value that references variables from outside its body.
the function may access and assign to to the referenced variables
the function is bound to the variables
*/

func adder() func(int, int) int {
	sum := 0

	return func(x int, y int) int {
		sum = x + y
		return sum
	}

}

// AddNumbers is a function that is used to illustrate the concept of closures
func AddNumbers() {
	addfunc := adder()
	addfunc(1, 2)
}

//TODO: FUN with functions

func Calculator() map[string]func(int, int) int {

	calculate := make(map[string]func(int, int) int)

	calculate["add"] = func(x int, y int) int {
		return x + y
	}
	calculate["subtract"] = func(x int, y int) int {
		return x - y
	}

	calculate["multiply"] = func(x int, y int) int {
		return x * y
	}

	calculate["divide"] = func(x int, y int) int {
		return x / y
	}

	return calculate
}
