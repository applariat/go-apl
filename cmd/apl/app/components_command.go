package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var componentParams apl.ComponentParams

// NewComponentsCommand Creates a cobra command for Components
func NewComponentsCommand() *cobra.Command {

	cmd := createListCommand(cmdListComponents, "components", "")
	getCmd := createGetCommand(cmdGetComponents, "component", "")

	// command flags
	cmd.Flags().StringVar(&componentParams.Name, "name", "", "Filter components by category")
	cmd.Flags().StringVar(&componentParams.Category, "category", "", "Filter components by category")

	// add sub commands
	cmd.AddCommand(getCmd)

	return cmd

}

// cmdListComponents returns a list of components
func cmdListComponents(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runListCommand(&componentParams, aplSvc.Components.List)

	if output != nil {
		fields := []string{"ID", "Name", "Category", "Versions"}
		printTableResultsCustom(output.([]apl.Component), fields)
	}
}

// cmdGetComponents gets a specified component by component-id
func cmdGetComponents(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runGetCommand(args, aplSvc.Components.Get)

	if output != nil {
		fields := []string{"ID", "Name", "Category", "Versions"}
		printTableResultsCustom(output.(apl.Component), fields)
	}
}
