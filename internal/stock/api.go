package stock

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/shalldie/leek/internal/utils"
)

func CreateUpdateFromEast(secid string) func() (float64, float64) {
	return func() (float64, float64) {
		body := utils.Fetch("https://push2.eastmoney.com/api/qt/stock/get", &utils.FetchOptions{
			Query: map[string]string{
				"fields": "f43,f169,f170",
				"secid":  secid,
			},
		})

		// 定义一个 map 来接收解析后的数据
		var result map[string]interface{}
		json.Unmarshal(body, &result)

		data := result["data"].(map[string]interface{}) // 强制转换为 map

		price := data["f43"].(float64) / 100
		rise := data["f169"].(float64) / 100
		prePrice := price - rise

		return price, prePrice
	}
}

func CreateUpdateFromSina(kw string, curIndex int32, preIndex int32) func() (float64, float64) {
	return func() (float64, float64) {
		body := utils.Fetch("https://hq.sinajs.cn?list="+kw, &utils.FetchOptions{
			Headers: map[string]string{
				"Referer": "https://finance.sina.com.cn",
			},
		})
		content := string(body)
		content = strings.Split(content, "=\"")[1]
		arr := strings.Split(content, ",")
		cur, _ := strconv.ParseFloat(arr[curIndex], 64)
		pre, _ := strconv.ParseFloat(arr[preIndex], 64)
		return cur, pre
	}
}
