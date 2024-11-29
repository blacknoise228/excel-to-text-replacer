package replacer

import (
	"excelkek/config"
	"excelkek/excel"
	"os"
	"strconv"
	"strings"
)

func Generate(config config.Config, data []excel.Schema) error {
	var result strings.Builder

	template, err := os.ReadFile(config.TextFile)
	if err != nil {
		return err
	}

	for _, d := range data {
		age := strconv.Itoa(d.Age)

		text := string(template)

		text = strings.Replace(text, "?", d.Name, 1)
		text = strings.Replace(text, "?", age, 1)
		text = strings.Replace(text, "?", d.Email, 1)
		text = strings.Replace(text, "?", d.Country, 1)

		result.WriteString(text + "\n")

	}

	err = os.WriteFile(config.OutputFile, []byte(result.String()), 0644)
	if err != nil {
		return err
	}
	return nil
}
