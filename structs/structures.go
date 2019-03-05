package structs

import "fmt"

// Coder this is a sample go structure
type Coder struct {
	name      string
	age       int
	languages []string
}

//Structurers this function defines a basic structure implementation
func Structurers() {
	david := Coder{name: "david muchiri", age: 22, languages: []string{"javascript", "golang"}}
	emptyCoder := Coder{}
	emptyAge := Coder{name: "david muchiri", languages: []string{"javascript", "golang"}}
	emptyName := Coder{age: 22, languages: []string{"javascript", "golang"}}

	pointer := &david

	fmt.Println("My name is", david.name)
	fmt.Println("My age is", pointer.age)
	fmt.Println("programming languages that i know are: ", pointer.languages[0], "and", pointer.languages[1])
	fmt.Println("no coder", emptyCoder)
	fmt.Println("no age", emptyAge)
	fmt.Println("no name", emptyName)
}
