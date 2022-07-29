package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var yankCmd = &cobra.Command{
	Use:   "yank",
	Short: "Remove a package version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(yankCmd).Standalone()
	yankCmd.Flags().StringP("version", "v", "", "Version")
	rootCmd.AddCommand(yankCmd)
}
