package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var (
	policyScheduleFilterName        string
	policyScheduleFilteResourceID   string
	policyScheduleFilteResourceType string
	policyScheduleFiltePolicyID     string
	policyScheduleFilteStatus       string

	policySchedulesCmd       = createListCommand(cmdListPolicySchedules, "policy-schedules", "")
	policySchedulesGetCmd    = createGetCommand(cmdGetPolicySchedules, "policy-schedule", "")
	policySchedulesCreateCmd = createCreateCommand(cmdCreatePolicySchedules, "policy-schedule", "")
	policySchedulesUpdateCmd = createUpdateCommand(cmdUpdatePolicySchedules, "policy-schedule", "")
	policySchedulesDeleteCmd = createDeleteCommand(cmdDeletePolicySchedules, "policy-schedule", "")
)

func init() {

	// command flags
	policySchedulesCmd.Flags().StringVar(&policyScheduleFilterName, "name", "", "Filter policy-schedules by name")
	policySchedulesCmd.Flags().StringVar(&policyScheduleFilteResourceID, "resource-id", "", "Filter policy-schedules by resource_id")
	policySchedulesCmd.Flags().StringVar(&policyScheduleFilteResourceType, "resource-type", "", "Filter policy-schedules by resource_type")
	policySchedulesCmd.Flags().StringVar(&policyScheduleFiltePolicyID, "policy-id", "", "Filter policy-schedules by policy_id")
	policySchedulesCmd.Flags().StringVar(&policyScheduleFilteStatus, "status", "", "Filter policy-schedules by status")

	// add sub commands
	policySchedulesCmd.AddCommand(policySchedulesGetCmd)
	policySchedulesCmd.AddCommand(policySchedulesCreateCmd)
	policySchedulesCmd.AddCommand(policySchedulesUpdateCmd)
	policySchedulesCmd.AddCommand(policySchedulesDeleteCmd)

	// Add this command to the main command
	AppLariatCmd.AddCommand(policySchedulesCmd)
}

// cmdListPolicySchedules returns a list of policySchedules
func cmdListPolicySchedules(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	params := &apl.PolicyScheduleParams{
		Name:         policyScheduleFilterName,
		ResourceID:   policyScheduleFilteResourceID,
		ResourceType: policyScheduleFilteResourceType,
		PolicyID:     policyScheduleFiltePolicyID,
		Status:       policyScheduleFilteStatus,
	}

	output := runListCommand(params, aplSvc.PolicySchedules.List)

	if output != nil {
		fields := []string{"ID", "Name", "ResourceType", "Status", "CreatedTime"}
		printTableResultsCustom(output.([]apl.PolicySchedule), fields)
	}
}

// cmdGetPolicySchedules gets a specified policySchedule by policySchedule-id
func cmdGetPolicySchedules(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runGetCommand(args, aplSvc.PolicySchedules.Get)

	if output != nil {
		fields := []string{"ID", "Name", "ResourceType", "Status", "CreatedTime"}
		printTableResultsCustom(output.(apl.PolicySchedule), fields)
	}
}

func cmdCreatePolicySchedules(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.PolicyScheduleCreateInput{}
	runCreateCommand(in, aplSvs.PolicySchedules.Create)
}

func cmdUpdatePolicySchedules(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.PolicyScheduleUpdateInput{}
	runUpdateCommand(args, in, aplSvs.PolicySchedules.Update)
}

func cmdDeletePolicySchedules(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	runDeleteCommand(args, aplSvc.PolicySchedules.Delete)
}
