package main

import "fmt"

func loop() {

	// for item, value := range [5]int{1, 2, 3, 4, 5} {
	// 	fmt.Println(item, value)
	// }

	// name := "Go Conference"

	// for item , value := range name {
	// 	fmt.Println(item, string(value))
	// }
	// create a map
	subjectMarks := map[string]float32{"Java": 80, "Python": 81, "Golang": 85}
	fmt.Println("Marks obtained:")

	// use for range to iterate through the key-value pair
	for subject, marks := range subjectMarks {
		fmt.Println(subject, ":", marks)
	}
	// // For loop
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println("Hello, World", i)
	// }

	// // While loop
	// num := 1
	// for num <= 5 {
	// 	fmt.Println("Hello, World")
	// 	num++
	// }

	// // Infinite loop
	// // for {
	// // 	fmt.Println("Hello, World")
	// // }

	// // Nested loop
	// for i := 1; i <= 5; i++ {
	// 	for j := 1; j <= 3; j++ {
	// 		fmt.Printf("i is %d and j is %d\n", i, j)
	// 	}
	// }

	// // Loop control statements
	// // break statement
	// for i := 1; i <= 5; i++ {
	// 	if i == 3 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// // continue statement
	// for i := 1; i <= 5; i++ {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
}
