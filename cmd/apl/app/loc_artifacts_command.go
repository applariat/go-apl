package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var locArtifactParams apl.LocArtifactParams

func NewLocArtifactsCommand() *cobra.Command {

	cmd := createListCommand(cmdListLocArtifacts, "loc-artifacts", "")
	getCmd := createGetCommand(cmdGetLocArtifacts, "loc-artifact", "")
	createCmd := createCreateCommand(cmdCreateLocArtifacts, "loc_artifact", "")
	updateCmd := createUpdateCommand(cmdUpdateLocArtifacts, "loc_artifact", "")
	deleteCmd := createDeleteCommand(cmdDeleteLocArtifacts, "loc_artifact", "")

	// command flags
	cmd.Flags().StringVar(&locArtifactParams.Name, "name", "", "Filter loc_artifacts by name")
	cmd.Flags().StringVar(&locArtifactParams.LocArtifactsType, "loc-artifacts-type", "", "Filter loc_artifacts by loc_artifacts_type")
	cmd.Flags().StringVar(&locArtifactParams.Bucket, "bucket", "", "Filter loc_artifacts by bucket")
	cmd.Flags().StringVar(&locArtifactParams.CredentialID, "credential-id", "", "Filter loc_artifacts by credential_id")
	cmd.Flags().StringVar(&locArtifactParams.CredentialType, "credential-type", "", "Filter loc_artifacts by credential_type")
	cmd.Flags().StringVar(&locArtifactParams.SecretCredentialID, "secret-credential-id", "", "Filter loc_artifacts by secret_credential_id")
	cmd.Flags().StringVar(&locArtifactParams.RegistryURI, "registry-uri", "", "Filter loc_artifacts by registry_uri")
	cmd.Flags().StringVar(&locArtifactParams.URL, "url", "", "Filter loc_artifacts by url")

	// add sub commands
	cmd.AddCommand(getCmd)
	cmd.AddCommand(createCmd)
	cmd.AddCommand(updateCmd)
	cmd.AddCommand(deleteCmd)

	return cmd
}

// cmdListLocArtifacts returns a list of loc_artifacts
func cmdListLocArtifacts(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runListCommand(&locArtifactParams, aplSvc.LocArtifacts.List)

	if output != nil {
		fields := []string{"ID", "Name", "LocArtifactsType", "CreatedTime"}
		printTableResultsCustom(output.([]apl.LocArtifact), fields)
	}
}

// cmdGetLocArtifacts gets a specified loc_artifact by loc_artifact-id
func cmdGetLocArtifacts(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runGetCommand(args, aplSvc.LocArtifacts.Get)

	if output != nil {
		fields := []string{"ID", "Name", "LocArtifactsType", "CreatedTime"}
		printTableResultsCustom(output.(apl.LocArtifact), fields)
	}
}

func cmdCreateLocArtifacts(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.LocArtifactCreateInput{}
	runCreateCommand(in, aplSvs.LocArtifacts.Create)
}

func cmdUpdateLocArtifacts(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.LocArtifactUpdateInput{}
	runUpdateCommand(args, in, aplSvs.LocArtifacts.Update)
}

func cmdDeleteLocArtifacts(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	runDeleteCommand(args, aplSvc.LocArtifacts.Delete)
}
