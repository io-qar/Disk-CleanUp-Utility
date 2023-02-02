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
	logger := adapters.NewLogger()
	logger.Info("Запуск утилиты")
	var configPath string

	flag.StringVar(&configPath, "config", "config.json", "Path to a config file")
	flag.Parse()
	content, err := ioutil.ReadFile(configPath)
	logger.Infof("Чтение файла конфигурации", configPath)
	if err != nil {
		logger.Info("Ошибка при чтении файла конфигурации")
		log.Fatalf(entity.CreationError, err)
	}
	
	cfg := config.NewConfig(content)
	logger.Info("Чтение файла конфигурации завершено")
	fs := adapters.NewFS()
	notifications := adapters.NewTgBot(cfg.Telegram.BotToken)
	app, err := app.NewAppication(cfg, fs, notifications, logger)
	logger.Infof("Инициализация бота с токеном %s", cfg.Telegram.BotToken)
	if err != nil {
		logger.Error("Ошибка при создании приложения")
		log.Fatalf(entity.CreationError, err)
	}
	if err := app.Run(); err != nil {
		log.Fatalf(entity.RunError, err)
	}
	logger.Info("Завершение работы утилиты")
}
