package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Download a bottle  or source packages for formulae and binaries for casks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fetchCmd).Standalone()

	fetchCmd.Flags().Bool("HEAD", false, "Fetch HEAD version instead of stable version")
	fetchCmd.Flags().Bool("arch", false, "Download for the given CPU architecture")
	fetchCmd.Flags().Bool("bottle-tag", false, "Download a bottle for given tag")
	fetchCmd.Flags().Bool("build-bottle", false, "Download source packages rather than a bottle")
	fetchCmd.Flags().BoolP("build-from-source", "s", false, "Download source packages rather than a bottle")
	fetchCmd.Flags().Bool("cask", false, "Treat all named arguments as casks")
	fetchCmd.Flags().Bool("casks", false, "Treat all named arguments as casks")
	fetchCmd.Flags().BoolP("debug", "d", false, "Display any debugging information")
	fetchCmd.Flags().Bool("deps", false, "Also download dependencies for any listed formula")
	fetchCmd.Flags().BoolP("force", "f", false, "Remove a previously cached version and re-fetch")
	fetchCmd.Flags().Bool("force-bottle", false, "Download a bottle if it exists for the current or newest version of macOS")
	fetchCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae")
	fetchCmd.Flags().Bool("formulae", false, "Treat all named arguments as formulae")
	fetchCmd.Flags().BoolP("help", "h", false, "Show this message")
	fetchCmd.Flags().Bool("no-quarantine", false, "Disable quarantining of downloads")
	fetchCmd.Flags().Bool("os", false, "Download for the given operating system")
	fetchCmd.Flags().Bool("quarantine", false, "Enable quarantining of downloads")
	fetchCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet")
	fetchCmd.Flags().Bool("retry", false, "Retry if downloading fails")
	fetchCmd.Flags().BoolP("verbose", "v", false, "Do a verbose VCS checkout, if the URL represents a VCS")
	rootCmd.AddCommand(fetchCmd)

	carapace.Gen(fetchCmd).PositionalAnyCompletion(
		action.ActionSearch(fetchCmd),
	)
}
