package stock

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/shalldie/leek/internal/utils"
)

// # 汇总文档
// # https://www.joinquant.com/community/post/detailMobile?postId=30016&page=&limit=20&replyId=&tag=

// # AU9999，黄金现货，东方财富，https://quote.eastmoney.com/q/118.AU9999.html
// curl 'https://push2.eastmoney.com/api/qt/stock/get?fields=f43%2Cf169%2Cf170&secid=118.AU9999'

// # 沪金连续
// curl -H 'Referer: https://finance.sina.com.cn' 'https://hq.sinajs.cn?list=nf_AU0'

var (
	COLOR_RED   = lipgloss.Color("#f00")
	COLOR_GREEN = lipgloss.Color("#090")
)

func f2str(num float64, maxDec int32) string {
	numStr := fmt.Sprintf("%d", maxDec)
	result := fmt.Sprintf("%."+numStr+"f", num)
	return result
	// 去掉末尾的零
	// result = strings.TrimRight(result, "0")

	// // 如果小数点后没有数字，去掉小数点
	// result = strings.TrimRight(result, ".")
	// return result
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
	s.Price = "-"
	s.PrePrice = "-"
	s.Rise = "-"
	s.Rate = "-"
}

func (s *Stock) Update() {
	err := utils.Try(func() {
		price, prePrice := s.UpdateFn()
		rise := price - prePrice
		rate := rise / prePrice

		s.Price = f2str(price, 3)
		s.PrePrice = f2str(prePrice, 3)
		s.Rise = f2str(rise, 3)
		s.Rate = f2str(rate*100, 2) + "%"

		// 着色，添加 「+」
		if rise >= 0 {
			s.Rise = lipgloss.NewStyle().Foreground(COLOR_RED).Render("+" + s.Rise)
			s.Rate = lipgloss.NewStyle().Foreground(COLOR_RED).Render("+" + s.Rate)
		} else {
			s.Rise = lipgloss.NewStyle().Foreground(COLOR_GREEN).Render(s.Rise)
			s.Rate = lipgloss.NewStyle().Foreground(COLOR_GREEN).Render(s.Rate)
		}
	})

	if err != nil {
		s.Reset()
	}
}
