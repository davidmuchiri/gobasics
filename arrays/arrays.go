package arrays

import "fmt"

/*
*Arrays
- The type [n]T is an array of n values of type T
- Example books := [5]string An array of five books
- An array's length is part of its type
!an array cannot be resized
*/

//ArrOne is a function that illustrates a basic implementation of arrays in go
func ArrOne() {
	var a [2]string

	a[0] = "Hello"
	a[1] = "Konichiwa"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	bed := [2]string{"my", "bed"}
	fmt.Println(bed)

}
