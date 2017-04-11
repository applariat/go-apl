package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var (
	//stackVersionFilterName string

	stackVersionsCmd       = createListCommand(cmdListStackVersions, "stack-versions", "")
	stackVersionsGetCmd    = createGetCommand(cmdGetStackVersions, "stack-version", "")
	stackVersionsCreateCmd = createCreateCommand(cmdCreateStackVersions, "stack-version", "")
	stackVersionsDeleteCmd = createDeleteCommand(cmdDeleteStackVersions, "stack-version", "")
)

func init() {

	// command flags
	//stackVersionsCmd.Flags().StringVar(&stackVersionFilterName, "name", "", "Filter stack-versions by name")

	// add sub commands
	stackVersionsCmd.AddCommand(stackVersionsGetCmd)
	stackVersionsCmd.AddCommand(stackVersionsCreateCmd)
	stackVersionsCmd.AddCommand(stackVersionsDeleteCmd)

	// Add this command to the main command
	AppLariatCmd.AddCommand(stackVersionsCmd)
}

// cmdListStackVersions returns a list of stackVersions
func cmdListStackVersions(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	//params := &apl.StackVersionParams{
	//	Name: stackVersionFilterName,
	//}

	output := runListCommand(nil, aplSvc.StackVersions.List)

	if output != nil {
		//fields := []string{"StackID", "StackVersions.StackVersion.ID"}
		//printTableResultsCustom(output.([]apl.StackVersionList), fields)

		header := []string{"Stack ID", "Stack Version ID"}
		// Print results out here ourselved due to the funky response.

		out := output.([]apl.StackVersionList)

		result := make([][]string, len(out))

		for _, row := range out {
			//fmt.Println("id count", len(row.StackVersions))
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
