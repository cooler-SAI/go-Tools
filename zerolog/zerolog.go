package zerolog

import (
	"github.com/rs/zerolog"
	"os"
)

// Log is the global logger instance pre-configured with:
// - Output to stderr
// - Timestamp field
// - Default text output (non-JSON)
var Log = zerolog.New(os.Stderr).With().Timestamp().Logger()

// Init configures the global logger with:
// - Colored console output
// - Human-readable timestamps
// - Structured logging capabilities
func Init() {
	Log = zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "2006-01-02 15:04:05",
	}).With().Timestamp().Logger()
}
