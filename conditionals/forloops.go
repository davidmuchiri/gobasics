package conditionals

import "fmt"

// ForloopOne loops over 10 numbers and adds them up
func ForloopOne() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

//ForloopTwo loops over the passed in numbers and adds them up
func ForloopTwo(x int) int {
	sum := 1
	for sum < x {
		sum += sum
	}
	return sum
}

// Forever is a function that loops forever
func Forever() {
	for {
	}
}
