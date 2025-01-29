package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/estifanos-neway/CLC/config"
	"github.com/estifanos-neway/CLC/internal/api/clc"
	"github.com/estifanos-neway/CLC/internal/api/gemini"
	"github.com/estifanos-neway/CLC/internal/pkg/display"
	"github.com/estifanos-neway/CLC/internal/pkg/logging"
)

func main() {
	if err := config.Load(); err != nil {
		display.ExitOnError(fmt.Errorf("unable to load app configuration \n %s", err), display.MsgTryAgain)
	}
	logging.ConfigureLogging()
	slog.Info("Logger configured.")

	printCmds := flag.Bool("p", false, "If set to true, the commands will be printed instead of being executed.")
	flag.Parse()

	prompt := strings.Join(flag.Args(), " ")
	if prompt == "" {
		fmt.Println("Please add a prompt!")
		os.Exit(1)
	}

	gemini := gemini.Gemini{
		Url:    config.AppConfig.Gemini.Url,
		ApiKey: config.AppConfig.Gemini.ApiKey,
	}

	clc := &clc.CLC{
		Gemini: &gemini,
		Prompt: prompt,
	}

	if err := clc.GetResponse(); err != nil {
		display.ExitOnError(err, display.MsgTryAgain)
	}

	switch {
	case *printCmds:
		clc.Print()
	default:
		clc.Exec()
	}

}
