package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Restarts the remote supervisord",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reloadCmd).Standalone()

	rootCmd.AddCommand(reloadCmd)
}
