package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var notifications_pinCmd = &cobra.Command{
	Use:     "pin",
	Short:   "Mark all filtered or a specific notification as pinned",
	Aliases: []string{"p"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(notifications_pinCmd).Standalone()

	notifications_pinCmd.Flags().String("limit", "", "specify limit of items per page")
	notifications_pinCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	notifications_pinCmd.Flags().BoolP("mine", "m", false, "Show notifications across all your repositories instead of the current repository only")
	notifications_pinCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	notifications_pinCmd.Flags().StringP("page", "p", "", "specify page, default is 1")
	notifications_pinCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	notifications_pinCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	notifications_pinCmd.Flags().StringP("states", "s", "", "Comma-separated list of notification states to filter by. Available values:")
	notificationsCmd.AddCommand(notifications_pinCmd)

	// TODO completion
	carapace.Gen(notifications_pinCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
		"states": tea.ActionNotificationStates().UniqueList(","),
	})
}
