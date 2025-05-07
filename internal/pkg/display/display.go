package display

import (
	"fmt"
	"log/slog"
	"os"
)

func ExitOnError(err error, msg string) {
	slog.Error(err.Error())
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

const (
	MsgTryAgain string = "Something went wrong! Please try again."
)
