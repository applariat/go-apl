package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var policyResultParams apl.PolicyResultParams

// NewPolicyResultsCommand Creates a cobra command for PolicyResults
func NewPolicyResultsCommand() *cobra.Command {

	cmd := createListCommand(cmdListPolicyResults, "policy-results", "")
	getCmd := createGetCommand(cmdGetPolicyResults, "policy-result", "")
	createCmd := createCreateCommand(cmdCreatePolicyResults, "policy-result", "")

	// command flags
	cmd.Flags().StringVar(&policyResultParams.PolicyID, "policy-id", "", "Filter policy-results by policy_id")
	cmd.Flags().StringVar(&policyResultParams.PolicyScheduleID, "policy-schedule-id", "", "Filter policy-results by policy_schedule_id")
	cmd.Flags().StringVar(&policyResultParams.ProjectID, "project-id", "", "Filter policy-results by project_id")

	// add sub commands
	cmd.AddCommand(getCmd)
	cmd.AddCommand(createCmd)

	return cmd
}

// cmdListPolicyResults returns a list of policyResults
func cmdListPolicyResults(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runListCommand(&policyResultParams, aplSvc.PolicyResults.List)

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
