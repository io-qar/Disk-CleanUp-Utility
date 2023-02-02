package adapters

import (
	"clean-utility/internal/interfaces"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct{
	output zerolog.ConsoleWriter
}

func NewLogger() interfaces.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
    return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}

	return Logger{output}
}

func (lg Logger) Info(s string, vs ...string) {
	log := zerolog.New(lg.output).With().Timestamp().Logger()
	if vs != nil {
		for _, v := range vs {
			log.Info().Msgf(s, v)
		}
	} else {
		log.Info().Msg(s)
	}
}

func (lg Logger) Error(s string, vs ...string) {
	log := zerolog.New(lg.output).With().Timestamp().Logger()
	if vs != nil {
		for _, v := range vs {
			log.Error().Msgf(s, v)
		}
	} else {
		log.Error().Msg(s)
	}
}

func (lg Logger) Warn(s string, vs ...string) {
	log := zerolog.New(lg.output).With().Timestamp().Logger()
	if vs != nil {
		for _, v := range vs {
			log.Warn().Msgf(s, v)
		}
	} else {
		log.Warn().Msg(s)
	}
}