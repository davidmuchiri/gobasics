package arrays

import "fmt"

/*
* Slices
 - Slices are dynamically sized, flexible view into the elements of an array
 - Slices are used more often than arrays
 - Type []T is a slice with elements of type T
 - a slice if formed by specifying two indices, a low and high bound separated by a colon
 * a[low:high]

 * Slice literals
 - A slice literal is like an array literal without the length
 ? example of an array literal [3]bool{true, true, false}

- This reates the same array as above then builds a slice that references it:
?  Example of a slice literal []bool{true, true, true}

*/

//SliceOne is a function that demonstrates the concept of a slice
func SliceOne() {
	fruits := [5]string{"apples", "banana", "mango", "melon", "pineapple"}
	var fruitsSlice = fruits[0:4] // will return an "array" with apples bananas mango and melon
	fmt.Println(fruitsSlice)
}

/*
 ! A slice does not store any data, it just describes a section of an underlying array
 ! Changing the elements os a slice modifies the corresponding elements of its underlying array
 ! Other slices that share the same underlying array will see those changes
*/

// SliceTwo illustrates the statements above
func SliceTwo() {
	names := [4]string{"john", "paul", "george", "Ringo"}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "xxx"
	fmt.Println(a, b)
	fmt.Println(names)

}
