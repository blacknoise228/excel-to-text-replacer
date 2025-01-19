package main

import (
	"excelkek/config"
	"excelkek/excel"
	"excelkek/replacer"
	"fmt"
)

func main() {
	config := config.NewConfig()

	excel, colCount := excel.ReadFromFile(config)

	err := replacer.Generate(config, excel, colCount)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Done")
}
