package conditionals

import (
	"fmt"
	"runtime"
	"time"
)

//SwitchOne is a function that prints out the operating system
func SwitchOne() {
	fmt.Println("Go is running on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X. ")
	case "linux":
		fmt.Println("Linux. ")
	default:
		fmt.Printf("%s. ", os)
	}
}

//SwitchTwo is a function that checks how close to monday today is
func SwitchTwo() {
	fmt.Print("When's Monday ? ")
	today := time.Now().Weekday()

	switch time.Monday {
	case today + 0:
		fmt.Println("today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away")
	}
}

//SwitchThree checks what hour it is and prints a message based on what part of the day it is
func SwitchThree() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	case t.Hour() < 19:
		fmt.Println("Good evening")
	case t.Hour() < 24:
		fmt.Println("Good night", t.Hour())
	default:
		fmt.Println("sleep")
	}
}
