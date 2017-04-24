package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"os"
)

var genDocsOutputFolder string

// NewGenerateDocumentationCommand Creates a cobra command for GenerateDocumentation
func NewGenerateDocumentationCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "gen-docs",
		Short: "Generate markdown documentation",
		Run: func(cmd *cobra.Command, args []string) {
			if _, err := os.Stat(genDocsOutputFolder); err != nil {
				if os.IsNotExist(err) {
					fmt.Println("Directory not found", genDocsOutputFolder)
					return
				}
			}
			doc.GenMarkdownTree(AppLariatCmd, genDocsOutputFolder)
		},
	}

	cmd.Flags().StringVar(&genDocsOutputFolder, "doc-dir", "docs", "The directory to write docs to")

	return cmd
}
