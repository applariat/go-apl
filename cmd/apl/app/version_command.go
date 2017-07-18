package app

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	// VERSION is set during build
	VERSION string
)

// NewComponentsCommand Creates a cobra command for Components
func NewVersionCommand() *cobra.Command {

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Show applariat.io client version",
		Run: func(cmd *cobra.Command, args []string) {
			if VERSION == "" {
				VERSION = "unknown"
			}
			fmt.Println(AppLariatCmd.Use + " version " + VERSION)
		},
	}

	return versionCmd

}
