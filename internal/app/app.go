package app

import (
	"clean-utility/internal/adapters"
	"clean-utility/internal/config"
	"clean-utility/internal/interfaces"
	"strings"
)

type Application struct {
	NotificationService interfaces.Notifications
	FSService           interfaces.FS
	MaxVolume           uint64
	Folders             []string
	To                  string
}

func NewAppication(cfg config.Config) (Application, error) {

	notifications := adapters.NewFakeTgBot(cfg.Telegram.BotToken)
	fs := adapters.NewFaleFS()

	app := Application{
		NotificationService: notifications,
		FSService:           fs,
		MaxVolume:           cfg.MaxVolume,
		Folders:             cfg.Folders,
		To:                  cfg.Telegram.Channel,
	}
	return app, nil
}

func (a Application) Run() error {

	info, err := a.FSService.DiskInfo()
	if err != nil {
		return err
	}
	msg := a.NotificationService.NewMessage()
	msg.To = a.To
	if a.MaxVolume < info.Used {
		logs := a.FSService.ClearedFolders(a.Folders)
		// проводитм анализ логс на наличие ошибок, чтоб определить как отправить сообщение
		// собираем все сообщение

		// простой пример
		msg.Text = strings.Join(logs.Info, "-")

	} else {
		// тут решаем из настроек надо ли информировать
		// о том что надо отправлять сообщение если ничего не делалось

		// если надо
		//  тогда заполняем msg нужной инфой

	}

	err = a.NotificationService.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}
