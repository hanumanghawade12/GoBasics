package main

import "fmt"

func Array() {

	// declare array variable of type integer
	// defined size [size=5]
	//   var arrayOfInteger = [5]int{1, 5, 8, 0, 3}
	//   var array = [2]string{"Go", "Conference"}

	//   fmt.Println(arrayOfInteger[0], array[0])
	// Initialize an Array in Golang
	//    var array1 [5]int
	//    array1[0] = 10
	//    array1[1] = 20
	//    array1[2] = 30
	//    array1[3] = 40
	//    array1[4] = 50

	//    fmt.Println("Array1:", array1)

	// create two slices
	evenNumbers := []int{2, 4}
	oddNumbers := []int{1, 3}

	// add elements of oddNumbers to evenNumbers
	evenNumbers = append(evenNumbers, oddNumbers...)
	fmt.Println("Numbers:", evenNumbers)
}
