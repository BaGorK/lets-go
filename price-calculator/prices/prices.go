package prices

import (
	"fmt"

	"github.com/BaGorK/lets-go/price-calc/conversion"
	"github.com/BaGorK/lets-go/price-calc/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println("Error parsing price:", err)
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", price*(1+job.TaxRate))
	}

	job.TaxIncludedPrices = result

	return job.IOManager.WriteJSON(job)
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		InputPrices: []float64{},
		TaxRate:     taxRate,
	}
}
