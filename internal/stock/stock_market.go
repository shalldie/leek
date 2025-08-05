package stock

type StockMarket struct {
	Name   string
	Stocks []*Stock
}

func NewStockMarket(name string, stocks []*Stock) *StockMarket {
	return &StockMarket{
		Name:   name,
		Stocks: stocks,
	}
}
