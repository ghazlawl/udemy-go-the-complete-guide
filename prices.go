package prices

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) Process() {
	result := make(map[string]float64)
}

func NewTaxIncludedPriceJob(taxRate float64, inputPrices []float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:           taxRate,
		InputPrices:       inputPrices,
		TaxIncludedPrices: make(map[string]float64),
	}
}
