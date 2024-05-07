package config

import (
	"fmt"
	"user-service/lib/pkg/utils"

	"gopkg.in/gcfg.v1"
)

func Setup() (*Config, error) {

	if !utils.IsFileOrDirExist("config.toml") {
		fmt.Println("error config not found")
	}

	cfgFile := "config.toml"

	cfg := &Config{}

	err := gcfg.ReadFileInto(cfg, cfgFile)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
