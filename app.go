package main

import (
	"clean-utility/internal/app"
	"clean-utility/internal/config"
	"clean-utility/internal/entity"
	"flag"
	"io/ioutil"

	"log"
)

// FIX_ME:
// Rename `example.config.json` to `config.json`
// 1. нет файла example.config.json, пропал куда-то
// 2. добавить config.json в гитигнор

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "config.json", "Path to a config file")
	flag.Parse()

	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf(entity.CreationError, err)
	}

	cfg := config.NewConfig(content)
	app, err := app.NewAppication(cfg)
	if err != nil {
		log.Fatalf(entity.CreationError, err)
	}
	if err := app.Run(); err != nil {
		log.Fatalf(entity.RunError, err)
	}
}
