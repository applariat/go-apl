package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var credentialsParams apl.CredentialParams

func NewCredentialsCommand() *cobra.Command {

	cmd := createListCommand(cmdListCredentials, "credentials", "")
	getCmd := createGetCommand(cmdGetCredentials, "credential", "")
	createCmd := createCreateCommand(cmdCreateCredentials, "credential", "")
	updateCmd := createUpdateCommand(cmdUpdateCredentials, "credential", "")
	deleteCmd := createDeleteCommand(cmdDeleteCredentials, "credential", "")

	// command flags
	cmd.Flags().StringVar(&credentialsParams.Name, "name", "", "Filter credentials by name")
	cmd.Flags().StringVar(&credentialsParams.CredentialType, "credential-type", "", "Filter credentials by type")

	// add sub commands
	cmd.AddCommand(getCmd)
	cmd.AddCommand(createCmd)
	cmd.AddCommand(updateCmd)
	cmd.AddCommand(deleteCmd)

	return cmd
}

// cmdListCredentials returns a list of credentials
func cmdListCredentials(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runListCommand(&credentialsParams, aplSvc.Credentials.List)

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
