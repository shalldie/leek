package stock

import (
	"encoding/json"

	"github.com/shalldie/leek/internal/utils"
)

// 获取 「AU9999」
func NewAU9999() *Stock {
	return &Stock{
		Name: "AU9999",
		UpdateFn: func() (float64, float64) {
			body := utils.Fetch("https://push2.eastmoney.com/api/qt/stock/get?fields=f43%2Cf169%2Cf170&secid=118.AU9999", nil)
			// content := string(body)

			// 定义一个 map 来接收解析后的数据
			var result map[string]interface{}
			json.Unmarshal(body, &result)

			data := result["data"].(map[string]interface{}) // 强制转换为 map

			price := data["f43"].(float64) / 100
			rise := data["f169"].(float64) / 100
			prePrice := price - rise

			return price, prePrice
		},
	}
}
