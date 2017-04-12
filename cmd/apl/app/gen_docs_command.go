package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"os"
)

var (
	outputFolder  string
	genDocCommand = &cobra.Command{
		Use:   "gen-docs",
		Short: "Generate markdown documentation",
		Run: func(cmd *cobra.Command, args []string) {
			if _, err := os.Stat(outputFolder); err != nil {
				if os.IsNotExist(err) {
					fmt.Println("Directory not found", outputFolder)
					return
				}
			}
			doc.GenMarkdownTree(AppLariatCmd, outputFolder)
		},
	}
)

func init() {
	genDocCommand.Flags().StringVar(&outputFolder, "doc-dir", "docs", "The directory to write docs to")
	AppLariatCmd.AddCommand(genDocCommand)
}
