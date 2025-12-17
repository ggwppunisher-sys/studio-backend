package observability

import (
	"fmt"
	"log/slog"
	"os"
)

const LogKeyAppVersion = "app_version"
const LogKeyAppName = "app_name"

// InitGlobalLogger initializes the global logger with app-specific attributes, log format, and log level.
func InitGlobalLogger(appName, appVersion string, logLevel string, logToJson bool) error {
	var level slog.Level
	err := level.UnmarshalText([]byte(logLevel))
	if err != nil {
		return fmt.Errorf("failed to parse log level (%s): %w", logLevel, err)
	}
	var handler slog.Handler
	if logToJson {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level})
	}
	attrs := []slog.Attr{
		slog.String(LogKeyAppName, appName),
		slog.String(LogKeyAppVersion, appVersion),
	}
	slog.SetDefault(slog.New(handler.WithAttrs(attrs)))
	return nil
}
