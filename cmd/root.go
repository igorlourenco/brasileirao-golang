package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "brzao",
	Short: "brzao is a cli tool for getting brzao data",
	Long:  "brzao is a cli tool for getting brzao data",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dados do brzao!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing Zero '%s'\n", err)
		os.Exit(1)
	}
}
