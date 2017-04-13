package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

func NewStackVersionsCommand() *cobra.Command {

	cmd := createListCommand(cmdListStackVersions, "stack-versions", "")
	getCmd := createGetCommand(cmdGetStackVersions, "stack-version", "")
	createCmd := createCreateCommand(cmdCreateStackVersions, "stack-version", "")
	deleteCmd := createDeleteCommand(cmdDeleteStackVersions, "stack-version", "")

	// add sub commands
	cmd.AddCommand(getCmd)
	cmd.AddCommand(createCmd)
	cmd.AddCommand(deleteCmd)

	return cmd
}

// cmdListStackVersions returns a list of stackVersions
func cmdListStackVersions(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	output := runListCommand(nil, aplSvc.StackVersions.List)
	if output != nil {
		header := []string{"Stack ID", "Stack Version ID"}
		// Print results out here ourselved due to the funky response.

		out := output.([]apl.StackVersionList)
		result := make([][]string, len(out))

		for _, row := range out {
			for _, ver := range row.StackVersions {
				result = append(result, []string{row.StackID, ver.ID})
			}
		}
		printTableResults(result, header)
	}
}

// cmdGetStackVersions gets a specified stackVersion by stackVersion-id
func cmdGetStackVersions(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	output := runGetCommand(args, aplSvc.StackVersions.Get)
	if output != nil {
		fields := []string{"ID", "Name", "CreatedTime"}
		printTableResultsCustom(output.(apl.StackVersion), fields)
	}
}

func cmdCreateStackVersions(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.StackVersionCreateInput{}
	runCreateCommand(in, aplSvs.StackVersions.Create)
}

func cmdDeleteStackVersions(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	runDeleteCommand(args, aplSvc.StackVersions.Delete)
}
