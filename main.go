package main

import (
	"excelkek/config"
	"excelkek/excel"
	"excelkek/replacer"
	"fmt"
)

func main() {
	config := config.NewConfig()

	excel := excel.ReadFromFile(config)

	err := replacer.Generate(config, excel)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Done")
}
