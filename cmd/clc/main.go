package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/estifanos-neway/CLC/config"
	"github.com/estifanos-neway/CLC/internal/pkg/helpers"
)

func main() {
	if err := config.Load(); err != nil {
		helpers.ExitOnError(fmt.Errorf("unable to load app configuration \n %s", err))
	}
	slog.Info("Configuration loaded.")

	logLevel := slog.LevelDebug
	if config.AppConfig.Environment == config.EnvironmentProd {
		logLevel = slog.LevelInfo
	}

	logHandlerOptions := slog.HandlerOptions{
		Level: logLevel,
	}

	logHandler := slog.NewJSONHandler(os.Stderr, &logHandlerOptions)
	logger := slog.New(logHandler)
	slog.SetDefault(logger)
	slog.Info("Logger configured.")

	// printCmds := flag.Bool("p", true, "If set to true, the commands will be printed instead of being executed.")
	flag.Parse()

}
