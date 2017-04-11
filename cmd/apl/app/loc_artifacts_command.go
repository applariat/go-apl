package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var (
	locArtifactFilterName               string
	locArtifactFilterLocArtifactsType   string
	locArtifactFilterBucket             string
	locArtifactFilterCredentialID       string
	locArtifactFilterCredentialType     string
	locArtifactFilterSecretCredentialID string
	locArtifactFilterRegistryURI        string
	locArtifactFilterURL                string

	locArtifactCmd       = createListCommand(cmdListLocArtifacts, "loc-artifacts", "")
	locArtifactGetCmd    = createGetCommand(cmdGetLocArtifacts, "loc-artifact", "")
	locArtifactCreateCmd = createCreateCommand(cmdCreateLocArtifacts, "loc_artifact", "")
	locArtifactUpdateCmd = createUpdateCommand(cmdUpdateLocArtifacts, "loc_artifact", "")
	locArtifactDeleteCmd = createDeleteCommand(cmdDeleteLocArtifacts, "loc_artifact", "")
)

func init() {

	// command flags
	locArtifactCmd.Flags().StringVar(&locArtifactFilterName, "name", "", "Filter loc_artifacts by name")
	locArtifactCmd.Flags().StringVar(&locArtifactFilterLocArtifactsType, "loc-artifacts-type", "", "Filter loc_artifacts by loc_artifacts_type")
	locArtifactCmd.Flags().StringVar(&locArtifactFilterBucket, "bucket", "", "Filter loc_artifacts by bucket")
	locArtifactCmd.Flags().StringVar(&locArtifactFilterCredentialID, "credential-id", "", "Filter loc_artifacts by credential_id")
	locArtifactCmd.Flags().StringVar(&locArtifactFilterCredentialType, "credential-type", "", "Filter loc_artifacts by credential_type")
	locArtifactCmd.Flags().StringVar(&locArtifactFilterSecretCredentialID, "secret-credential-id", "", "Filter loc_artifacts by secret_credential_id")
	locArtifactCmd.Flags().StringVar(&locArtifactFilterRegistryURI, "registry-uri", "", "Filter loc_artifacts by registry_uri")
	locArtifactCmd.Flags().StringVar(&locArtifactFilterURL, "url", "", "Filter loc_artifacts by url")

	// add sub commands
	locArtifactCmd.AddCommand(locArtifactGetCmd)
	locArtifactCmd.AddCommand(locArtifactCreateCmd)
	locArtifactCmd.AddCommand(locArtifactUpdateCmd)
	locArtifactCmd.AddCommand(locArtifactDeleteCmd)

	// Add this command to the main command
	AppLariatCmd.AddCommand(locArtifactCmd)
}

// cmdListLocArtifacts returns a list of loc_artifacts
func cmdListLocArtifacts(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	params := &apl.LocArtifactParams{
		Name:               locArtifactFilterName,
		LocArtifactsType:   locArtifactFilterLocArtifactsType,
		Bucket:             locArtifactFilterBucket,
		CredentialID:       locArtifactFilterCredentialID,
		CredentialType:     locArtifactFilterCredentialType,
		SecretCredentialID: locArtifactFilterSecretCredentialID,
		RegistryURI:        locArtifactFilterRegistryURI,
		URL:                locArtifactFilterURL,
	}

	output := runListCommand(params, aplSvc.LocArtifacts.List)

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
