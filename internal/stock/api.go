package stock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/shalldie/leek/internal/utils"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func CreateUpdateFromEast(secid string) func() *Stock {
	return func() *Stock {
		body := utils.Fetch("https://push2.eastmoney.com/api/qt/stock/get", &utils.FetchOptions{
			Query: map[string]string{
				"fields": "f43,f169,f170,f58",
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

		return &Stock{
			Name:     data["f58"].(string),
			Price:    fmt.Sprint(price),
			PrePrice: fmt.Sprint(prePrice),
		}
	}
}

func CreateUpdateFromSina(kw string, curIndex int32, preIndex int32) func() *Stock {
	return func() *Stock {
		// curl -H 'Referer: https://finance.sina.com.cn' 'https://hq.sinajs.cn?list=sh512820' | iconv -f GB2312 -t UTF-8
		body := utils.Fetch("https://hq.sinajs.cn?list="+kw, &utils.FetchOptions{
			Headers: map[string]string{
				"Referer": "https://finance.sina.com.cn",
			},
		})

		decoder := simplifiedchinese.GBK.NewDecoder()
		render := transform.NewReader(bytes.NewReader(body), decoder)
		utf8bytes, _ := io.ReadAll(render)

		content := string(utf8bytes)
		content = strings.Split(content, "=\"")[1]
		arr := strings.Split(content, ",")
		// cur, _ := strconv.ParseFloat(arr[curIndex], 64)
		// pre, _ := strconv.ParseFloat(arr[preIndex], 64)
		return &Stock{
			Price:    arr[curIndex],
			PrePrice: arr[preIndex],
		}
	}
}
