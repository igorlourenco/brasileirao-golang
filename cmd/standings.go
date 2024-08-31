package cmd

import (
	brzao "igorlourenco/brzao/brzaoapi"

	"github.com/spf13/cobra"
)

var standingsCmd = &cobra.Command{
	Use:     "standings",
	Aliases: []string{"standings"},
	Short:   "Get brzao standings",
	Long:    "Get brzao standings",
	Run: func(cmd *cobra.Command, args []string) {
		brzao.Standings()
	},
}

func init() {
	rootCmd.AddCommand(standingsCmd)
}
