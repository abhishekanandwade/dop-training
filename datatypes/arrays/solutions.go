package arrays

import "fmt"

// Sum all numbers from 1 to 100 and return the total sum.
func sumNumbers() int {
	total := 0
	for i := 1; i <= 100; i++ {
		total += i
	}
	return total
}

// Print all even numbers from 1 to 100 using a loop.
func printEvenNumbers() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Print(" ", i)
		}
	}
}

// Print all odd numbers from 1 to 100 using a loop.
func printOddNumbers() {
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			fmt.Print(" ", i)
		}
	}
}

func LoopsSolutions() {
	// Call the functions to demonstrate their functionality
	fmt.Println("Sum of numbers from 1 to 100:", sumNumbers())
	fmt.Println("")
	fmt.Print("Even numbers from 1 to 100: ")
	printEvenNumbers()
	fmt.Println("")
	fmt.Println("")
	fmt.Print("Odd numbers from 1 to 100: ")
	printOddNumbers()
}
