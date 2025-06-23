package zerolog

import (
	"github.com/rs/zerolog"
	"os"
)

var Log zerolog.Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

func InitConsoleLogger() {
	Log = zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "2006-01-02 15:04:05",
	}).With().Timestamp().Logger()
}
