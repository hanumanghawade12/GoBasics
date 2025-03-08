package main

import (
	"fmt"
	"price-calculator/filemanager"
	"price-calculator/prices"
)

func main() {
	taxRate := []float64{0.1, 0.2, 0.3, 0.4, 0.5}
	dones := make([]chan bool, len(taxRate))

	for index, taxValue := range taxRate {
		dones[index] = make(chan bool)
		fm := filemanager.NewFileManager("input.json", fmt.Sprintf("output_%v.json", taxValue*100))
		pricesJob := prices.NewTaxIncludedPriceJob(fm, taxValue)
		go pricesJob.Process(dones[index])
	}
	for _, done := range dones {
		<-done
	}
}
