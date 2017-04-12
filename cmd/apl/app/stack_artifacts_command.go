package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var (
	stackArtifactFilterName          string
	stackArtifactFilterLocArtifactID string
	stackArtifactFilterProjectID     string
	stackArtifactFilterStackID       string
	stackArtifactFilterType          string
	stackArtifactFilterVersion       string
	stackArtifactFilterArtifactName  string
	stackArtifactFilterPackage       string

	stackArtifactsCmd       = createListCommand(cmdListStackArtifacts, "stack-artifacts", "")
	stackArtifactsGetCmd    = createGetCommand(cmdGetStackArtifacts, "stack-artifact", "")
	stackArtifactsCreateCmd = createCreateCommand(cmdCreateStackArtifacts, "stack-artifact", "")
	stackArtifactsUpdateCmd = createUpdateCommand(cmdUpdateStackArtifacts, "stack-artifact", "")
	stackArtifactsDeleteCmd = createDeleteCommand(cmdDeleteStackArtifacts, "stack-artifact", "")
)

func init() {

	// command flags
	stackArtifactsCmd.Flags().StringVar(&stackArtifactFilterName, "name", "", "Filter stack-artifacts by name")
	stackArtifactsCmd.Flags().StringVar(&stackArtifactFilterLocArtifactID, "loc-artifact-id", "", "Filter stack-artifacts by loc_artifact_id")
	stackArtifactsCmd.Flags().StringVar(&stackArtifactFilterProjectID, "project-id", "", "Filter stack-artifacts by project_id")
	stackArtifactsCmd.Flags().StringVar(&stackArtifactFilterStackID, "stack-id", "", "Filter stack-artifacts by stack_id")
	stackArtifactsCmd.Flags().StringVar(&stackArtifactFilterType, "type", "", "Filter stack-artifacts by type")
	stackArtifactsCmd.Flags().StringVar(&stackArtifactFilterVersion, "version", "", "Filter stack-artifacts by version")
	stackArtifactsCmd.Flags().StringVar(&stackArtifactFilterArtifactName, "artifact-name", "", "Filter stack-artifacts by artifact_name")
	stackArtifactsCmd.Flags().StringVar(&stackArtifactFilterPackage, "package", "", "Filter stack-artifacts by package")

	// add sub commands
	stackArtifactsCmd.AddCommand(stackArtifactsGetCmd)
	stackArtifactsCmd.AddCommand(stackArtifactsCreateCmd)
	stackArtifactsCmd.AddCommand(stackArtifactsUpdateCmd)
	stackArtifactsCmd.AddCommand(stackArtifactsDeleteCmd)

	// Add this command to the main command
	AppLariatCmd.AddCommand(stackArtifactsCmd)
}

// cmdListStackArtifacts returns a list of stackArtifacts
func cmdListStackArtifacts(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	params := &apl.StackArtifactParams{
		Name:          stackArtifactFilterName,
		LocArtifactID: stackArtifactFilterLocArtifactID,
		ProjectID:     stackArtifactFilterProjectID,
		StackID:       stackArtifactFilterStackID,
		Type:          stackArtifactFilterType,
		Version:       stackArtifactFilterVersion,
		ArtifactName:  stackArtifactFilterArtifactName,
		Package:       stackArtifactFilterPackage,
	}

	output := runListCommand(params, aplSvc.StackArtifacts.List)

	if output != nil {
		fields := []string{"ID", "Name", "ArtifactName", "Type", "Version", "Package", "CreatedTime"}
		printTableResultsCustom(output.([]apl.StackArtifact), fields)
	}
}

// cmdGetStackArtifacts gets a specified stackArtifact by stackArtifact-id
func cmdGetStackArtifacts(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runGetCommand(args, aplSvc.StackArtifacts.Get)

	if output != nil {
		fields := []string{"ID", "Name", "ArtifactName", "Type", "Version", "Package", "CreatedTime"}
		printTableResultsCustom(output.(apl.StackArtifact), fields)
	}
}

func cmdCreateStackArtifacts(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.StackArtifactCreateInput{}
	runCreateCommand(in, aplSvs.StackArtifacts.Create)
}

func cmdUpdateStackArtifacts(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.StackArtifactUpdateInput{}
	runUpdateCommand(args, in, aplSvs.StackArtifacts.Update)
}

func cmdDeleteStackArtifacts(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	runDeleteCommand(args, aplSvc.StackArtifacts.Delete)
}
