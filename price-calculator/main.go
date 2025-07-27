package main

import (
	"fmt"

	"github.com/BaGorK/lets-go/price-calc/filemanager"
	"github.com/BaGorK/lets-go/price-calc/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.10, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.Filemanager{
			InputFilePath:  "prices.txt",
			OutputFilePath: fmt.Sprintf("result_%.0f.json", taxRate*100),
		}
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Printf("Could not process prices with tax rate %.0f%%: %v\n", taxRate*100, err)
		}
	}
}
