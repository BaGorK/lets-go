package prices

import (
	"fmt"

	"github.com/BaGorK/lets-go/price-calc/conversion"
	"github.com/BaGorK/lets-go/price-calc/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := filemanager.ReadLines("prices.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println("Error parsing price:", err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", price*(1+job.TaxRate))
	}

	fmt.Printf("Tax rate %.2f: %v\n", job.TaxRate, result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{},
		TaxRate:     taxRate,
	}
}
