package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var formulaeCmd = &cobra.Command{
	Use:   "formulae",
	Short: "List all locally installable formulae including short names",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(formulaeCmd).Standalone()

	rootCmd.AddCommand(formulaeCmd)
}
