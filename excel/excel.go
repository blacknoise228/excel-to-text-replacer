package excel

import (
	"excelkek/config"
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func ReadFromFile(config config.Config) []Schema {
	ex, err := excelize.OpenFile(config.ExcelFile)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := ex.Close(); err != nil {
			panic(err)
		}
	}()

	data := make([]Schema, 0)

	rows, err := ex.GetRows(config.PageName)
	if err != nil {
		panic(err)
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}
		id, err := strconv.Atoi(row[0])
		if err != nil {
			fmt.Println(err)
			return nil
		}
		age, err := strconv.Atoi(row[2])
		if err != nil {
			fmt.Println(err)
			return nil
		}
		data = append(data, Schema{
			ID:      id,
			Name:    row[1],
			Age:     age,
			Email:   row[3],
			Country: row[4],
		})
	}
	return data
}
