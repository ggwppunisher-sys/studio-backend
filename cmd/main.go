package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"studio-backend/internal/app/config"
	"studio-backend/internal/app/startup"
	"syscall"

	"github.com/joho/godotenv"
)

const appName = "studio-backend"

var version = "v0.1.0" // version передается через -ldflags.

func main() {
	processCmdArgs()

	appConfig, err := config.New(appName, version)
	if err != nil {
		slog.Error("failed to load config", slog.Any("error", err))
		return
	}

	appCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err = startup.Run(appCtx, appConfig); err != nil {
		slog.Error("application stopped with an error", slog.Any("error", err))
	} else {
		slog.Info("application stopped gracefully")
	}
}

// processCmdArgs configures command-line flags, handles version display, and optionally loads environment variables from a file.
func processCmdArgs() {
	var versionCheck bool
	var envFileToLoad string

	flag.BoolVar(&versionCheck, "version", false, "Show version and exit.")
	flag.StringVar(&envFileToLoad, "env", "", "File to preload environment from. Existing values won't be changed.")
	flag.Parse()

	if versionCheck {
		fmt.Println(appName, version)
		os.Exit(0)
	}
	if envFileToLoad != "" {
		if err := godotenv.Load(envFileToLoad); err != nil {
			slog.Error("failed to load specified environment file", slog.Any("error", err))
			os.Exit(1)
		} else {
			slog.Info(fmt.Sprintf("environment updated from file: %s", envFileToLoad))
		}
	}
}
