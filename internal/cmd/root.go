package cmd

import (
	"os"

	"github.com/saman2000hoseini/k8s-training/internal/cmd/server"
	"github.com/saman2000hoseini/k8s-training/internal/config"

	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
)

// ExitFailure status code.
const ExitFailure = 1

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cfg := config.New()

	// nolint: exhaustivestruct
	root := &cobra.Command{
		Use:   "visitor",
		Short: "Simple webpage visitor counter",
	}

	server.Register(root, cfg)

	if err := root.Execute(); err != nil {
		log.Errorf("failed to execute root command: %s", err)
		os.Exit(ExitFailure)
	}
}
