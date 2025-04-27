package main

import (
	"fmt"
	"strconv"

	"github.com/shalldie/leek/internal/app"
)

func main() {
	app.Run()
	// au9 := stock.NewAU9999()
	// au0 := stock.NewAU0()
	// xau := stock.NewXAU()
	// gc := stock.NewGC()

	// println(au9.String())
	// println(au0.String())
	// println(xau.String())
	// println(gc.String())
}

func main2() {
	num := 123.4
	str := fmt.Sprintf("%.3f", num) // 保留最多3位小数
	fmt.Println(str)
}

func formatFloat(num float64) string {
	// 先用 %.3f 保留最多 3 位小数
	str := fmt.Sprintf("%.3f", num)
	// 去掉末尾的零
	str = strconv.FormatFloat(num, 'f', -1, 64)
	return str
}

func main5() {
	num := 123.45000
	fmt.Println(formatFloat(num)) // 输出 123.45

	num2 := 123.0
	fmt.Println(formatFloat(num2)) // 输出 123

	num3 := 123.456789
	fmt.Println(formatFloat(num3)) // 输出 123.457
}
