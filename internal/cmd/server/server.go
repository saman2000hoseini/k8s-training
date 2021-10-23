package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/saman2000hoseini/k8s-training/internal/config"
	"github.com/saman2000hoseini/k8s-training/internal/db"
	"github.com/saman2000hoseini/k8s-training/internal/handler"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

func main(cfg config.Config) {
	app := echo.New()

	redisClient := db.New(cfg.Redis)
	if err := redisClient.Set(handler.KEY, 0, 0).Err(); err != nil {
		log.Fatalf("init visitors failed: %s", err)
	}

	handler.Visit{
		Store: redisClient,
	}.Register(app.Group("/api"))

	handler.Healthz{}.Register(app.Group(""))

	if err := app.Start(fmt.Sprintf(":%d", cfg.Server.Port)); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("echo initiation failed: %s", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

// Register server command.
func Register(root *cobra.Command, cfg config.Config) {
	root.AddCommand(
		// nolint: exhaustivestruct
		&cobra.Command{
			Use:   "server",
			Short: "Run server to serve the requests",
			Run: func(cmd *cobra.Command, args []string) {
				main(cfg)
			},
		},
	)
}
