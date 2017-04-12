package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var (
	policyResultFilterPolicyID         string
	policyResultFilterPolicyScheduleID string
	policyResultFilterProjectID        string

	policyResultsCmd       = createListCommand(cmdListPolicyResults, "policy-results", "")
	policyResultsGetCmd    = createGetCommand(cmdGetPolicyResults, "policy-result", "")
	policyResultsCreateCmd = createCreateCommand(cmdCreatePolicyResults, "policy-result", "")
)

func init() {

	// command flags
	policyResultsCmd.Flags().StringVar(&policyResultFilterPolicyID, "policy-id", "", "Filter policy-results by policy_id")
	policyResultsCmd.Flags().StringVar(&policyResultFilterPolicyScheduleID, "policy-schedule-id", "", "Filter policy-results by policy_schedule_id")
	policyResultsCmd.Flags().StringVar(&policyResultFilterProjectID, "project-id", "", "Filter policy-results by project_id")

	// add sub commands
	policyResultsCmd.AddCommand(policyResultsGetCmd)
	policyResultsCmd.AddCommand(policyResultsCreateCmd)

	// Add this command to the main command
	AppLariatCmd.AddCommand(policyResultsCmd)
}

// cmdListPolicyResults returns a list of policyResults
func cmdListPolicyResults(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	params := &apl.PolicyResultParams{
		PolicyID:         policyResultFilterPolicyID,
		PolicyScheduleID: policyResultFilterPolicyScheduleID,
		ProjectID:        policyResultFilterProjectID,
	}

	output := runListCommand(params, aplSvc.PolicyResults.List)

	if output != nil {
		fields := []string{"ID", "PolicyID", "PolicyScheduleID", "ProjectID", "CreatedTime"}
		printTableResultsCustom(output.([]apl.PolicyResult), fields)
	}
}

// cmdGetPolicyResults gets a specified policyResult by policyResult-id
func cmdGetPolicyResults(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runGetCommand(args, aplSvc.PolicyResults.Get)

	if output != nil {
		fields := []string{"ID", "PolicyID", "PolicyScheduleID", "ProjectID", "CreatedTime"}
		printTableResultsCustom(output.(apl.PolicyResult), fields)
	}
}

func cmdCreatePolicyResults(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.PolicyResultCreateInput{}
	runCreateCommand(in, aplSvs.PolicyResults.Create)
}
