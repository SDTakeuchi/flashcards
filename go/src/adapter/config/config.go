package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-yaml/yaml"
)

type (
	Config struct {
		ENV             string `env:"ENV" yaml:"env"`
		Debug           bool   `yaml:"debug" default:"false"`
		SheetID         string `env:"SHEET_ID" yaml:"sheet_id"`
		SheetURL        string
		SheetCredential string `env:"SHEET_CREDENTIAL" yaml:"sheet_credential"`
		TabName         string `env:"TAB_NAME" yaml:"tab_name"`
		DB              DB     `yaml:"db"`
		Token           Token  `yaml:"token"`
	}

	DB struct {
		Database    string `required:"true" yaml:"name"`
		UserName    string `required:"true" yaml:"user_name"`
		Password    string `required:"true" yaml:"password"`
		Port        string `required:"true" yaml:"port" default:"5432"`
		HostName    string `required:"true" yaml:"host_name"`
		SSLMode     string `required:"true" yaml:"ssl_mode"`
		ConnTimeout int    `required:"true" yaml:"conn_timeout" default:"9"`
		// 0 means no termination
		MaxOpen int `required:"true" yaml:"max_open" default:"0"`
		MaxIdle int `required:"true" yaml:"max_idle" default:"2"`
		// 0 means no termination as long as the process is running
		// set values that can be read by time.ParseDuration()
		MaxLifeTime time.Duration `required:"true" yaml:"max_life_time" default:"0"`
	}

	Token struct {
		SecretKey            string        `required:"true" yaml:"secret_key"`
		AccessTokenDuration  time.Duration `required:"true" yaml:"access_token_duration" default:"10m"`
		RefreshTokenDuration time.Duration `required:"true" yaml:"refresh_token_duration" default:"24h"`
		MinSecretKeySize     int           `required:"true" yaml:"min_secret_key_size"`
	}
)

var globalConfig Config

func Load(path string) {
	f, err := os.Open(path)
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
