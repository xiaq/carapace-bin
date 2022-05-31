package cmd

import (
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kotlin",
	Short: "run Kotlin programs, scripts or REPL",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "-D") || strings.HasPrefix(arg, "-J") || strings.HasPrefix(arg, "-X") {
			name := strings.TrimPrefix(strings.Split(arg, "=")[0], "-")
			rootCmd.Flags().String(name, "", "")
			rootCmd.Flag(name).NoOptDefVal = " "
		}
	}

	carapace.Override(carapace.Opts{
		LongShorthand: true,
	})

	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("D", "D", "Set a system JVM property")
	rootCmd.Flags().String("J", "J", "Pass an option directly to JVM")
	rootCmd.Flags().String("X", "X", "Pass -X argument to the compiler")
	rootCmd.Flags().String("classpath", "", "Paths where to find user class files")
	rootCmd.Flags().Bool("compiler-path", false, "Kotlin compiler classpath for compiling script or expression or running REPL")
	rootCmd.Flags().StringP("expression", "e", "", "Evaluates the expression and prints the result")
	rootCmd.Flags().BoolP("help", "h", false, "Print a synopsis of options")
	rootCmd.Flags().String("howtorun", "", "How to run the supplied command with arguments")
	rootCmd.Flags().Bool("no-reflect", false, "Don't include Kotlin reflection implementation into classpath")
	rootCmd.Flags().Bool("no-stdlib", false, "Don't include Kotlin standard library into classpath")
	rootCmd.Flags().Bool("version", false, "Display Kotlin version")

	rootCmd.Flag("D").NoOptDefVal = " "
	rootCmd.Flag("J").NoOptDefVal = " "
	rootCmd.Flag("X").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"classpath": carapace.ActionDirectories(),
		"howtorun":  carapace.ActionValues("guess", "classfile", "jar", "script"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
