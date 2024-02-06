package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Remove stale lock files and outdated downloads",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanupCmd).Standalone()

	cleanupCmd.Flags().BoolS("s", "s", false, "Scrub the cache")
	cleanupCmd.Flags().BoolP("debug", "d", false, "Display any debugging information")
	cleanupCmd.Flags().BoolP("dry-run", "n", false, "Show what would be removed, but do not actually remove anything")
	cleanupCmd.Flags().BoolP("help", "h", false, "Show this message")
	cleanupCmd.Flags().String("prune", "", "Remove all cache files older than specified days")
	cleanupCmd.Flags().Bool("prune-prefix", false, "Only prune the symlinks and directories from the prefix and remove no other files")
	cleanupCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet")
	cleanupCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose")
	rootCmd.AddCommand(cleanupCmd)

	carapace.Gen(cleanupCmd).FlagCompletion(carapace.ActionMap{
		"prune": carapace.ActionValues("all"),
	})
}
