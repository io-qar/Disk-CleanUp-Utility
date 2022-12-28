package cmd

import (
	"clean-utility/internal/app"
	"clean-utility/internal/config"

	"log"
)

func main() {

	cfg := config.NewConfig()
	app, err := app.NewAppication(cfg)
	if err != nil {
		log.Fatalf("Ошибка при создании приложения: %v", err)
	}
	if err := app.Run(); err != nil {
		log.Fatalf("Ошибка выполнения: %v", err)
	}

}
