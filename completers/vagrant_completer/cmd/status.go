package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/vagrant"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "outputs status of the vagrant machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statusCmd).Standalone()

	rootCmd.AddCommand(statusCmd)

	carapace.Gen(statusCmd).PositionalCompletion(
		vagrant.ActionMachines(),
	)
}
