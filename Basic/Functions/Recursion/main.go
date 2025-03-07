package main

import "fmt"

func main() {
	// fact := factorial(5)
	// fmt.Println(fact)
	sum := sumNumbers(1, 2, 3, 4, 5)
	fmt.Println(sum)
}

// func factorial(n int) int {
// 	if n == 0 {
// 		return 1
// 	}
// 	return n * factorial(n-1)
// }

func sumNumbers(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}
