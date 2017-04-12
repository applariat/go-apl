package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var (
	componentFilterName     string
	componentFilterCategory string

	componentsCmd    = createListCommand(cmdListComponents, "components", "")
	componentsGetCmd = createGetCommand(cmdGetComponents, "component", "")
)

func init() {

	// command flags
	componentsCmd.Flags().StringVar(&componentFilterName, "name", "", "Filter components by category")
	componentsCmd.Flags().StringVar(&componentFilterCategory, "category", "", "Filter components by category")

	// add sub commands
	componentsCmd.AddCommand(componentsGetCmd)

	// Add this command to the main command
	AppLariatCmd.AddCommand(componentsCmd)
}

// cmdListComponents returns a list of components
func cmdListComponents(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	params := &apl.ComponentParams{
		Name:     componentFilterName,
		Category: componentFilterCategory,
	}

	output := runListCommand(params, aplSvc.Components.List)

	if output != nil {
		fields := []string{"ID", "Name", "Category"}
		printTableResultsCustom(output.([]apl.Component), fields)
	}
}

// cmdGetComponents gets a specified component by component-id
func cmdGetComponents(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runGetCommand(args, aplSvc.Components.Get)

	if output != nil {
		fields := []string{"ID", "Name", "Category"}
		printTableResultsCustom(output.(apl.Component), fields)
	}
}
