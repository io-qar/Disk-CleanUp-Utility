package adapters

import (
	"clean-utility/internal/entity"
	"clean-utility/internal/interfaces"
	"errors"
	"fmt"
)

type FakeTgBot struct {
	Token string
}

func NewFakeTgBot(token string) interfaces.Notifications {
	return FakeTgBot{token}
}

func (f FakeTgBot) SendMessage(msg entity.Message) error {
	if msg.To == "" {
		return errors.New("Empty receiver\n")
	}

	fmt.Printf("[fake-notification] Получатель: %s\n", msg.To)
	fmt.Printf("[fake-notification] Текст: %s\n", msg.Text)

	return nil
}

func (f FakeTgBot) NewMessage() entity.Message {
	return entity.Message{}
}