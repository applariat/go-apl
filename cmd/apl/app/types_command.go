package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

func NewTypesCommand() *cobra.Command {
	cmd := createListCommand(cmdListTypes, "types", "")
	getCmd := createGetCommand(cmdGetTypes, "type", "")
	cmd.AddCommand(getCmd)
	return cmd
}

// cmdListTypes returns a list of types
func cmdListTypes(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	output := runListCommand(nil, aplSvc.Types.List)
	if output != nil {
		fields := []string{"ID", "Types"}
		printTableResultsCustom(output.([]apl.Type), fields)
	}
}

// cmdGetTypes gets a specified type by type-id
func cmdGetTypes(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	output := runGetCommand(args, aplSvc.Types.Get)
	if output != nil {
		fields := []string{"ID", "Types"}
		printTableResultsCustom(output.(apl.Type), fields)
	}
}
