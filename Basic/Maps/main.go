package main

import (
	"fmt"
)

type flatMap map[string]float64

func main() {
	usernames := make([]string, 2, 5)
	usernames = append(usernames, "john")
	usernames = append(usernames, "jane")
	usernames = append(usernames, "james")
	usernames = append(usernames, "jake")
	usernames = append(usernames, "joe")
	usernames = append(usernames, "jim")
	usernames = append(usernames, "jill")
	usernames = append(usernames, "jess")
	usernames = append(usernames, "jerry")
	fmt.Println(usernames)

	courseRating := make(flatMap, 3)
	courseRating["Go Fundamentals"] = 4.2
	courseRating["Docker Deep Dive"] = 4.0
	courseRating["Kubernetes for Developers"] = 4.6
	courseRating["Microservices with Java"] = 4.7
	courseRating["Microservices with Go"] = 4.8
	fmt.Println(courseRating)

	for index, value := range usernames {
		fmt.Println(index, value)
	}
	for key, value := range courseRating {
		fmt.Println(key, value)
	}

}

// func main() {
// 	websites := map[string]string{
// 		"SitePoint": "https://www.sitepoint.com",
// 		"Google":    "https://www.google.com",
// 		"Facebook":  "https://www.facebook.com",
// 	}
// 	fmt.Println(websites)
// 	fmt.Println(websites["Google"])
// 	websites["linkedIn"] = "https://www.linkedIn.co.uk"
// 	fmt.Println(websites)
// 	delete(websites, "SitePoint")
// 	fmt.Println(websites)
// }
