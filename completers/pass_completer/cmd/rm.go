package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pass_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove existing password or directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rmCmd).Standalone()
	rmCmd.Flags().BoolP("recursive", "r", false, "remove recursively")
	rmCmd.Flags().BoolP("force", "f", false, "remove forcefully")

	rootCmd.AddCommand(rmCmd)

	carapace.Gen(rmCmd).PositionalCompletion(
		carapace.ActionCallback(func(args []string) carapace.Action {
			return action.ActionPassNames().Invoke(args).ToMultipartsA("/")
		}),
	)
}
