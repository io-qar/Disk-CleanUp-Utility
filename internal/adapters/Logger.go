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

func (lg Logger) Info(s string) {
	log := zerolog.New(lg.output).With().Timestamp().Logger()

	log.Info().Msg(s)
}

func (lg Logger) Infof(s string, v string) {
	log := zerolog.New(lg.output).With().Timestamp().Logger()
	
	log.Info().Msgf(s, v)
}

func (lg Logger) Error(s string) {
	log := zerolog.New(lg.output).With().Timestamp().Logger()

	log.Error().Msg(s)
}

func (lg Logger) Errorf(s string, v string) {
	log := zerolog.New(lg.output).With().Timestamp().Logger()

	log.Error().Msgf(s, v)
}

func (lg Logger) Warn(s string) {
	log := zerolog.New(lg.output).With().Timestamp().Logger()

	log.Warn().Msg(s)
}

func (lg Logger) Warnf(s string, v string) {
	log := zerolog.New(lg.output).With().Timestamp().Logger()

	log.Warn().Msgf(s, v)
}