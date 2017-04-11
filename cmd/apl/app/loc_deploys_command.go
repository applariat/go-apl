package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var (
	locDeployFilterName           string
	locDeployFilterLocDeploysType string
	locDeployFilterCredentialID   string
	locDeployFilterCredentialType string

	locDeploysCmd       = createListCommand(cmdListLocDeploys, "loc-deploys", "")
	locDeploysGetCmd    = createGetCommand(cmdGetLocDeploys, "loc-deploy", "")
	locDeploysCreateCmd = createCreateCommand(cmdCreateLocDeploys, "loc-deploy", "")
	locDeploysDeleteCmd = createDeleteCommand(cmdDeleteLocDeploys, "loc-deploy", "")
)

func init() {

	// command flags
	locDeploysCmd.Flags().StringVar(&locDeployFilterName, "name", "", "Filter loc_deploys by name")
	locDeploysCmd.Flags().StringVar(&locDeployFilterLocDeploysType, "loc-deploys-type", "", "Filter loc_deploys by loc_deploys_type")
	locDeploysCmd.Flags().StringVar(&locDeployFilterCredentialID, "credential-id", "", "Filter loc_deploys by credential_id")
	locDeploysCmd.Flags().StringVar(&locDeployFilterCredentialType, "credential-type", "", "Filter loc_deploys by credential_type")

	// add sub commands
	locDeploysCmd.AddCommand(locDeploysGetCmd)
	locDeploysCmd.AddCommand(locDeploysCreateCmd)
	locDeploysCmd.AddCommand(locDeploysDeleteCmd)

	// Add this command to the main command
	AppLariatCmd.AddCommand(locDeploysCmd)
}

// cmdListLocDeploys returns a list of locDeploys
func cmdListLocDeploys(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	params := &apl.LocDeployParams{
		Name:           locDeployFilterName,
		LocDeploysType: locDeployFilterLocDeploysType,
		CredentialID:   locDeployFilterCredentialID,
		CredentialType: locDeployFilterCredentialType,
	}

	output := runListCommand(params, aplSvc.LocDeploys.List)

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
