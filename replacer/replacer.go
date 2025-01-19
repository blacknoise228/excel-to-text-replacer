package replacer

import (
	"excelkek/config"
	"fmt"
	"os"
	"strings"
)

func Generate(config config.Config, data map[int]map[string]string, colCount int) error {
	var result strings.Builder

	template, err := os.ReadFile(config.TextFile)
	if err != nil {
		return err
	}

	for _, m := range data {
		text := string(template)
		for i := 1; i <= colCount; i++ {

			v := fmt.Sprintf("$%d", i)
			d := m[v]

			text = strings.Replace(text, v, d, -1)
		}
		result.WriteString(text + "\n")
	}

	err = os.WriteFile(config.OutputFile, []byte(result.String()), 0644)
	if err != nil {
		return err
	}
	return nil
}
