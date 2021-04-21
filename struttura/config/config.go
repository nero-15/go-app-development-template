package config

import (
	"log"
	"os"

	ini "gopkg.in/ini.v1"
)

// ConfigList is api key struct
type ConfigList struct {
	DbName    string
	SQLDriver string
	Port      string
}

// Config is ConfigList
var Config ConfigList

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		DbName:    cfg.Section("db").Key("name").String(),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		Port:      cfg.Section("web").Key("port").String(),
	}
}