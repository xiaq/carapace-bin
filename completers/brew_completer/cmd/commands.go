package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var commandsCmd = &cobra.Command{
	Use:   "commands",
	Short: "Show lists of built-in and external commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(commandsCmd).Standalone()

	commandsCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	commandsCmd.Flags().BoolP("help", "h", false, "Show this message.")
	commandsCmd.Flags().Bool("include-aliases", false, "Include aliases of internal commands.")
	commandsCmd.Flags().BoolP("quiet", "q", false, "List only the names of commands without")
	commandsCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(commandsCmd)
}
