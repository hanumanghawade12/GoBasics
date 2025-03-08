package prices

import (
	"fmt"
	"price-calculator/conversion"
	"price-calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager `json:"-"`
	InputPrices       []float64               `json:"input_prices"`
	TaxRate           float64                 `json:"tax_rate"`
	TaxIncludedPrices map[string]string       `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println("error converting strings to floats")
		return
	}
	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println("error converting strings to floats")
		return
	}
	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process(done chan bool) {
	job.LoadData()
	result := make(map[string]string)
	for _, priceValue := range job.InputPrices {
		includesPrice := priceValue * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", priceValue)] = result[fmt.Sprintf("%.2f", includesPrice)]
	}
	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)
	done <- true
}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   fm,
		InputPrices: []float64{10, 20, 30, 40, 50},
		TaxRate:     taxRate,
	}
}
