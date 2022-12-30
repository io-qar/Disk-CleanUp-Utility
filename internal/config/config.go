package config

const (
	defaultVolume uint64 = 50
)

type Config struct {
	Folders   []string
	MaxVolume uint64
	Telegram  struct {
		Channel  string
		BotToken string
	}
}

func NewConfig() Config {

	return Config{}
}
