package app

import (
	"fmt"
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var (
	deploymentParams apl.DeploymentParams
)

// NewDeploymentsCommand Creates a cobra command for Deployments
func NewDeploymentsCommand() *cobra.Command {

	cmd := createListCommand(cmdListDeployments, "deployments", "")
	cmd.Flags().StringVar(&deploymentParams.Name, "name", "", "Filter deployments by name")
	cmd.Flags().StringVar(&deploymentParams.StackVersionID, "stack-version-id", "", "Filter deployments by stack_version_id")
	cmd.Flags().StringVar(&deploymentParams.ProjectID, "project-id", "", "Filter deployments by project_id")
	cmd.Flags().StringVar(&deploymentParams.WorkloadID, "workload-id", "", "Filter deployments by workload_id")
	cmd.Flags().StringVar(&deploymentParams.WorkloadType, "workload-type", "", "Filter deployments by workload_type")

	// Get
	getCmd := createGetCommand(cmdGetDeployments, "deployment", "")
	cmd.AddCommand(getCmd)

	// Create
	createCmd := NewDeploymentsCreateCommand()
	cmd.AddCommand(createCmd)

	// Update
	updateCmd := createUpdateCommand(cmdUpdateDeployments, "deployment", "")
	cmd.AddCommand(updateCmd)

	// Delete
	deleteCmd := createDeleteCommand(cmdDeleteDeployments, "deployment", "")
	cmd.AddCommand(deleteCmd)

	// Pods
	//podsCmd := NewDeploymentsPodsCommand()
	//cmd.AddCommand(podsCmd)

	// Override
	overrideCmd := NewDeploymentsOverridesCommand()
	cmd.AddCommand(overrideCmd)

	// Scale Component
	scaleComponentCmd := NewDeploymentsScaleCommand()
	cmd.AddCommand(scaleComponentCmd)

	return cmd
}

// cmdListDeployments returns a list of deployments
func cmdListDeployments(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runListCommand(&deploymentParams, aplSvc.Deployments.List)

	if output != nil {
		fields := []string{"ID", "Name", "CreatedTime"}
		printTableResultsCustom(output.([]apl.Deployment), fields)
	}
}

// cmdGetDeployments gets a specified deployment by deployment-id
func cmdGetDeployments(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runGetCommand(args, aplSvc.Deployments.Get)

	if output != nil {
		fields := []string{"ID", "Name", "CreatedTime"}
		printTableResultsCustom(output.(apl.Deployment), fields)
	}
}

func cmdUpdateDeployments(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	in := &apl.DeploymentUpdateInput{}
	runUpdateCommand(args, in, aplSvc.Deployments.Update)
}

func cmdDeleteDeployments(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	runDeleteCommand(args, aplSvc.Deployments.Delete)
}

// Utility functions

// deploymentArtifactFactory fetches the type and builds the struct
func deploymentArtifactFactory(aplSvc *apl.Client, stackArtifactIDs []string) apl.Artifact {

	var artifact apl.Artifact

	for _, saID := range stackArtifactIDs {
		sa, _, err := aplSvc.StackArtifacts.Get(saID)
		if err != nil {
			panic(err.Error())
		}

		switch sa.StackArtifactType {
		case "code":
			artifact.Code = saID
		case "image":
			artifact.Image = saID
		case "config":
			artifact.Config = saID
		case "data":
			artifact.Data = saID
		default:
			panic(fmt.Errorf("Unsupported StackArtifactType %s", sa.StackArtifactType))
		}

	}
	return artifact
}
