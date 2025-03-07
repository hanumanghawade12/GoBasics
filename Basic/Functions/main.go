package main

import "fmt"

type transform func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	moreNumbers := []int{6, 7, 8, 9, 10}
	doubleNumber := transformNumbers(&numbers, doubleNumber)
	tripleNumber := transformNumbers(&numbers, trippleNumber)
	// fmt.Println(doubleNumber, tripleNumber)

	transformFunction1 := getTransformFunction(&numbers)
	transformFunction2 := getTransformFunction(&moreNumbers)

	doubleNumber = transformNumbers(&numbers, transformFunction1)
	tripleNumber = transformNumbers(&moreNumbers, transformFunction2)
	fmt.Println(doubleNumber, tripleNumber)
}

func transformNumbers(numbers *[]int, transform transform) []int {
	doubleNumber := []int{}
	for _, val := range *numbers {
		doubleNumber = append(doubleNumber, transform(val))
	}
	return doubleNumber
}

func getTransformFunction(numbers *[]int) transform {
	if (*numbers)[0] == 1 {
		return doubleNumber
	} else {
		return trippleNumber

	}
}

func doubleNumber(val int) int {
	return val * 2
}

func trippleNumber(val int) int {
	return val * 3
}
