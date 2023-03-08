package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var codespace_rebuildCmd = &cobra.Command{
	Use:   "rebuild",
	Short: "Rebuild a codespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_rebuildCmd).Standalone()

	codespace_rebuildCmd.PersistentFlags().StringP("codespace", "c", "", "Name of the codespace")
	codespace_rebuildCmd.Flags().Bool("full", false, "perform a full rebuild")
	codespace_rebuildCmd.PersistentFlags().StringP("repo", "R", "", "Filter codespace selection by repository name (user/repo)")
	codespaceCmd.AddCommand(codespace_rebuildCmd)

	carapace.Gen(codespace_rebuildCmd).FlagCompletion(carapace.ActionMap{
		"codespace": action.ActionCodespaces(),
		"repo":      gh.ActionOwnerRepositories(gh.HostOpts{}),
	})
}
