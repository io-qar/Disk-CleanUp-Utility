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
		return errors.New("Empty reciever")
	}

	fmt.Printf("[fake-notification] Получатель: %s", msg.To)
	fmt.Printf("[fake-notification] Текст: \n%s", msg.Text)

	return nil
}

func (f FakeTgBot) NewMessage() entity.Message {
	return entity.Message{}
}
