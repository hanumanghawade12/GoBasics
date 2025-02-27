package main

import (
	"fmt"

	"github.com/structs/user"
)

func main() {
	userFirstname := getUserData("Enter your name: ")
	userLastname := getUserData("Enter your last name: ")
	userBirthDate := getUserData("Enter your birth date: ")

	var appUser *user.User
	appUser, err := user.NewUser(userFirstname, userLastname, userBirthDate)
	if err != nil {
		fmt.Println(err)
		return
	}

	adminUser := user.NewAdmin("hanuman@gmail.com", "123456")

	adminUser.OutputUser()

	appUser.OutputUser()
	// appUser.ClearUserData()
	// appUser.OutputUser()
}

func getUserData(output string) string {
	fmt.Print(output)
	var value string
	fmt.Scanln(&value)
	return value
}
