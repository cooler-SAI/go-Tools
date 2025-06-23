package zerolog

import (
	zl "github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

// Logger is the global logger instance
var Logger = log.Logger

// ConfigureZerologConsoleWriter sets up the global zerolog logger
func ConfigureZerologConsoleWriter() {
	zl.SetGlobalLevel(zl.InfoLevel)
	zl.TimeFieldFormat = "2006-01-02 15:04:05"
	log.Logger = zl.New(zl.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	Logger = log.Logger // Update our exported logger
}
