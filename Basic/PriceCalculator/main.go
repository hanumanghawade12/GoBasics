package main

import (
	"fmt"
	"price-calculator/filemanager"
	"price-calculator/prices"
)

func main() {
	taxRate := []float64{0.1, 0.2, 0.3, 0.4, 0.5}

	for _, taxValue := range taxRate {
		fm := filemanager.NewFileManager("input.json", fmt.Sprintf("output_%v.json", taxValue*100))
		pricesJob := prices.NewTaxIncludedPriceJob(fm, taxValue)
		pricesJob.Process()
	}
}
