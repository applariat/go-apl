package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var (
	stackComponentFilterName               string
	stackComponentFilterComponentID        string
	stackComponentFilterComponentVersionID string
	stackComponentFilterStackID            string
	stackComponentFilterStackVersionID     string
	stackComponentFilterProjectID          string

	stackComponentsCmd       = createListCommand(cmdListStackComponents, "stack-components", "")
	stackComponentsGetCmd    = createGetCommand(cmdGetStackComponents, "stack-component", "")
	stackComponentsCreateCmd = createCreateCommand(cmdCreateStackComponents, "stack-component", "")
	stackComponentsUpdateCmd = createUpdateCommand(cmdUpdateStackComponents, "stack-component", "")
	stackComponentsDeleteCmd = createDeleteCommand(cmdDeleteStackComponents, "stack-component", "")
)

func init() {

	// command flags
	stackComponentsCmd.Flags().StringVar(&stackComponentFilterName, "name", "", "Filter stack-components by name")
	stackComponentsCmd.Flags().StringVar(&stackComponentFilterComponentID, "component-id", "", "Filter stack-components by component_id")
	stackComponentsCmd.Flags().StringVar(&stackComponentFilterComponentVersionID, "component-version-id", "", "Filter stack-components by component_version_id")
	stackComponentsCmd.Flags().StringVar(&stackComponentFilterStackID, "stack-id", "", "Filter stack-components by stack_id")
	stackComponentsCmd.Flags().StringVar(&stackComponentFilterStackVersionID, "stack-version-id", "", "Filter stack-components by stack_version_id")
	stackComponentsCmd.Flags().StringVar(&stackComponentFilterProjectID, "project-id", "", "Filter stack-components by project_id")

	// add sub commands
	stackComponentsCmd.AddCommand(stackComponentsGetCmd)
	stackComponentsCmd.AddCommand(stackComponentsCreateCmd)
	stackComponentsCmd.AddCommand(stackComponentsUpdateCmd)
	stackComponentsCmd.AddCommand(stackComponentsDeleteCmd)

	// Add this command to the main command
	AppLariatCmd.AddCommand(stackComponentsCmd)
}

// cmdListStackComponents returns a list of stackComponents
func cmdListStackComponents(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	params := &apl.StackComponentParams{
		Name:               stackComponentFilterName,
		ComponentID:        stackComponentFilterComponentID,
		ComponentVersionID: stackComponentFilterComponentVersionID,
		StackID:            stackComponentFilterStackID,
		StackVersionID:     stackComponentFilterStackVersionID,
		ProjectID:          stackComponentFilterProjectID,
	}

	output := runListCommand(params, aplSvc.StackComponents.List)

	if output != nil {
		fields := []string{"ID", "Name", "ComponentID", "ComponentVersionID", "CreatedTime"}

		printTableResultsCustom(output.([]apl.StackComponent), fields)
	}
}

// cmdGetStackComponents gets a specified stackComponent by stackComponent-id
func cmdGetStackComponents(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runGetCommand(args, aplSvc.StackComponents.Get)

	if output != nil {
		fields := []string{"ID", "Name", "ComponentID", "ComponentVersionID", "CreatedTime"}
		printTableResultsCustom(output.(apl.StackComponent), fields)
	}
}

func cmdCreateStackComponents(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.StackComponentCreateInput{}
	runCreateCommand(in, aplSvs.StackComponents.Create)
}

func cmdUpdateStackComponents(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.StackComponentUpdateInput{}
	runUpdateCommand(args, in, aplSvs.StackComponents.Update)
}

func cmdDeleteStackComponents(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	runDeleteCommand(args, aplSvc.StackComponents.Delete)
}
