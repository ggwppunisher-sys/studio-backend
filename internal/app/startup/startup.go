package startup

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"studio-backend/internal/app/config"
	"studio-backend/internal/transport/apiserver"
)

func Run(ctx context.Context, cfg config.Config) error {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
	})
	server, err := apiserver.NewServer(cfg.Server, handler)
	if err != nil {
		return fmt.Errorf("failed to init server: %w", err)
	}

	go func() {
		<-ctx.Done()
		err = server.Shutdown(ctx)
		if err != nil {
			slog.Error("failed to shutdown", "error", err)
		}
	}()

	slog.Info("Server started", "port", cfg.Server.Port)
	err = server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}
