package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var stackParams apl.StackParams

// NewStacksCommand Creates a cobra command for Stacks
func NewStacksCommand() *cobra.Command {

	cmd := createListCommand(cmdListStacks, "stacks", "")
	getCmd := createGetCommand(cmdGetStacks, "stack", "")
	createCmd := createCreateCommand(cmdCreateStacks, "stack", "")
	updateCmd := createUpdateCommand(cmdUpdateStacks, "stack", "")
	deleteCmd := createDeleteCommand(cmdDeleteStacks, "stack", "")

	// command flags
	cmd.Flags().StringVar(&stackParams.Name, "name", "", "Filter stacks by name")

	// add sub commands
	cmd.AddCommand(getCmd)
	cmd.AddCommand(createCmd)
	cmd.AddCommand(updateCmd)
	cmd.AddCommand(deleteCmd)

	return cmd
}

// cmdListStacks returns a list of stacks
func cmdListStacks(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	output := runListCommand(&stackParams, aplSvc.Stacks.List)
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
