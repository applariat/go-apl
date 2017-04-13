package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var deploymentParams apl.DeploymentParams

func NewDeploymentsCommand() *cobra.Command {

	cmd := createListCommand(cmdListDeployments, "deployments", "")
	getCmd := createGetCommand(cmdGetDeployments, "deployment", "")
	createCmd := createCreateCommand(cmdCreateDeployments, "deployment", "")
	updateCmd := createUpdateCommand(cmdUpdateDeployments, "deployment", "")
	deleteCmd := createDeleteCommand(cmdDeleteDeployments, "deployment", "")

	// command flags
	//cmd.Flags().IntVar(&deploymentParams.ReleaseVersion, "release-version", 0, "Filter deployments by release_version")

	cmd.Flags().StringVar(&deploymentParams.Name, "name", "", "Filter deployments by name")
	cmd.Flags().StringVar(&deploymentParams.StackVersionID, "stack-version-id", "", "Filter deployments by stack_version_id")
	cmd.Flags().StringVar(&deploymentParams.ProjectID, "project-id", "", "Filter deployments by project_id")
	cmd.Flags().StringVar(&deploymentParams.WorkloadID, "workload-id", "", "Filter deployments by workload_id")
	cmd.Flags().StringVar(&deploymentParams.LeaseType, "lease-type", "", "Filter deployments by lease_type")
	cmd.Flags().StringVar(&deploymentParams.WorkloadName, "workload-name", "", "Filter deployments by workload_name")
	cmd.Flags().StringVar(&deploymentParams.LeaseExpiration, "lease-expiration", "", "Filter deployments by lease_expiration")
	cmd.Flags().StringVar(&deploymentParams.QosLevel, "qos-level", "", "Filter deployments by qos_level")

	// add sub commands
	cmd.AddCommand(getCmd)
	cmd.AddCommand(createCmd)
	cmd.AddCommand(updateCmd)
	cmd.AddCommand(deleteCmd)

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

func cmdCreateDeployments(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.DeploymentCreateInput{}
	runCreateCommand(in, aplSvs.Deployments.Create)
}

func cmdUpdateDeployments(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.DeploymentUpdateInput{}
	runUpdateCommand(args, in, aplSvs.Deployments.Update)
}

func cmdDeleteDeployments(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	runDeleteCommand(args, aplSvc.Deployments.Delete)
}
