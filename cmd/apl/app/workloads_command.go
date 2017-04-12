package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var (
	workloadFilterMaxLeasePeriodDays int
	workloadFilterPriority           int

	workloadFilterName             string
	workloadFilterWorkloadType     string
	workloadFilterLeaseType        string
	workloadFilterQualityOfService string

	workloadsCmd       = createListCommand(cmdListWorkloads, "workloads", "")
	workloadsGetCmd    = createGetCommand(cmdGetWorkloads, "workload", "")
	workloadsUpdateCmd = createUpdateCommand(cmdUpdateWorkloads, "workload", "")
)

func init() {

	// command flags
	workloadsCmd.Flags().StringVar(&workloadFilterName, "name", "", "Filter workloads by name")
	workloadsCmd.Flags().StringVar(&workloadFilterWorkloadType, "workload-type", "", "Filter workloads by workload_type")
	workloadsCmd.Flags().StringVar(&workloadFilterLeaseType, "lease-type", "", "Filter workloads by lease_type")
	workloadsCmd.Flags().StringVar(&workloadFilterQualityOfService, "quality-of-service", "", "Filter workloads by quality_of_service")

	workloadsCmd.Flags().IntVar(&workloadFilterMaxLeasePeriodDays, "max-lease-period-days", -1, "Filter workloads by max_lease_period_days")
	workloadsCmd.Flags().IntVar(&workloadFilterPriority, "priority", -1, "Filter workloads by priority")

	// add sub commands
	workloadsCmd.AddCommand(workloadsGetCmd)
	workloadsCmd.AddCommand(workloadsUpdateCmd)

	// Add this command to the main command
	AppLariatCmd.AddCommand(workloadsCmd)
}

// cmdListWorkloads returns a list of workloads
func cmdListWorkloads(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	params := &apl.WorkloadParams{
		Name:             workloadFilterName,
		WorkloadType:     workloadFilterWorkloadType,
		LeaseType:        workloadFilterLeaseType,
		QualityOfService: workloadFilterQualityOfService,
	}

	if workloadFilterMaxLeasePeriodDays != -1 {
		params.MaxLeasePeriodDays = workloadFilterMaxLeasePeriodDays
	}

	if workloadFilterPriority != -1 {
		params.Priority = workloadFilterPriority
	}

	output := runListCommand(params, aplSvc.Workloads.List)

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
