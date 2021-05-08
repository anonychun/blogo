package logger

import (
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

var once sync.Once
var logger zerolog.Logger

func Log() *zerolog.Logger {
	once.Do(func() {
		logger = zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			NoColor:    true,
			TimeFormat: time.RFC3339,
		}).With().Timestamp().Logger().Level(zerolog.GlobalLevel())
	})
	return &logger
}
