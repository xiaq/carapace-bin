package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var casksCmd = &cobra.Command{
	Use:   "casks",
	Short: "List all locally installable casks including short name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(casksCmd).Standalone()


	rootCmd.AddCommand(casksCmd)
}
