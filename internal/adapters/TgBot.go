package adapters

import (
	"clean-utility/internal/entity"
	"clean-utility/internal/interfaces"

	"github.com/yanzay/tbot/v2"
)

type TgBot struct {
	Token  string
	Client *tbot.Client
}

func NewTgBot(token string) interfaces.Notifications {
	bot := tbot.New(token)
	client := bot.Client()

	return TgBot{token, client}
}

func (f TgBot) SendMessage(msg entity.Message) error {
	_, err := f.Client.SendMessage(msg.To, msg.Text)
	return err
}

func (f TgBot) NewMessage() entity.Message {
	return entity.Message{}
}
