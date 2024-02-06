package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var depsCmd = &cobra.Command{
	Use:   "deps",
	Short: "Show dependencies for formula",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(depsCmd).Standalone()

	depsCmd.Flags().BoolS("1", "1", false, "Show only the direct dependencies declared in the formula")
	depsCmd.Flags().Bool("HEAD", false, "Show dependencies for HEAD version instead of stable version")
	depsCmd.Flags().Bool("annotate", false, "Mark any build, test, implicit, optional, or recommended dependencies")
	depsCmd.Flags().Bool("cask", false, "Treat all named arguments as casks")
	depsCmd.Flags().Bool("casks", false, "Treat all named arguments as casks")
	depsCmd.Flags().BoolP("debug", "d", false, "Display any debugging information")
	depsCmd.Flags().Bool("declared", false, "Show only the direct dependencies declared in the formula")
	depsCmd.Flags().Bool("direct", false, "Show only the direct dependencies declared in the formula")
	depsCmd.Flags().Bool("dot", false, "Show text-based graph description in DOT format")
	depsCmd.Flags().Bool("eval-all", false, "Evaluate all available formulae and casks")
	depsCmd.Flags().Bool("for-each", false, "List dependencies for each provided formula, one formula per line")
	depsCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae")
	depsCmd.Flags().Bool("formulae", false, "Treat all named arguments as formulae")
	depsCmd.Flags().Bool("full-name", false, "List dependencies by their full name")
	depsCmd.Flags().Bool("graph", false, "Show dependencies as a directed graph")
	depsCmd.Flags().BoolP("help", "h", false, "Show this message")
	depsCmd.Flags().Bool("include-build", false, "Include :build dependencies for formula")
	depsCmd.Flags().Bool("include-optional", false, "Include :optional dependencies for formula")
	depsCmd.Flags().Bool("include-requirements", false, "Include requirements in addition to dependencies for formula")
	depsCmd.Flags().Bool("include-test", false, "Include :test dependencies for formula")
	depsCmd.Flags().Bool("installed", false, "List dependencies for formulae that are currently installed")
	depsCmd.Flags().Bool("missing", false, "Show only missing dependencies")
	depsCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet")
	depsCmd.Flags().Bool("skip-recommended", false, "Skip :recommended dependencies for formula")
	depsCmd.Flags().BoolP("topological", "n", false, "Sort dependencies in topological order")
	depsCmd.Flags().Bool("tree", false, "Show dependencies as a tree")
	depsCmd.Flags().Bool("union", false, "Show the union of dependencies for multiple formula")
	depsCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose")
	rootCmd.AddCommand(depsCmd)

	depsCmd.MarkFlagsMutuallyExclusive("1", "direct", "declared")

	carapace.Gen(depsCmd).PositionalAnyCompletion(
		action.ActionSearch(depsCmd).FilterArgs(),
	)
}
