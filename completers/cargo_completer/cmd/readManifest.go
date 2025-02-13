package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var readManifestCmd = &cobra.Command{
	Use:   "read-manifest",
	Short: "Print a JSON representation of a Cargo.toml manifest.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(readManifestCmd).Standalone()

	readManifestCmd.Flags().BoolP("help", "h", false, "Print help")
	readManifestCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	readManifestCmd.Flags().BoolP("quiet", "q", false, "Do not print cargo log messages")
	rootCmd.AddCommand(readManifestCmd)

	carapace.Gen(readManifestCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
