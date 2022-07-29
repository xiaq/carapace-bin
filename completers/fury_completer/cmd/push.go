package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Upload a new version of a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pushCmd).Standalone()
	pushCmd.Flags().Bool("public", false, "Create as public package")
	pushCmd.Flags().Bool("quiet", false, "Do not show progress bar")
	rootCmd.AddCommand(pushCmd)
}
