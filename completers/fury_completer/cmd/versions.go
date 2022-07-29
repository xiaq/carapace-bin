package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var versionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "List versions for a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionsCmd).Standalone()
	rootCmd.AddCommand(versionsCmd)
}
