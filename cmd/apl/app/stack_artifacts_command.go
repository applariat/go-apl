package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var (
	stackArtifactParams      apl.StackArtifactParams
	stackArtifactCreateInput apl.StackArtifactCreateInput
)

// NewStackArtifactsCommand Creates a cobra command for StackArtifacts
func NewStackArtifactsCommand() *cobra.Command {

	// command flags
	cmd := createListCommand(cmdListStackArtifacts, "stack-artifacts", "")
	cmd.Flags().StringVar(&stackArtifactParams.Name, "name", "", "Filter stack-artifacts by name")
	cmd.Flags().StringVar(&stackArtifactParams.LocArtifactID, "loc-artifact-id", "", "Filter stack-artifacts by loc_artifact_id")
	cmd.Flags().StringVar(&stackArtifactParams.ProjectID, "project-id", "", "Filter stack-artifacts by project_id")
	cmd.Flags().StringVar(&stackArtifactParams.StackID, "stack-id", "", "Filter stack-artifacts by stack_id")
	cmd.Flags().StringVar(&stackArtifactParams.StackArtifactType, "type", "", "Filter stack-artifacts by [code|config|image|data|builder]")
	cmd.Flags().StringVar(&stackArtifactParams.Version, "version", "", "Filter stack-artifacts by version")
	cmd.Flags().StringVar(&stackArtifactParams.ArtifactName, "artifact-name", "", "Filter stack-artifacts by artifact_name")
	cmd.Flags().StringVar(&stackArtifactParams.Package, "package", "", "Filter stack-artifacts by package")

	// Get
	getCmd := createGetCommand(cmdGetStackArtifacts, "stack-artifact", "")
	cmd.AddCommand(getCmd)

	// Create
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "create a stack-artifact",
		Run:   cmdCreateStackArtifacts,
	}
	addInputFileFlag(createCmd)
	createCmd.Flags().StringVar(&stackArtifactCreateInput.ID, "id", "", "")
	createCmd.Flags().StringVar(&stackArtifactCreateInput.LocArtifactID, "loc-artifact-id", "", "")
	createCmd.Flags().StringVar(&stackArtifactCreateInput.StackID, "stack-id", "", "")
	createCmd.Flags().StringVar(&stackArtifactCreateInput.ArtifactName, "artifact-name", "", "")
	createCmd.Flags().StringVar(&stackArtifactCreateInput.Name, "name", "", "")
	createCmd.Flags().StringVar(&stackArtifactCreateInput.StackArtifactType, "stack-artifact-type", "", "[code|config|image|data|builder]")
	cmd.AddCommand(createCmd)

	// Update
	updateCmd := createUpdateCommand(cmdUpdateStackArtifacts, "stack-artifact", "")
	cmd.AddCommand(updateCmd)

	// Delete
	deleteCmd := createDeleteCommand(cmdDeleteStackArtifacts, "stack-artifact", "")
	cmd.AddCommand(deleteCmd)

	return cmd
}

// cmdListStackArtifacts returns a list of stackArtifacts
func cmdListStackArtifacts(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	output := runListCommand(&stackArtifactParams, aplSvc.StackArtifacts.List)
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
	if !isInputFileDefined() {
		in = &stackArtifactCreateInput
	}
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
