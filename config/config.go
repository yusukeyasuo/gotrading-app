package config

import (
	"log"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	ApiKey string
	ApiSecret string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
	}

	Config = ConfigList{
		ApiKey: cfg.Section("bitflyer").Key("api_key").String(),
		ApiSecret: cfg.Section("bitflyer").Key("api_secret").String(),
	}
}