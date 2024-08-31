package cmd

import (
	"igorlourenco/brzao/brzao"

	"github.com/spf13/cobra"
)

var date string

var matchesCmd = &cobra.Command{
	Use:     "matches",
	Aliases: []string{"matches"},
	Short:   "Get brzao matches",
	Long:    "Get brzao matches",
	Run: func(cmd *cobra.Command, args []string) {
		brzao.Matches(brzao.DateOption(date))
	},
}

func init() {
	matchesCmd.Flags().StringVarP(&date, "date", "d", "today", "Date to filter matches (today, tomorrow, yesterday)")
	rootCmd.AddCommand(matchesCmd)
}
