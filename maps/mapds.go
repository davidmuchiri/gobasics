package maps

import "fmt"

/*
*MAPS
- a map maps keys to values
? the zero value of a map is null
! a nil map has no keys nor can keys be added
- the make function returns a map of the given type initialized and ready for use
*/

//Coordinates is a lat lng structure
type Coordinates struct {
	lat, lng float64
}

/*
*	The variable below is an initializes a map
? The keys are of type string and their values are structures of type Coordinates
*/
var m map[string]Coordinates

//MapOne simple mapFunction
func MapOne() {
	m = make(map[string]Coordinates)
	m["Priwanna apartments"] = Coordinates{
		40.68533, -74.39967,
	}
	fmt.Println(m["Priwanna apartments"])
}

// MapTwo is a map literal
func MapTwo() {
	m := map[string]Coordinates{
		"Nairobi": Coordinates{
			36.03343, -73.23434,
		},
		"Google": Coordinates{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
}

/*
* MUTATING MAPS
? inserting or updating an element in a map
 - m[key] = elem
? retrieve an element
- elem = m[key]

? deleting an element
- delete(m, key)

? test that a key is present with a two-value assignment
-- elem, ok = m[key]
- if key is in m , ok is true. if not, ok is false

*/

// MapMutation addresses mutations of maps
func MapMutation() {
	m := make(map[string]int)

	m["answer"] = 42 // adding a key and value
	fmt.Println("The value: ", m["answer"])

	m["answer"] = 48 // setting key to another value
	fmt.Println("The value:", m["answer"])

	delete(m, "answer") // deleting a key from map
	fmt.Println("The value:", m["answer"])

	v, ok := m["Answer"]
	fmt.Println("The value: ", v, "Present?", ok)
}

//TODO: maps exercise
// create a function that takes in a word and returns a map of all letters in the word and how many times they are repeated

//WordCount takes int a word and returns a map of letters and their count
func WordCount(str string) map[string]int {

	return map[string]int{"x": 1}
}
