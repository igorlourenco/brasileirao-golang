package cmd

import (
	"igorlourenco/brzao/brzao"

	"github.com/spf13/cobra"
)

var matchDate string

var getMatchesCmd = &cobra.Command{
	Use:     "matches",
	Aliases: []string{"matches"},
	Short:   "Get brzao matches",
	Long:    "Get brzao matches",
	Run: func(cmd *cobra.Command, args []string) {
		brzao.Matches(brzao.DateOption(matchDate))
	},
}

func init() {
	getMatchesCmd.Flags().StringVarP(&matchDate, "date", "d", "today", "Date to filter matches (today, tomorrow, yesterday)")
	rootCmd.AddCommand(getMatchesCmd)
}
