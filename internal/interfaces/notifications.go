package interfaces

import "clean-utility/internal/entity"

type Notifications interface {
	NewMessage() entity.Message
	SendMessage(entity.Message) error
}
