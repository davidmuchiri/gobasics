package main

import (
	"fmt"

	"github.com/dinobambino7/gobasics/functions"
	"github.com/dinobambino7/gobasics/maps"
)

func main() {
	fmt.Println(functions.Add(1, 2))
	maps.MapOne()
	maps.MapTwo()
	maps.MapMutation()

	calculator := functions.Calculator()

	add := calculator["add"]
	subtract := calculator["subtract"]
	multiply := calculator["multiply"]
	divide := calculator["divide"]

	fmt.Println(add(3, 3))
	fmt.Println(subtract(4, 2))
	fmt.Println(multiply(3, 3))
	fmt.Println(divide(4, 2))
}
