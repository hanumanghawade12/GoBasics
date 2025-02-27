package main
import "fmt"

func Struct() {

  // declare a struct
  type Person struct {
    name string
    age  int
  }

  // assign value to struct while creating an instance
  person1 := Person{ "John", 25}
  fmt.Println(person1)

  // define an instance
//   var person2 Person

//   // assign value to struct variables
//   person2 = Person {
//     name: "Sara",
//     age: 29,
//   }

//   fmt.Println(person2)
  // create string
  message := "Welcome to Programiz"
    
  // use len() function to count length
  stringLength := len(message)

  fmt.Println("Length of a string is:", len(stringLength))
}