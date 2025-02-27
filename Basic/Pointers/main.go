package main

import "fmt"

func main() {
	fmt.Println("Welcome to pinters")
	age := 32
	fmt.Println("Age before", age)
	updateAge(&age)
	adultAge := getAdultAge(age)
	fmt.Println("Age after", adultAge)
}

func getAdultAge(age int) int {
	return age - 18
}

func updateAge(age *int) {
	*age = 50
}
