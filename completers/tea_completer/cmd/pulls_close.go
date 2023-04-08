package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var pulls_closeCmd = &cobra.Command{
	Use:   "close",
	Short: "Change state of one or more pull requests to 'closed'",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pulls_closeCmd).Standalone()

	pulls_closeCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	pulls_closeCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	pulls_closeCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	pulls_closeCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	pullsCmd.AddCommand(pulls_closeCmd)

	// TODO completion
	carapace.Gen(pulls_closeCmd).FlagCompletion(carapace.ActionMap{
		"output": tea.ActionOutputFormats(),
	})
}
