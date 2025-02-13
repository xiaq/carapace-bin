package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/jj"
	"github.com/rsteube/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff [OPTIONS] [PATHS]...",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diffCmd).Standalone()

	diffCmd.Flags().Bool("color-words", false, "Show a word-level diff with changes indicated only by color")
	diffCmd.Flags().String("from", "", "Show changes from this revision")
	diffCmd.Flags().Bool("git", false, "Show a Git-format diff")
	diffCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	diffCmd.Flags().StringP("revision", "r", "", "Show changes in this revision, compared to its parent(s)")
	diffCmd.Flags().Bool("stat", false, "Show a histogram of the changes")
	diffCmd.Flags().BoolP("summary", "s", false, "For each path, show only whether it was modified, added, or removed")
	diffCmd.Flags().String("to", "", "Show changes to this revision")
	diffCmd.Flags().String("tool", "", "Generate diff by external command")
	diffCmd.Flags().Bool("types", false, "For each path, show only its type before and after")
	rootCmd.AddCommand(diffCmd)

	carapace.Gen(diffCmd).FlagCompletion(carapace.ActionMap{
		"from":     jj.ActionRevs(jj.RevOption{}.Default()),
		"revision": jj.ActionRevs(jj.RevOption{}.Default()),
		"to":       jj.ActionRevs(jj.RevOption{}.Default()),
		"tool":     bridge.ActionCarapaceBin().Split(),
	})

	carapace.Gen(diffCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			revisions := []string{"@"}
			if f := diffCmd.Flag("from"); f.Changed {
				revisions = []string{f.Value.String()}
			}
			if f := diffCmd.Flag("to"); f.Changed {
				revisions = append(revisions, f.Value.String())
			}
			return jj.ActionRevDiffs(revisions...)
		}),
	)
}
