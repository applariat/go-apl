package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var orgFilterName string

// NewOrgsCommand Creates a cobra command for Orgs
func NewOrgsCommand() *cobra.Command {

	cmd := createListCommand(cmdListOrgs, "orgs", "")
	getCmd := createGetCommand(cmdGetOrgs, "org", "")
	updateCmd := createUpdateCommand(cmdUpdateOrgs, "org", "")

	// command flags
	cmd.Flags().StringVar(&orgFilterName, "name", "", "Filter orgs by name")

	// add sub commands
	cmd.AddCommand(getCmd)
	cmd.AddCommand(updateCmd)

	return cmd
}

// cmdListOrgs returns a list of orgs
func cmdListOrgs(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runListCommand(nil, aplSvc.Orgs.List)

	if output != nil {
		fields := []string{"ID", "CompanyName", "OrgType", "NumOfEmployees"}
		printTableResultsCustom(output.(apl.Org), fields)
	}
}

// cmdGetOrgs gets a specified org by org-id
func cmdGetOrgs(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runGetCommand(args, aplSvc.Orgs.Get)

	if output != nil {
		fields := []string{"ID", "CompanyName", "OrgType", "NumOfEmployees"}
		printTableResultsCustom(output.(apl.Org), fields)
	}
}

func cmdUpdateOrgs(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.OrgUpdateInput{}
	runUpdateCommand(args, in, aplSvs.Orgs.Update)
}
