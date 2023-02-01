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
	const MAX_SIZE = 4096
	var err error = nil
	last := 0

	for x := 0; x <= len(msg.Text); x += MAX_SIZE {
		if x+MAX_SIZE < len(msg.Text) {
			_, err = f.Client.SendMessage(msg.To, msg.Text[x:x+MAX_SIZE])
		} else {
			last = x
			break
		}
	}
	_, err = f.Client.SendMessage(msg.To, msg.Text[last:])

	return err
}

func (f TgBot) NewMessage() entity.Message {
	return entity.Message{}
}
