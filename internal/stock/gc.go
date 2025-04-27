package stock

import (
	"strconv"
	"strings"

	"github.com/shalldie/leek/internal/utils"
)

// 获取 「纽约金」
func NewGC() *Stock {
	return &Stock{
		Name: "纽约金",
		UpdateFn: func() (float64, float64) {
			body := utils.Fetch("https://hq.sinajs.cn?list=hf_GC", &utils.FetchOptions{
				Headers: map[string]string{
					"Referer": "https://finance.sina.com.cn",
				},
			})
			content := string(body)
			content = strings.Split(content, "=\"")[1]
			arr := strings.Split(content, ",")
			cur, _ := strconv.ParseFloat(arr[0], 64)
			pre, _ := strconv.ParseFloat(arr[7], 64)
			return cur, pre
		},
	}
}
