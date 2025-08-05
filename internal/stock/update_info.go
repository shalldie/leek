package stock

import (
	"bytes"
	"io"
	"regexp"
	"strings"

	"github.com/shalldie/leek/internal/utils"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// 判断是哪个交易所
func shorsz(code string) string {
	// 沪市主板	600/601	600519	大盘蓝筹股
	// 深市主板	000/001	000001	传统行业龙头股
	marketMap := map[string]string{
		"sh": "^(60|68|688|900|51|50|52)",
		"sz": "^(00|30|20|15|16)",
	}

	for key, pattern := range marketMap {
		if m, _ := regexp.MatchString(pattern, code); m {
			return key
		}
	}
	return ""
}

func GetInfoFromSina(codes []string) map[string]*Stock {
	// curl -H 'Referer: https://finance.sina.com.cn' 'https://hq.sinajs.cn?list=sh512820' | iconv -f GB2312 -t UTF-8
	var codeList []string

	for _, code := range codes {
		codeList = append(codeList, shorsz(code)+code)
	}

	body := utils.Fetch("https://hq.sinajs.cn?list="+strings.Join(codeList, ","), &utils.FetchOptions{
		// Query: map[string]string{
		// 	"list": strings.Join(codeList, ","),
		// },
		Headers: map[string]string{
			"Referer": "https://finance.sina.com.cn",
		},
	})

	decoder := simplifiedchinese.GBK.NewDecoder()
	render := transform.NewReader(bytes.NewReader(body), decoder)
	utf8bytes, _ := io.ReadAll(render)
	content := string(utf8bytes)

	// lines := strings.Split(content, ";\n")

	result := map[string]*Stock{}

	for _, code := range codes {
		re := regexp.MustCompile(code + `="(.*?)"`)
		submatches := re.FindStringSubmatch(content)

		if len(submatches) <= 0 {
			continue
		}

		params := strings.Split(submatches[1], ",")

		result[code] = &Stock{
			Code:     code,
			Name:     params[0],
			Price:    params[3],
			PrePrice: params[2],
		}
	}

	return result
}
