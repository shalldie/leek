package store

import (
	"sync"

	"github.com/shalldie/leek/internal/stock"
)

type storeState struct {
	MarketIndex int
	Markets     []*stock.StockMarket
}

// 当前分类下所有 Stock
func (s *storeState) Stocks() []*stock.Stock {
	return s.Markets[s.MarketIndex].Stocks
}

// 重置数据
func (s *storeState) Reset() {
	for _, market := range s.Markets {
		for _, s := range market.Stocks {
			s.Reset()
		}
	}
}

// 更新当前分类所有数据
func (s *storeState) Update() {
	var wg sync.WaitGroup

	for _, item := range s.Stocks() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			item.Update()
		}()
	}
	wg.Wait()
}

// 全局 state
var State = &storeState{
	MarketIndex: 0,
	Markets: []*stock.StockMarket{
		{
			Name: "黄金",
			Stocks: []*stock.Stock{
				{
					Name:     "AU9999",
					UpdateFn: stock.CreateUpdateFromEast("118.AU9999"),
				},
				{
					Name:     "沪金主连",
					UpdateFn: stock.CreateUpdateFromSina("nf_AU0", 6, 10),
				},
				{
					Name:     "伦敦金",
					UpdateFn: stock.CreateUpdateFromSina("hf_XAU", 0, 7),
				},
				{
					Name:     "纽约金",
					UpdateFn: stock.CreateUpdateFromSina("hf_GC", 0, 7),
				},
			},
		},
		{
			Name: "A股", // sz:0 sh:1
			Stocks: []*stock.Stock{
				{
					// Name: "银行ETF龙头",
					Code: "512820",
					// UpdateFn: stock.CreateUpdateFromEast("1.512820"),
					// UpdateFn: stock.CreateUpdateFromSina("sh512820", 3, 2),
					// UpdateFn: func() *stock.Stock {
					// 	return stock.GetInfoFromSina([]string{"512820"})["512820"]
					// },
				},
				{
					// Name: "王子新材",
					Code: "002735",
					// UpdateFn: func() *stock.Stock {
					// 	return stock.GetInfoFromSina([]string{"002735"})["002735"]
					// },
				},
				// {
				// 	Name:     "王子新材",
				// 	UpdateFn: stock.CreateUpdateFromEast("0.002735"),
				// },
				{
					Code: "000565",
					// Name:     "三峡",
					// UpdateFn: stock.CreateUpdateFromEast("0.000565"),
					// UpdateFn: func() *stock.Stock {
					// 	return stock.GetInfoFromSina([]string{"000565"})["000565"]
					// },
				},
			},
		},
		{
			Name: "美股",
			Stocks: []*stock.Stock{
				{
					Name:     "纳斯达克100",
					UpdateFn: stock.CreateUpdateFromEast("100.NDX100"),
				},
			},
		},
	},
}
