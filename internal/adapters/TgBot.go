package adapters

import (
	"clean-utility/internal/entity"
	"clean-utility/internal/interfaces"
	// "fmt"

	"github.com/yanzay/tbot/v2"
)

type TgBot struct {
	Token string
	Client *tbot.Client
}

func NewTgBot(token string) interfaces.Notifications {
	bot := tbot.New(token)
	client := bot.Client()

	return TgBot{token, client}
}

func (f TgBot) SendMessage(msg entity.Message) error {
	f.Client.SendMessage(msg.To, msg.Text)
	return nil
}

func (f TgBot) NewMessage() entity.Message {
	return entity.Message{}
}
