package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_renameCmd = &cobra.Command{
	Use:     "rename [<new-name>]",
	Short:   "Rename a repository",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_renameCmd).Standalone()

	repo_renameCmd.Flags().Bool("confirm", false, "Skip confirmation prompt")
	repo_renameCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	repo_renameCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	repo_renameCmd.Flag("confirm").Hidden = true
	repoCmd.AddCommand(repo_renameCmd)

	carapace.Gen(repo_renameCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepoOverride(repo_renameCmd),
	})
}
