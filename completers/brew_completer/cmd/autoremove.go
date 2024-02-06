package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var autoremoveCmd = &cobra.Command{
	Use:   "autoremove",
	Short: "Uninstall formulae that are no longer needed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoremoveCmd).Standalone()

	autoremoveCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	autoremoveCmd.Flags().BoolP("dry-run", "n", false, "List what would be uninstalled, but do not")
	autoremoveCmd.Flags().BoolP("help", "h", false, "Show this message.")
	autoremoveCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	autoremoveCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(autoremoveCmd)
}
