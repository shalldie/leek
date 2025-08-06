package stock

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"regexp"
	"strings"

	"github.com/shalldie/leek/internal/utils"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// 判断是哪个交易所
// 沪: sh, 1
// 深: sz, 0
func shorsz10(code string) string {
	// 沪市主板	600/601	600519	大盘蓝筹股
	// 深市主板	000/001	000001	传统行业龙头股
	marketMap := map[string]string{
		"1": "^(60|68|688|900|51|50|52)",
		"0": "^(00|30|20|15|16)",
	}

	for key, pattern := range marketMap {
		if m, _ := regexp.MatchString(pattern, code); m {
			return key
		}
	}
	return ""
}

func CreateUpdateFromEast(code string) func() *Stock {

	secid := code
	prefix := shorsz10(code)
	if len(prefix) > 0 {
		secid = prefix + "." + secid
	}

	return func() *Stock {
		body := utils.Fetch("https://push2.eastmoney.com/api/qt/stock/get", &utils.FetchOptions{
			Query: map[string]string{
				// f59（缩小倍数，10的n次方）、f43（最新价格）、f57（股票代码）、f58（股票名称）、f169（涨跌额）、f170（涨跌幅）
				"fields": "f59,f43,f169,f170,f58,f57",
				"secid":  secid,
			},
		})

		// 定义一个 map 来接收解析后的数据
		var result map[string]interface{}
		json.Unmarshal(body, &result)

		data := result["data"].(map[string]interface{}) // 强制转换为 map

		scale := math.Pow(10, data["f59"].(float64))

		price := data["f43"].(float64) / scale
		rise := data["f169"].(float64) / scale
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
