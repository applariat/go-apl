package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var locDeployParams apl.LocDeployParams

func NewLocDeploysCommand() *cobra.Command {

	cmd := createListCommand(cmdListLocDeploys, "loc-deploys", "")
	getCmd := createGetCommand(cmdGetLocDeploys, "loc-deploy", "")
	createCmd := createCreateCommand(cmdCreateLocDeploys, "loc-deploy", "")
	deleteCmd := createDeleteCommand(cmdDeleteLocDeploys, "loc-deploy", "")

	// command flags
	cmd.Flags().StringVar(&locDeployParams.Name, "name", "", "Filter loc_deploys by name")
	cmd.Flags().StringVar(&locDeployParams.LocDeploysType, "loc-deploys-type", "", "Filter loc_deploys by loc_deploys_type")
	cmd.Flags().StringVar(&locDeployParams.CredentialID, "credential-id", "", "Filter loc_deploys by credential_id")
	cmd.Flags().StringVar(&locDeployParams.CredentialType, "credential-type", "", "Filter loc_deploys by credential_type")

	// add sub commands
	cmd.AddCommand(getCmd)
	cmd.AddCommand(createCmd)
	cmd.AddCommand(deleteCmd)

	return cmd
}

// cmdListLocDeploys returns a list of locDeploys
func cmdListLocDeploys(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runListCommand(&locDeployParams, aplSvc.LocDeploys.List)

	if output != nil {
		fields := []string{"ID", "Name", "LocDeploysType", "CreatedTime"}
		printTableResultsCustom(output.([]apl.LocDeploy), fields)
	}
}

// cmdGetLocDeploys gets a specified locDeploy by locDeploy-id
func cmdGetLocDeploys(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runGetCommand(args, aplSvc.LocDeploys.Get)

	if output != nil {
		fields := []string{"ID", "Name", "LocDeploysType", "CreatedTime"}
		printTableResultsCustom(output.(apl.LocDeploy), fields)
	}
}

func cmdCreateLocDeploys(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.LocDeployCreateInput{}
	runCreateCommand(in, aplSvs.LocDeploys.Create)
}

func cmdDeleteLocDeploys(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	runDeleteCommand(args, aplSvc.LocDeploys.Delete)
}
