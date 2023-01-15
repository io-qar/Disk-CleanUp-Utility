package main

import (
	"clean-utility/internal/app"
	"clean-utility/internal/config"
	"flag"
	"io/ioutil"

	"log"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "../config.json", "Path to a config file")
	flag.Parse()

	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Println(err)
	}

	cfg := config.NewConfig(content)
	app, err := app.NewAppication(cfg)
	if err != nil {
		log.Fatalf("Ошибка при создании приложения: %v", err)
	}
	if err := app.Run(); err != nil {
		log.Fatalf("Ошибка выполнения: %v", err)
	}
}
