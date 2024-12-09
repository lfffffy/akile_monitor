package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	AuthSecret  string `json:"auth_secret"`
	Listen      string `json:"listen"`
	EnableWeCom bool   `json:"enable_weCom"`
	WeComKey    string `json:"weCom_key"`
	UpdateUri   string `json:"update_uri"`
	WebUri      string `json:"web_uri"`
	HookUri     string `json:"hook_uri"`
	HookToken   string `json:"hook_token"`
}

var cfg *Config

func LoadConfig() {
	file, err := os.ReadFile("config.json")
	if err != nil {
		log.Panic(err)
	}
	cfg = &Config{}
	err = json.Unmarshal(file, cfg)
	if err != nil {
		log.Panic(err)
	}
	return
}
