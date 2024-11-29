package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ExcelFile  string `yaml:"excelfile"`
	PageName   string `yaml:"pagename"`
	TextFile   string `yaml:"textfile"`
	OutputFile string `yaml:"outputfile"`
}

func NewConfig() Config {
	var configs Config

	viper.SetConfigFile("./config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := viper.Unmarshal(&configs); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Loaded settings: %+v\n", configs)

	return configs
}
