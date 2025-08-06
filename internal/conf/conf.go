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
	Codes []string `json:"codes"`
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
		return
	}

	// 解析JSON
	var config ConfigModel
	err = json.Unmarshal(content, &config)
	if err != nil {
		return
	}

	// 无配置
	if len(config.Codes) <= 0 {
		return
	}

	// 使用配置
	market := &stock.StockMarket{
		Name:   "自选",
		Stocks: []*stock.Stock{},
	}
	for _, code := range config.Codes {
		market.Stocks = append(market.Stocks, &stock.Stock{Code: code})
	}

	store.State.Markets = append([]*stock.StockMarket{market}, store.State.Markets...)
}
