package slices

import "fmt"

// 1. Create a slice of length 10 and capacity 20. Print its length and capacity.
func question1() {
	s := make([]int, 10, 20)
	fmt.Printf("Length: %d, Capacity: %d\n", len(s), cap(s))
}

// 2. Append 15 integers to an empty slice and print the slice, its length, and capacity after each append operation.
func question2() {
	s := []int{}
	for i := 1; i <= 15; i++ {
		s = append(s, i)
		fmt.Printf("After appending %d: slice=%v, length=%d, capacity=%d\n", i, s, len(s), cap(s))
	}
	// b := [15]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// for i, v := range b {
	// 	s = append(s, v)
	// 	fmt.Printf("After appending index %d, value=%d, slice=%v, length=%d, capacity=%d\n", i, v, s, len(s), cap(s))
	// }
}

// 3. Write a function that takes a slice of integers and returns a new slice containing only the even numbers from the original slice.
func filterEven(nums []int) []int {
	evens := []int{}
	for _, n := range nums {
		if n%2 == 0 {
			evens = append(evens, n)
		}
	}
	return evens
}

func SliceExamples() {
	// question1()
	// question2()
	fmt.Println("Even numbers:", filterEven([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
