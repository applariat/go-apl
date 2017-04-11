package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var (
	stackFilterName          string
	stackFilterVersionNumber int
	stackFilterReleaseNumber int

	stacksCmd       = createListCommand(cmdListStacks, "stacks", "")
	stacksGetCmd    = createGetCommand(cmdGetStacks, "stack", "")
	stacksCreateCmd = createCreateCommand(cmdCreateStacks, "stack", "")
	stacksUpdateCmd = createUpdateCommand(cmdUpdateStacks, "stack", "")
	stacksDeleteCmd = createDeleteCommand(cmdDeleteStacks, "stack", "")
)

func init() {

	// command flags
	stacksCmd.Flags().StringVar(&stackFilterName, "name", "", "Filter stacks by name")
	stacksCmd.Flags().IntVar(&stackFilterVersionNumber, "version-number", -1, "Filter stacks by version_number")
	stacksCmd.Flags().IntVar(&stackFilterReleaseNumber, "release-number", -1, "Filter stacks by release_number")

	// add sub commands
	stacksCmd.AddCommand(stacksGetCmd)
	stacksCmd.AddCommand(stacksCreateCmd)
	stacksCmd.AddCommand(stacksUpdateCmd)
	stacksCmd.AddCommand(stacksDeleteCmd)

	// Add this command to the main command
	AppLariatCmd.AddCommand(stacksCmd)
}

// cmdListStacks returns a list of stacks
func cmdListStacks(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	params := &apl.StackParams{
		Name: stackFilterName,
	}

	if stackFilterVersionNumber != -1 {
		params.VersionNumber = stackFilterVersionNumber
	}

	if stackFilterReleaseNumber != -1 {
		params.ReleaseNumber = stackFilterReleaseNumber
	}

	output := runListCommand(params, aplSvc.Stacks.List)

	if output != nil {
		fields := []string{"ID", "Name", "VersionNumber", "ReleaseNumber", "CreatedTime"}
		printTableResultsCustom(output.([]apl.Stack), fields)
	}
}

// cmdGetStacks gets a specified stack by stack-id
func cmdGetStacks(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runGetCommand(args, aplSvc.Stacks.Get)

	if output != nil {
		fields := []string{"ID", "Name", "VersionNumber", "ReleaseNumber", "CreatedTime"}
		printTableResultsCustom(output.(apl.Stack), fields)
	}
}

func cmdCreateStacks(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.StackCreateInput{}
	runCreateCommand(in, aplSvs.Stacks.Create)
}

func cmdUpdateStacks(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.StackUpdateInput{}
	runUpdateCommand(args, in, aplSvs.Stacks.Update)
}

func cmdDeleteStacks(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	runDeleteCommand(args, aplSvc.Stacks.Delete)
}
