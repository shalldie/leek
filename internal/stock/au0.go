package stock

import (
	"strconv"
	"strings"

	"github.com/shalldie/leek/internal/utils"
)

// 获取 「沪金主连」
func NewAU0() *Stock {
	return &Stock{
		Name: "沪金主连",
		UpdateFn: func() (float64, float64) {
			body := utils.Fetch("https://hq.sinajs.cn?list=nf_AU0", &utils.FetchOptions{
				Headers: map[string]string{
					"Referer": "https://finance.sina.com.cn",
				},
			})
			content := string(body)
			arr := strings.Split(content, ",")
			cur, _ := strconv.ParseFloat(arr[6], 64)
			pre, _ := strconv.ParseFloat(arr[10], 64)
			return cur, pre
		},
	}
}
