package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
	"github.com/tommartensen/jaga/pkg/client/segment"
	"github.com/tommartensen/jaga/pkg/logging"
)

var (
	version = "local"
)

func newRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "jagactl",
		Short:   "CLI for the Jaga project",
		Long:    "jagactl is a tool to find segments within your area to achieve a PR or KOM.",
		Version: version,
	}

	rootCmd.AddCommand(segment.Command())
	return rootCmd
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	defer logging.Logger.Sync()

	ctx := context.Background()
	rootCmd := newRootCommand()
	rootCmd.ExecuteContext(ctx)
}
