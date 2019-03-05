package pointers

import "fmt"

// A pointer holds the memory adress of a value
// The type *T is a pointer to a T value
// Its zero value is nul
// var p *int
// the & operator generates a pointer to its operand
// i := 2
// p = &i
// the * operator denotes the pointer's underlying value
//fmt.Println(*p)
// *p = 21 with reassign i to 21

//PointerOne simple pointer function
func PointerOne() {
	i, j := 42, 2701

	p := &i         // variable p points to var i
	fmt.Println(*p) // read var i through pointer p
	*p = 21         // reassign i throught the pointer p
	fmt.Println(*p) //variable i should be a new variable

	p = &j       //variable p not points to variable j
	*p = *p / 37 // reassign variable j through pointer p j is set to itself divided by 37
	fmt.Println(j)

}
