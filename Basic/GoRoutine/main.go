package main

import (
	"fmt"
	"time"
)

func main() {
	dones := make([]chan bool, 3)
	// done := make(chan bool)
	dones[0] = make(chan bool)
	go greet("1st line", dones[0])
	dones[1] = make(chan bool)
	go greet("2nd line", dones[1])
	dones[2] = make(chan bool)
	go slowGreet("3rd line", dones[2])
	// <-done
	// <-done
	// <-done
	for _, done := range dones {
		<-done
	}

	// greet("4 th line")
}

func greet(title string, done chan bool) {
	fmt.Println("Hello, World!", title)
	done <- true
}

func slowGreet(title string, done chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello, World!", title)
	done <- true
}
