package adapters

import (
	"clean-utility/internal/entity"
	"clean-utility/internal/interfaces"

	// FIX-ME:
	//  убрать закомментированый код
	// "fmt"

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
	// FIX-ME:
	// функция имеет второй возвращаемый аргумент в виде ошибки
	// его необходимо проверять
	f.Client.SendMessage(msg.To, msg.Text)
	return nil
}

func (f TgBot) NewMessage() entity.Message {
	return entity.Message{}
}
