package datatypes

import "fmt"

func Arrays() {
	// var a [5]int //length 5 of integers. Index starts from 0 to 4
	// fmt.Println(a)
	// fmt.Println("Length of array a:", len(a))
	// //Insert an element at a particular position
	// // n is the length and n-1 is the last index
	// a[0] = 10
	// a[1] = 20
	// a[2] = 30
	// a[3] = 40
	// a[4] = 50
	// fmt.Println("Array is:", a)
	// fmt.Println("Last element of array a:", a[len(a)-1])
	// fmt.Println("Third element of array a:", a[2])
	// a[2] = 25
	// fmt.Println("Updated array is:", a)
	//Iteration using for loop
	for i := 0; i <= 5; i++ {
		// if i == 2 {
		// 	fmt.Println("value of i which is: ", i)
		// } else if i == 4 {
		// 	fmt.Println("value of i: ", i)
		// } else {
		// 	fmt.Println("i: ", i)
		// }

		// fmt.Println("i: ", i)

		// break exits the loop on a particular condition
		// if i == 3 {
		// 	fmt.Println("Reached i=3, exit the loop")
		// 	break
		// }

		//continue skips the current iteration and moves to the next iteration
		if i == 3 {
			fmt.Println("Skipping i=3")
			continue
		}
		fmt.Println("i: ", i)

	}
}
