package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var descCmd = &cobra.Command{
	Use:   "desc",
	Short: "Display formula's name and one-line description",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(descCmd).Standalone()

	descCmd.Flags().Bool("cask", false, "Treat all named arguments as casks")
	descCmd.Flags().Bool("casks", false, "Treat all named arguments as casks")
	descCmd.Flags().Bool("debug", false, "Display any debugging information")
	descCmd.Flags().BoolP("description", "d", false, "Search just descriptions for text")
	descCmd.Flags().Bool("eval-all", false, "Evaluate all available formulae and casks")
	descCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae")
	descCmd.Flags().Bool("formulae", false, "Treat all named arguments as formulae")
	descCmd.Flags().BoolP("help", "h", false, "Show this message")
	descCmd.Flags().BoolP("name", "n", false, "Search just names for text")
	descCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet")
	descCmd.Flags().BoolP("search", "s", false, "Search both names and descriptions for text")
	descCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose")
	rootCmd.AddCommand(descCmd)

	carapace.Gen(descCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if descCmd.Flag("eval-all").Changed {
				return action.ActionSearch(descCmd)
			}
			return action.ActionList(uninstallCmd)
		}).FilterArgs(),
	)
}
