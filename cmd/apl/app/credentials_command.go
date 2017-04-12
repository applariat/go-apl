package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var (
	credentialFilterName string
	credentialFilterType string

	credentialsCmd       = createListCommand(cmdListCredentials, "credentials", "")
	credentialsGetCmd    = createGetCommand(cmdGetCredentials, "credential", "")
	credentialsCreateCmd = createCreateCommand(cmdCreateCredentials, "credential", "")
	credentialsUpdateCmd = createUpdateCommand(cmdUpdateCredentials, "credential", "")
	credentialsDeleteCmd = createDeleteCommand(cmdDeleteCredentials, "credential", "")
)

func init() {

	// command flags
	credentialsCmd.Flags().StringVar(&credentialFilterName, "name", "", "Filter credentials by name")
	credentialsCmd.Flags().StringVar(&credentialFilterType, "credential-type", "", "Filter credentials by type")

	// add sub commands
	credentialsCmd.AddCommand(credentialsGetCmd)
	credentialsCmd.AddCommand(credentialsCreateCmd)
	credentialsCmd.AddCommand(credentialsUpdateCmd)
	credentialsCmd.AddCommand(credentialsDeleteCmd)

	// Add this command to the main command
	AppLariatCmd.AddCommand(credentialsCmd)
}

// cmdListCredentials returns a list of credentials
func cmdListCredentials(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	params := &apl.CredentialParams{
		Name:           credentialFilterName,
		CredentialType: credentialFilterType,
	}

	output := runListCommand(params, aplSvc.Credentials.List)

	if output != nil {
		fields := []string{"ID", "Name", "CredentialType"}
		printTableResultsCustom(output.([]apl.Credential), fields)
	}
}

// cmdGetCredentials gets a specified credential by credential-id
func cmdGetCredentials(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runGetCommand(args, aplSvc.Credentials.Get)

	if output != nil {
		fields := []string{"ID", "Name", "CredentialType"}
		printTableResultsCustom(output.(apl.Credential), fields)
	}
}

func cmdCreateCredentials(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.CredentialCreateInput{}
	runCreateCommand(in, aplSvs.Credentials.Create)
}

func cmdUpdateCredentials(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.CredentialUpdateInput{}
	runUpdateCommand(args, in, aplSvs.Credentials.Update)
}

func cmdDeleteCredentials(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	runDeleteCommand(args, aplSvc.Credentials.Delete)
}
