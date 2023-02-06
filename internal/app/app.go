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
	Logger              interfaces.Logger
}

func NewApplication(cfg config.Config, fs interfaces.FS, notifications interfaces.Notifications, lg interfaces.Logger) (Application, error) {
	app := Application{
		NotificationService: notifications,
		FSService:           fs,
		MaxVolume:           cfg.MaxVolume,
		Folders:             cfg.Folders,
		To:                  cfg.Telegram.Channel,
		Logger:              lg,
	}
	return app, nil
}

func (a Application) Run() error {
	infoBefore, err := a.FSService.DiskInfo()
	if err != nil {
		a.Logger.Error("Ошибка во время сбора информации об объёме диска до очистки")
		return err
	}
	a.Logger.Info("Собрана информация об объёме диска до очистки")
	msg := a.NotificationService.NewMessage()
	msg.To = a.To
	if a.MaxVolume < infoBefore.Used {
		a.Logger.Info("Объём в конфиге меньше, чем занятно на диске, очистка каталогов...")
		a.Logger.Info("Начало очистки каталогов")
		logs := a.FSService.ClearedFolders(a.Folders)
		a.Logger.Info("Очистка каталогов завершена")

		infoAfter, err := a.FSService.DiskInfo()
		if err != nil {
			a.Logger.Error("Ошибка во время сбора информации об объёме диска после очистки")
			return err
		}
		a.Logger.Info("Собрана информация об объёме диска после очистки, сбор логов для отправки в ТГ")

		logs.Info = append(logs.Info, fmt.Sprintf(entity.TxtAfterClean, infoBefore.Used, infoAfter.Used))
		msg.Text = strings.Join(logs.Errors, "-")
		msg.Text += strings.Join(logs.Info, "-")
		a.Logger.Info("Логи для отправки в ТГ собраны")
	} else {
		a.Logger.Info("Очистка каталогов не производилась")
		msg.Text = fmt.Sprintf(entity.TxtNotClean, infoBefore.Used)
	}

	a.Logger.Info("Отправка логов в ТГ")
	err = a.NotificationService.SendMessage(msg)
	if err != nil {
		a.Logger.Error("Ошибка при отправке логов в ТГ")
		return err
	}
	a.Logger.Info("Логи отправлены в ТГ")

	return nil
}
