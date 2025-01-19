package excel

import (
	"excelkek/config"
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func ReadFromFile(config config.Config) (data map[int]map[string]string, colCount int) {
	ex, err := excelize.OpenFile(config.ExcelFile)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := ex.Close(); err != nil {
			panic(err)
		}
	}()

	rows, err := ex.GetRows(config.PageName)
	if err != nil {
		panic(err)
	}

	columns, err := ex.GetCols(config.PageName)
	if err != nil {
		panic(err)
	}

	data = make(map[int]map[string]string)

	for i, row := range rows {
		if i == 0 {
			continue
		}
		if _, ok := data[i]; !ok {
			data[i] = make(map[string]string)
		}
		for j, colCell := range row {
			coLname := fmt.Sprintf("$%d", j+1)
			data[i][coLname] = colCell
			log.Println(coLname, row)
		}

	}
	return data, len(columns)
}
