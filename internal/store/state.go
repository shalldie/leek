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
			Name: "A股", // sz:0 sh:1
			Stocks: []*stock.Stock{
				{
					Code: "1.000001", // 上证指数
				},
				{
					Code: "0.399001", // 深证成指
				},
				{
					Code: "1.000016", // 上证50
				},
				{
					Code: "1.000300", // 沪深300
				},
				{
					Code: "1.000688", // 科创50
				},
			},
		},
		{
			Name: "黄金",
			Stocks: []*stock.Stock{
				{
					Code: "118.AU9999", // AU9999
				},
				{
					Code: "113.aum",
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
			Name: "美股",
			Stocks: []*stock.Stock{
				{
					Code: "100.DJIA", // 道琼斯
				},
				{
					Code: "100.SPX", // 标普500
				},
				{
					Code: "100.NDX", // 纳斯达克
				},
				{
					Code: "100.NDX100", // 纳斯达克100
				},
			},
		},
	},
}
