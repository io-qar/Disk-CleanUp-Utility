package app

import (
	"clean-utility/internal/adapters"
	"clean-utility/internal/config"
	"clean-utility/internal/entity"
	"clean-utility/internal/interfaces"
	"fmt"

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
	notifications := adapters.NewTgBot(cfg.Telegram.BotToken)
	fs := adapters.NewFS()

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
	infoBefore, err := a.FSService.DiskInfo()
	if err != nil {
		return err
	}
	msg := a.NotificationService.NewMessage()
	msg.To = a.To
	if a.MaxVolume < infoBefore.Used {
		logs := a.FSService.ClearedFolders(a.Folders)
		
		infoAfter, err := a.FSService.DiskInfo()
		if err != nil {
			return err
		}

		// проводитм анализ логс на наличие ошибок, чтоб определить как отправить сообщение
		// собираем все сообщение
		logs.Info = append(logs.Info, fmt.Sprintf(entity.TxtAfterClean, infoAfter.Used, infoBefore.Used))
		msg.Text = strings.Join(logs.Errors, "-")
		msg.Text += strings.Join(logs.Info, "-")
	} else {
		// тут решаем из настроек надо ли информировать
		// о том что надо отправлять сообщение если ничего не делалось
		msg.Text = fmt.Sprintf(entity.TxtNotClean, infoBefore.Used)
		// если надо
		//  тогда заполняем msg нужной инфой
	}

	err = a.NotificationService.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}