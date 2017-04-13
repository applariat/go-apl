package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var policyParams apl.PolicyParams

func NewPoliciesCommand() *cobra.Command {

	cmd := createListCommand(cmdListPolicies, "policies", "")
	getCmd := createGetCommand(cmdGetPolicies, "policy", "")
	createCmd := createCreateCommand(cmdCreatePolicies, "policy", "")
	updateCmd := createUpdateCommand(cmdUpdatePolicies, "policy", "")
	deleteCmd := createDeleteCommand(cmdDeletePolicies, "policy", "")

	// command flags
	cmd.Flags().StringVar(&policyParams.Name, "name", "", "Filter policiess by name")
	cmd.Flags().StringVar(&policyParams.PolicyType, "policy-type", "", "Filter policiess by policy_type")
	cmd.Flags().StringVar(&policyParams.PolicyGroup, "policy-group", "", "Filter policiess by policy_group")
	cmd.Flags().StringVar(&policyParams.PolicyTemplateID, "policy-template-id", "", "Filter policiess by policy_template_id")
	cmd.Flags().StringVar(&policyParams.Return, "return", "", "Filter policiess by return")

	// add sub commands
	cmd.AddCommand(getCmd)
	cmd.AddCommand(createCmd)
	cmd.AddCommand(updateCmd)
	cmd.AddCommand(deleteCmd)

	return cmd
}

// cmdListPolicies returns a list of policies
func cmdListPolicies(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runListCommand(&policyParams, aplSvc.Policies.List)

	if output != nil {
		fields := []string{"ID", "Name", "PolicyType", "PolicyGroup", "CreatedTime"}
		printTableResultsCustom(output.([]apl.Policy), fields)
	}
}

// cmdGetPolicies gets a specified policies by policies-id
func cmdGetPolicies(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runGetCommand(args, aplSvc.Policies.Get)

	if output != nil {
		fields := []string{"ID", "Name", "PolicyType", "PolicyGroup", "CreatedTime"}
		printTableResultsCustom(output.(apl.Policy), fields)
	}
}

func cmdCreatePolicies(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.PolicyCreateInput{}
	runCreateCommand(in, aplSvs.Policies.Create)
}

func cmdUpdatePolicies(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.PolicyUpdateInput{}
	runUpdateCommand(args, in, aplSvs.Policies.Update)
}

func cmdDeletePolicies(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	runDeleteCommand(args, aplSvc.Policies.Delete)
}
