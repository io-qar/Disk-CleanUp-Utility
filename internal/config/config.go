package config

import (
	"clean-utility/internal/adapters"

	"github.com/tidwall/gjson"
)

const defaultVolume uint64 = 50

type Config struct {
	Folders   []string
	MaxVolume uint64
	Telegram  struct {
		Channel  string
		BotToken string
	}
}

func NewConfig(jsonConfig []byte) Config {
	folders := gjson.Get(string(jsonConfig), "folders").Array()
	var configFolders []string
	logger := adapters.NewLogger()
	logger.Info("Чтение папок")

	for _, folder := range folders {
		configFolders = append(configFolders, folder.Str)
		logger.Info("Папка %s добавлена в очередь очистки", folder.Str)
	}

	v := Config{
		configFolders,
		gjson.Get(string(jsonConfig), "maxVolume").Uint(),
		struct{Channel string; BotToken string} {
			gjson.Get(string(jsonConfig), "telegram-bot.channel").Str,
			gjson.Get(string(jsonConfig), "telegram-bot.token").Str,
		},
	}

	return v
}
