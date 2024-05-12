package main

import (
	"fmt"
	"log/slog"
	"os"

	"mate/pkg/cmd"

	"github.com/charmbracelet/log"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	if version == "dev" {
		slog.SetDefault(slog.New(log.NewWithOptions(os.Stdout, log.Options{
			Level: log.DebugLevel,
		})))

		slog.Info("development mode detected")
	}

	c := cmd.NewRoot(version, commit, date)
	if err := c.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
