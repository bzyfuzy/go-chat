package cmd

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "chat-cli",
	Short:   "Peer-to-Peer CLI chat application",
	Version: "v0.0.1",
}

func Execute() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		for s := range sig {
			fmt.Println(s.String())
			os.Exit(0)
		}
	}()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
