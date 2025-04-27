package stock

import (
	"fmt"
	"strings"
)

func f2str(num float64, maxDec int32) string {
	numStr := fmt.Sprintf("%d", maxDec)
	result := fmt.Sprintf("%."+numStr+"f", num)
	// 去掉末尾的零
	result = strings.TrimRight(result, "0")

	// 如果小数点后没有数字，去掉小数点
	result = strings.TrimRight(result, ".")
	return result
}

type Stock struct {
	// 名称
	Name string
	// 用于更新数据的 fn，返回 「当前价格、昨收价格」
	UpdateFn func() (float64, float64)

	// 当前价格
	Price string
	// 昨收价格
	PrePrice string
	// 涨跌
	Rise string
	// 涨幅
	Rate string
}

func (s *Stock) Reset() {
	s.Price = "---"
	s.PrePrice = "---"
	s.Rise = "---"
	s.Rate = "---"
}

func (s *Stock) Update() {
	price, prePrice := s.UpdateFn()
	rise := price - prePrice
	rate := rise / prePrice

	// s.Price = fmt.Sprintf("%.3f", price)
	s.Price = f2str(price, 3)
	// s.PrePrice = fmt.Sprintf("%.3f", prePrice)
	s.PrePrice = f2str(prePrice, 3)
	// s.Rise = fmt.Sprintf("%.3f", rise)
	s.Rise = f2str(rise, 3)
	// s.Rate = fmt.Sprintf("%.2f", rate*100) + "%"
	s.Rate = f2str(rate*100, 2) + "%"
}

func (s *Stock) String() string {
	result := fmt.Sprintf("%s：\n当前价格：%s，涨跌：%s，涨幅：%s", s.Name, s.Price, s.Rise, s.Rate)
	return result
}
