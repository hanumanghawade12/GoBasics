package main

import "fmt"

func main() {
	prices := []int{100, 200, 300, 400, 500}
	prices[4] = 600
	updatedPrice := append(prices, 700)
	discountedPrice := []int{50, 100, 150}
	newprices := append(updatedPrice, discountedPrice...)

	fmt.Println(newprices, prices)
}

// func main() {
// 	prices := [5]float64{100.12, 200.45, 300.56, 400.45, 500.54}
// 	fmt.Println(prices)
// 	// Accesing the specific element
// 	fmt.Println(prices[0])
// 	fmt.Println(prices[1])
// 	// Prodct names array
// 	products := [3]string{"Apple", "banana", "Orange"}
// 	// type of product[2] is string
// 	fmt.Println(products[2])

// 	featuresPrices := prices[1:]
// 	featuresPrices[0] = 1000.12
// 	highlightedFeatures := featuresPrices[:3]
// 	fmt.Println(highlightedFeatures, prices)
// 	// length of featuresPrices is 3
// 	// fmt.Println(len(featuresPrices))
// 	// capacity of featuresPrices is 4
// 	// fmt.Println(cap(featuresPrices))
// }
