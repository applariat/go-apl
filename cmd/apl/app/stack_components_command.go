package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var stackComponentParams apl.StackComponentParams

// NewStackComponentsCommand Creates a cobra command for StackComponents
func NewStackComponentsCommand() *cobra.Command {

	cmd := createListCommand(cmdListStackComponents, "stack-components", "")
	getCmd := createGetCommand(cmdGetStackComponents, "stack-component", "")
	createCmd := createCreateCommand(cmdCreateStackComponents, "stack-component", "")
	updateCmd := createUpdateCommand(cmdUpdateStackComponents, "stack-component", "")
	deleteCmd := createDeleteCommand(cmdDeleteStackComponents, "stack-component", "")

	// command flags
	cmd.Flags().StringVar(&stackComponentParams.Name, "name", "", "Filter stack-components by name")
	cmd.Flags().StringVar(&stackComponentParams.ComponentID, "component-id", "", "Filter stack-components by component_id")
	cmd.Flags().StringVar(&stackComponentParams.ComponentVersionID, "component-version-id", "", "Filter stack-components by component_version_id")
	cmd.Flags().StringVar(&stackComponentParams.StackID, "stack-id", "", "Filter stack-components by stack_id")
	cmd.Flags().StringVar(&stackComponentParams.StackVersionID, "stack-version-id", "", "Filter stack-components by stack_version_id")
	cmd.Flags().StringVar(&stackComponentParams.ProjectID, "project-id", "", "Filter stack-components by project_id")

	// add sub commands
	cmd.AddCommand(getCmd)
	cmd.AddCommand(createCmd)
	cmd.AddCommand(updateCmd)
	cmd.AddCommand(deleteCmd)

	return cmd
}

// cmdListStackComponents returns a list of stackComponents
func cmdListStackComponents(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	output := runListCommand(&stackComponentParams, aplSvc.StackComponents.List)
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
