package conf

import (
	"encoding/json"
	"os"
	"os/user"
	"path/filepath"

	"github.com/shalldie/leek/internal/stock"
	"github.com/shalldie/leek/internal/store"
)

type ConfigModel struct {
	A []string `json:"a"`
}

func init() {
	usr, err := user.Current()
	if err != nil {
		return
	}

	// 构建配置文件路径
	configPath := filepath.Join(usr.HomeDir, ".leek_config.json")

	// 读取文件内容
	content, err := os.ReadFile(configPath)
	if err != nil {
		// fmt.Printf("无法读取配置文件: %v\n", err)
		return
	}

	// 解析JSON
	var config ConfigModel
	err = json.Unmarshal(content, &config)
	if err != nil {
		// fmt.Printf("无法解析配置文件: %v\n", err)
		return
	}

	// 使用配置
	// fmt.Printf("成功读取配置: %+v\n", config)
	market := &stock.StockMarket{
		Name:   "自选",
		Stocks: []*stock.Stock{},
	}
	for _, code := range config.A {
		market.Stocks = append(market.Stocks, &stock.Stock{Code: code})
	}

	store.State.Markets = append([]*stock.StockMarket{market}, store.State.Markets...)
}
