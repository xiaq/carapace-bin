package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gum"
	"github.com/spf13/cobra"
)

var pagerCmd = &cobra.Command{
	Use:   "pager",
	Short: "Scroll through a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pagerCmd).Standalone()

	pagerCmd.Flags().String("align", "", "Text Alignment")
	pagerCmd.Flags().String("background", "", "Background Color")
	pagerCmd.Flags().Bool("bold", false, "Bold text")
	pagerCmd.Flags().String("border", "", "Border Style")
	pagerCmd.Flags().String("border-background", "", "Border Background Color")
	pagerCmd.Flags().String("border-foreground", "", "Border Foreground Color")
	pagerCmd.Flags().Bool("faint", false, "Faint text")
	pagerCmd.Flags().String("foreground", "", "Foreground Color")
	pagerCmd.Flags().String("height", "", "Text height")
	pagerCmd.Flags().String("help.align", "", "Text Alignment")
	pagerCmd.Flags().String("help.background", "", "Background Color")
	pagerCmd.Flags().Bool("help.bold", false, "Bold text")
	pagerCmd.Flags().String("help.border", "", "Border Style")
	pagerCmd.Flags().String("help.border-background", "", "Border Background Color")
	pagerCmd.Flags().String("help.border-foreground", "", "Border Foreground Color")
	pagerCmd.Flags().Bool("help.faint", false, "Faint text")
	pagerCmd.Flags().String("help.foreground", "", "Foreground Color")
	pagerCmd.Flags().String("help.height", "", "Text height")
	pagerCmd.Flags().Bool("help.italic", false, "Italicize text")
	pagerCmd.Flags().String("help.margin", "", "Text margin")
	pagerCmd.Flags().String("help.padding", "", "Text padding")
	pagerCmd.Flags().Bool("help.strikethrough", false, "Strikethrough text")
	pagerCmd.Flags().Bool("help.underline", false, "Underline text")
	pagerCmd.Flags().String("help.width", "", "Text width")
	pagerCmd.Flags().Bool("italic", false, "Italicize text")
	pagerCmd.Flags().String("line-number.align", "", "Text Alignment")
	pagerCmd.Flags().String("line-number.background", "", "Background Color")
	pagerCmd.Flags().Bool("line-number.bold", false, "Bold text")
	pagerCmd.Flags().String("line-number.border", "", "Border Style")
	pagerCmd.Flags().String("line-number.border-background", "", "Border Background Color")
	pagerCmd.Flags().String("line-number.border-foreground", "", "Border Foreground Color")
	pagerCmd.Flags().Bool("line-number.faint", false, "Faint text")
	pagerCmd.Flags().String("line-number.foreground", "", "Foreground Color")
	pagerCmd.Flags().String("line-number.height", "", "Text height")
	pagerCmd.Flags().Bool("line-number.italic", false, "Italicize text")
	pagerCmd.Flags().String("line-number.margin", "", "Text margin")
	pagerCmd.Flags().String("line-number.padding", "", "Text padding")
	pagerCmd.Flags().Bool("line-number.strikethrough", false, "Strikethrough text")
	pagerCmd.Flags().Bool("line-number.underline", false, "Underline text")
	pagerCmd.Flags().String("line-number.width", "", "Text width")
	pagerCmd.Flags().String("margin", "", "Text margin")
	pagerCmd.Flags().String("match-highlight.align", "", "Text Alignment")
	pagerCmd.Flags().String("match-highlight.background", "", "Background Color")
	pagerCmd.Flags().Bool("match-highlight.bold", false, "Bold text")
	pagerCmd.Flags().String("match-highlight.border", "", "Border Style")
	pagerCmd.Flags().String("match-highlight.border-background", "", "Border Background Color")
	pagerCmd.Flags().String("match-highlight.border-foreground", "", "Border Foreground Color")
	pagerCmd.Flags().Bool("match-highlight.faint", false, "Faint text")
	pagerCmd.Flags().String("match-highlight.foreground", "", "Foreground Color")
	pagerCmd.Flags().String("match-highlight.height", "", "Text height")
	pagerCmd.Flags().Bool("match-highlight.italic", false, "Italicize text")
	pagerCmd.Flags().String("match-highlight.margin", "", "Text margin")
	pagerCmd.Flags().String("match-highlight.padding", "", "Text padding")
	pagerCmd.Flags().Bool("match-highlight.strikethrough", false, "Strikethrough text")
	pagerCmd.Flags().Bool("match-highlight.underline", false, "Underline text")
	pagerCmd.Flags().String("match-highlight.width", "", "Text width")
	pagerCmd.Flags().String("match.align", "", "Text Alignment")
	pagerCmd.Flags().String("match.background", "", "Background Color")
	pagerCmd.Flags().Bool("match.bold", false, "Bold text")
	pagerCmd.Flags().String("match.border", "", "Border Style")
	pagerCmd.Flags().String("match.border-background", "", "Border Background Color")
	pagerCmd.Flags().String("match.border-foreground", "", "Border Foreground Color")
	pagerCmd.Flags().Bool("match.faint", false, "Faint text")
	pagerCmd.Flags().String("match.foreground", "", "Foreground Color")
	pagerCmd.Flags().String("match.height", "", "Text height")
	pagerCmd.Flags().Bool("match.italic", false, "Italicize text")
	pagerCmd.Flags().String("match.margin", "", "Text margin")
	pagerCmd.Flags().String("match.padding", "", "Text padding")
	pagerCmd.Flags().Bool("match.strikethrough", false, "Strikethrough text")
	pagerCmd.Flags().Bool("match.underline", false, "Underline text")
	pagerCmd.Flags().String("match.width", "", "Text width")
	pagerCmd.Flags().String("padding", "", "Text padding")
	pagerCmd.Flags().Bool("show-line-numbers", false, "Show line numbers")
	pagerCmd.Flags().Bool("soft-wrap", false, "Soft wrap lines")
	pagerCmd.Flags().Bool("strikethrough", false, "Strikethrough text")
	pagerCmd.Flags().String("timeout", "", "Timeout until command exits")
	pagerCmd.Flags().Bool("underline", false, "Underline text")
	pagerCmd.Flags().String("width", "", "Text width")
	rootCmd.AddCommand(pagerCmd)

	carapace.Gen(pagerCmd).FlagCompletion(carapace.ActionMap{
		"align":                             gum.ActionAlignments(),
		"background":                        gum.ActionColors(),
		"border":                            gum.ActionBorders(),
		"border-background":                 gum.ActionColors(),
		"border-foreground":                 gum.ActionColors(),
		"foreground":                        gum.ActionColors(),
		"help.align":                        gum.ActionAlignments(),
		"help.background":                   gum.ActionColors(),
		"help.border":                       gum.ActionBorders(),
		"help.border-background":            gum.ActionColors(),
		"help.border-foreground":            gum.ActionColors(),
		"help.foreground":                   gum.ActionColors(),
		"line-number.align":                 gum.ActionAlignments(),
		"line-number.background":            gum.ActionColors(),
		"line-number.border":                gum.ActionBorders(),
		"line-number.border-background":     gum.ActionColors(),
		"line-number.border-foreground":     gum.ActionColors(),
		"line-number.foreground":            gum.ActionColors(),
		"match-highlight.align":             gum.ActionAlignments(),
		"match-highlight.background":        gum.ActionColors(),
		"match-highlight.border":            gum.ActionBorders(),
		"match-highlight.border-background": gum.ActionColors(),
		"match-highlight.border-foreground": gum.ActionColors(),
		"match-highlight.foreground":        gum.ActionColors(),
		"match.align":                       gum.ActionAlignments(),
		"match.background":                  gum.ActionColors(),
		"match.border":                      gum.ActionBorders(),
		"match.border-background":           gum.ActionColors(),
		"match.border-foreground":           gum.ActionColors(),
		"match.foreground":                  gum.ActionColors(),
	})
}
