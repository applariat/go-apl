package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var (
	deploymentFilterReleaseVersion  int
	deploymentFilterName            string
	deploymentFilterStackVersionID  string
	deploymentFilterProjectID       string
	deploymentFilterWorkloadID      string
	deploymentFilterLeaseType       string
	deploymentFilterWorkloadName    string
	deploymentFilterLeaseExpiration string
	deploymentFilterQosLevel        string

	deploymentsCmd       = createListCommand(cmdListDeployments, "deployments", "")
	deploymentsGetCmd    = createGetCommand(cmdGetDeployments, "deployment", "")
	deploymentsCreateCmd = createCreateCommand(cmdCreateDeployments, "deployment", "")
	deploymentsUpdateCmd = createUpdateCommand(cmdUpdateDeployments, "deployment", "")
	deploymentsDeleteCmd = createDeleteCommand(cmdDeleteDeployments, "deployment", "")
)

func init() {

	// command flags
	deploymentsCmd.Flags().IntVar(&deploymentFilterReleaseVersion, "release-version", -1, "Filter deployments by release_version")

	deploymentsCmd.Flags().StringVar(&deploymentFilterName, "name", "", "Filter deployments by name")
	deploymentsCmd.Flags().StringVar(&deploymentFilterStackVersionID, "stack-version-id", "", "Filter deployments by stack_version_id")
	deploymentsCmd.Flags().StringVar(&deploymentFilterProjectID, "project-id", "", "Filter deployments by project_id")
	deploymentsCmd.Flags().StringVar(&deploymentFilterWorkloadID, "workload-id", "", "Filter deployments by workload_id")
	deploymentsCmd.Flags().StringVar(&deploymentFilterLeaseType, "lease-type", "", "Filter deployments by lease_type")
	deploymentsCmd.Flags().StringVar(&deploymentFilterWorkloadName, "workload-name", "", "Filter deployments by workload_name")
	deploymentsCmd.Flags().StringVar(&deploymentFilterLeaseExpiration, "lease-expiration", "", "Filter deployments by lease_expiration")
	deploymentsCmd.Flags().StringVar(&deploymentFilterQosLevel, "qos-level", "", "Filter deployments by qos_level")

	// add sub commands
	deploymentsCmd.AddCommand(deploymentsGetCmd)
	deploymentsCmd.AddCommand(deploymentsCreateCmd)
	deploymentsCmd.AddCommand(deploymentsUpdateCmd)
	deploymentsCmd.AddCommand(deploymentsDeleteCmd)

	// Add this command to the main command
	AppLariatCmd.AddCommand(deploymentsCmd)
}

// cmdListDeployments returns a list of deployments
func cmdListDeployments(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	params := &apl.DeploymentParams{
		Name:            deploymentFilterName,
		StackVersionID:  deploymentFilterStackVersionID,
		ProjectID:       deploymentFilterProjectID,
		WorkloadID:      deploymentFilterWorkloadID,
		LeaseType:       deploymentFilterLeaseType,
		WorkloadName:    deploymentFilterWorkloadName,
		LeaseExpiration: deploymentFilterLeaseExpiration,
		QosLevel:        deploymentFilterQosLevel,
	}

	if deploymentFilterReleaseVersion != -1 {
		params.ReleaseVersion = deploymentFilterReleaseVersion
	}

	output := runListCommand(params, aplSvc.Deployments.List)

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
