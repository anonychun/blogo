package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var logger = zerolog.New(zerolog.ConsoleWriter{
	Out:        os.Stdout,
	NoColor:    false,
	TimeFormat: time.RFC3339,
}).With().Timestamp().Logger().Level(zerolog.GlobalLevel())

func Log() *zerolog.Logger { return &logger }
