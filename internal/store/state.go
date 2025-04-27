package store

import "github.com/shalldie/leek/internal/stock"

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
	// 更新黄金
	for _, item := range s.Golds {
		item.Update()
	}
}

// 全局 state
var State = &storeState{
	Golds: []*stock.Stock{
		stock.NewAU9999(),
		stock.NewAU0(),
		stock.NewXAU(),
		stock.NewGC(),
	},
}

// func New() *storeState {
// 	return &storeState{
// 		Golds: []*stock.Stock{
// 			stock.NewAU9999(),
// 			stock.NewAU0(),
// 			stock.NewXAU(),
// 			stock.NewGC(),
// 		},
// 	}
// }
