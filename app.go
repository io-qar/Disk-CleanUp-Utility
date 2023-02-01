package main

import (
	// "bytes"
	"clean-utility/internal/adapters"
	"clean-utility/internal/app"
	"clean-utility/internal/config"
	"clean-utility/internal/entity"
	"flag"
	"fmt"
	"io/ioutil"

	"log"
)

func init() {
	entity.InfoLogger = log.New(&entity.Buf, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	entity.WarningLogger = log.New(&entity.Buf, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	entity.ErrorLogger = log.New(&entity.Buf, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	entity.InfoLogger.Println("Запуск утилиты")
	defer fmt.Println(&entity.Buf)
	var configPath string

	flag.StringVar(&configPath, "config", "config.json", "Path to a config file")
	flag.Parse()

	content, err := ioutil.ReadFile(configPath)
	entity.InfoLogger.Println("Чтение файла конфигурации", configPath)
	if err != nil {
		entity.ErrorLogger.Println("Ошибка при чтении файла конфигурации", configPath)
		log.Fatalf(entity.CreationError, err)
	}

	cfg := config.NewConfig(content)
	fs := adapters.NewFS()
	notifications := adapters.NewTgBot(cfg.Telegram.BotToken)
	app, err := app.NewAppication(cfg, fs, notifications)
	entity.InfoLogger.Println("Инициализация бота с токеном", cfg.Telegram.BotToken)
	if err != nil {
		entity.ErrorLogger.Println(entity.CreationError)
		log.Fatalf(entity.CreationError, err)
	}
	if err := app.Run(); err != nil {
		log.Fatalf(entity.RunError, err)
	}
	entity.InfoLogger.Println("Завершение работы утилиты")
}
