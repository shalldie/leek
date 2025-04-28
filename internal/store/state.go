package store

import (
	"sync"

	"github.com/shalldie/leek/internal/stock"
)

type storeState struct {
	Golds []*stock.Stock
}

// 重置数据
func (s *storeState) Reset() {
	for _, item := range s.Golds {
		item.Reset()
	}
}

// 更新所有数据
func (s *storeState) Update() {
	var wg sync.WaitGroup
	// 更新黄金
	for _, item := range s.Golds {
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
	Golds: []*stock.Stock{
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
		{
			Name:     "纳斯达克100",
			UpdateFn: stock.CreateUpdateFromEast("100.NDX100"),
		},
	},
}
