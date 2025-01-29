package logging

import (
	"log/slog"
	"os"

	"github.com/estifanos-neway/CLC/config"
)

func ConfigureLogging() {
	// TODO Create slog.Temp()
	logLevel := slog.LevelDebug
	if config.AppConfig.Environment == config.EnvironmentProd {
		logLevel = 10
	}

	logHandlerOptions := slog.HandlerOptions{
		Level: logLevel,
	}

	logHandler := slog.NewJSONHandler(os.Stderr, &logHandlerOptions)
	logger := slog.New(logHandler)
	slog.SetDefault(logger)
}
