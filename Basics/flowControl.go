package main

import "fmt"

func flowControl() {
	// if num > 0{
	// 	fmt.Println("Number is positive")
	// } else if num == 0{
	// 	fmt.Println("Number is Zero")
	// }else{
	// 	fmt.Println("Number is negative")
	// }

	dayOfWeek := 3

	switch dayOfWeek {

	case 1:
		fmt.Println("Sunday")

	case 2:
		fmt.Println("Monday")

	case 3:
		fmt.Println("Tuesday")

	case 4:
		fmt.Println("Wednesday")

	case 5:
		fmt.Println("Thursday")

	case 6:
		fmt.Println("Friday")

	case 7:
		fmt.Println("Saturday")

	default:
		fmt.Println("Invalid day")
	}
}
