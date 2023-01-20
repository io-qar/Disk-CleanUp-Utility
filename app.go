package main

import (
	"clean-utility/internal/adapters"
	"clean-utility/internal/app"
	"clean-utility/internal/config"
	"clean-utility/internal/entity"
	"flag"
	"io/ioutil"

	"log"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "config.json", "Path to a config file")
	flag.Parse()

	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf(entity.CreationError, err)
	}

	cfg := config.NewConfig(content)
	fs := adapters.NewFS()
	notifications := adapters.NewTgBot(cfg.Telegram.BotToken)
	app, err := app.NewAppication(cfg, fs, notifications)
	if err != nil {
		log.Fatalf(entity.CreationError, err)
	}
	if err := app.Run(); err != nil {
		log.Fatalf(entity.RunError, err)
	}
}
