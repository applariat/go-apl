package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var policyScheduleParams apl.PolicyScheduleParams

// NewPolicySchedulesCommand Creates a cobra command for PolicySchedules
func NewPolicySchedulesCommand() *cobra.Command {

	cmd := createListCommand(cmdListPolicySchedules, "policy-schedules", "")
	getCmd := createGetCommand(cmdGetPolicySchedules, "policy-schedule", "")
	createCmd := createCreateCommand(cmdCreatePolicySchedules, "policy-schedule", "")
	updateCmd := createUpdateCommand(cmdUpdatePolicySchedules, "policy-schedule", "")
	deleteCmd := createDeleteCommand(cmdDeletePolicySchedules, "policy-schedule", "")

	// command flags
	cmd.Flags().StringVar(&policyScheduleParams.Name, "name", "", "Filter policy-schedules by name")
	cmd.Flags().StringVar(&policyScheduleParams.ResourceID, "resource-id", "", "Filter policy-schedules by resource_id")
	cmd.Flags().StringVar(&policyScheduleParams.ResourceType, "resource-type", "", "Filter policy-schedules by resource_type")
	cmd.Flags().StringVar(&policyScheduleParams.PolicyID, "policy-id", "", "Filter policy-schedules by policy_id")
	cmd.Flags().StringVar(&policyScheduleParams.Status, "status", "", "Filter policy-schedules by status")

	// add sub commands
	cmd.AddCommand(getCmd)
	cmd.AddCommand(createCmd)
	cmd.AddCommand(updateCmd)
	cmd.AddCommand(deleteCmd)

	return cmd
}

// cmdListPolicySchedules returns a list of policySchedules
func cmdListPolicySchedules(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runListCommand(&policyScheduleParams, aplSvc.PolicySchedules.List)

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
