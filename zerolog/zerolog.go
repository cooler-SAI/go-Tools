package zerolog

import (
	"os"

	zl "github.com/rs/zerolog"  // Alias zerolog to zl to avoid confusion with our package
	"github.com/rs/zerolog/log" // Import the global zerolog.Logger instance
)

// ConfigureConsoleWriter sets up the global zerolog logger
// to output to the console (stderr) in a human-readable format.
//
// You should call this function once at the beginning of your main application.
func ConfigureConsoleWriter() {
	// Set the global logging level (e.g., zl.InfoLevel).
	// This can be changed to zl.DebugLevel for more verbose logs during development.
	zl.SetGlobalLevel(zl.InfoLevel)

	// Configure ConsoleWriter to output to os.Stderr.
	// ConsoleWriter formats logs for easy human readability.
	log.Logger = log.Output(zl.ConsoleWriter{Out: os.Stderr})

	// Optional: You can add time to logs.
	// zl.TimeFieldFormat = zl.TimeFormatUnix
	// Or a more human-readable format:
	// zl.TimeFieldFormat = "2006-01-02 15:04:05" // Example RFC3339 without milliseconds
}
