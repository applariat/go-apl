package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var workloadParams apl.WorkloadParams

func NewWorkloadsCommand() *cobra.Command {
	cmd := createListCommand(cmdListWorkloads, "workloads", "")
	getCmd := createGetCommand(cmdGetWorkloads, "workload", "")
	updateCmd := createUpdateCommand(cmdUpdateWorkloads, "workload", "")

	// command flags
	cmd.Flags().StringVar(&workloadParams.Name, "name", "", "Filter workloads by name")
	cmd.Flags().StringVar(&workloadParams.WorkloadType, "workload-type", "", "Filter workloads by workload_type")
	cmd.Flags().StringVar(&workloadParams.LeaseType, "lease-type", "", "Filter workloads by lease_type")
	cmd.Flags().StringVar(&workloadParams.QualityOfService, "quality-of-service", "", "Filter workloads by quality_of_service")

	// add sub commands
	cmd.AddCommand(getCmd)
	cmd.AddCommand(updateCmd)

	return cmd
}

// cmdListWorkloads returns a list of workloads
func cmdListWorkloads(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	output := runListCommand(&workloadParams, aplSvc.Workloads.List)
	if output != nil {
		fields := []string{"ID", "WorkloadType", "LeaseType", "MaxLeasePeriodDays", "Priority", "QualityOfService", "CreatedTime"}
		printTableResultsCustom(output.([]apl.Workload), fields)
	}
}

// cmdGetWorkloads gets a specified workload by workload-id
func cmdGetWorkloads(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	output := runGetCommand(args, aplSvc.Workloads.Get)
	if output != nil {
		fields := []string{"ID", "WorkloadType", "LeaseType", "MaxLeasePeriodDays", "Priority", "QualityOfService", "CreatedTime"}
		printTableResultsCustom(output.(apl.Workload), fields)
	}
}

func cmdUpdateWorkloads(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.WorkloadUpdateInput{}
	runUpdateCommand(args, in, aplSvs.Workloads.Update)
}
