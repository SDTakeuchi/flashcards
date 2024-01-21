package config

import (
	"fmt"
	"log"
	"os"

	"github.com/go-yaml/yaml"
)

type Config struct {
	ENV             string `env:"ENV" yaml:"env"`
	SheetID         string `env:"SHEET_ID" yaml:"sheet_id"`
	SheetURL        string
	SheetCredential string `env:"SHEET_CREDENTIAL" yaml:"sheet_credential"`
	TabName         string `env:"TAB_NAME" yaml:"tab_name"`
}

var globalConfig Config

func Load() {
	f, err := os.Open("adapter/config/config.yaml")
	if err != nil {
		log.Fatalf(
			"load config failed: os.Open err: %v",
			err.Error(),
		)
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(&globalConfig)
	if err != nil {
		log.Fatalf(
			"err: %v, globalConfig: %+v",
			err.Error(),
			globalConfig,
		)
	}

	globalConfig.SheetURL = fmt.Sprintf("https://docs.google.com/spreadsheets/d/%s/edit", globalConfig.SheetID)
}

func Get() Config {
	return globalConfig
}
