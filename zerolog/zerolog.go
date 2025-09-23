package zerolog

import (
	"os"

	"github.com/rs/zerolog"
)

// Log is the global logger instance
var Log = zerolog.New(os.Stderr).With().Timestamp().Logger()

// consoleOutput wraps ConsoleWriter to track the instance
type consoleOutput struct {
	writer *zerolog.ConsoleWriter
}

var currentOutput *consoleOutput

// Init configures the global logger
func Init() {
	consoleWriter := &zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "2006-01-02 15:04:05",
	}

	Log = zerolog.New(consoleWriter).With().Timestamp().Logger()
	currentOutput = &consoleOutput{writer: consoleWriter}
}

// Sync flushes any buffered log entries
func Sync() error {
	if currentOutput != nil && currentOutput.writer != nil {
		// If the ConsoleWriter has a Sync method (some implementations might)
		if syncer, ok := currentOutput.writer.Out.(interface{ Sync() error }); ok {
			return syncer.Sync()
		}
	}
	return nil
}
