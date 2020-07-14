package methods

// ProductPrice define product with value
type ProductPrice struct {
	Name        string
	Description string
	CostValue   float32
}

//CalculateSalesPrice calculate sales price from product cost value
func (productPrice ProductPrice) CalculateSalesPrice(percentageOfSale float32) float32 {
	return ((productPrice.CostValue * percentageOfSale) / 100) + productPrice.CostValue
}
