package app

import (
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

func NewAppication(cfg config.Config, fs interfaces.FS, notifications interfaces.Notifications) (Application, error) {
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
		entity.ErrorLogger.Println("Ошибка во время сбора информации об объёме диска до очистки")
		return err
	}
	entity.InfoLogger.Println("Собрана информация об объёме диска до очистки")
	msg := a.NotificationService.NewMessage()
	msg.To = a.To
	if a.MaxVolume < infoBefore.Used {
		entity.InfoLogger.Println("Объём в конфиге меньше, чем занятно на диске, очистка каталогов...")
		logs := a.FSService.ClearedFolders(a.Folders)

		infoAfter, err := a.FSService.DiskInfo()
		if err != nil {
			entity.ErrorLogger.Println("Ошибка во время сбора информации об объёме диска после очистки")
			return err
		}
		entity.InfoLogger.Println("Собрана информация об объёме диска после очистки, сбор логов для отправки в ТГ")

		logs.Info = append(logs.Info, fmt.Sprintf(entity.TxtAfterClean, infoBefore.Used, infoAfter.Used))
		msg.Text = strings.Join(logs.Errors, "-")
		msg.Text += strings.Join(logs.Info, "-")
		entity.InfoLogger.Println("Логи для отправки в ТГ собраны")
	} else {
		entity.InfoLogger.Println("Очистка каталогов не производилась")
		msg.Text = fmt.Sprintf(entity.TxtNotClean, infoBefore.Used)
	}

	entity.InfoLogger.Println("Отправка логов в ТГ")
	err = a.NotificationService.SendMessage(msg)
	if err != nil {
		entity.ErrorLogger.Println("Ошибка при отправке логов в ТГ")
		return err
	}
	entity.InfoLogger.Println("Логи отправлены в ТГ")

	return nil
}
