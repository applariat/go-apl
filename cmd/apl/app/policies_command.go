package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var (
	policiesFilterName             string
	policiesFilterPolicyType       string
	policiesFilterPolicyGroup      string
	policiesFilterPolicyTemplateID string
	policiesFilterReturn           string

	policiessCmd       = createListCommand(cmdListPoliciess, "policies", "")
	policiessGetCmd    = createGetCommand(cmdGetPoliciess, "policy", "")
	policiessCreateCmd = createCreateCommand(cmdCreatePoliciess, "policy", "")
	policiessUpdateCmd = createUpdateCommand(cmdUpdatePoliciess, "policy", "")
	policiessDeleteCmd = createDeleteCommand(cmdDeletePoliciess, "policy", "")
)

func init() {

	// command flags
	policiessCmd.Flags().StringVar(&policiesFilterName, "name", "", "Filter policiess by name")
	policiessCmd.Flags().StringVar(&policiesFilterPolicyType, "policy-type", "", "Filter policiess by policy_type")
	policiessCmd.Flags().StringVar(&policiesFilterPolicyGroup, "policy-group", "", "Filter policiess by policy_group")
	policiessCmd.Flags().StringVar(&policiesFilterPolicyTemplateID, "policy-template-id", "", "Filter policiess by policy_template_id")
	policiessCmd.Flags().StringVar(&policiesFilterReturn, "return", "", "Filter policiess by return")

	// add sub commands
	policiessCmd.AddCommand(policiessGetCmd)
	policiessCmd.AddCommand(policiessCreateCmd)
	policiessCmd.AddCommand(policiessUpdateCmd)
	policiessCmd.AddCommand(policiessDeleteCmd)

	// Add this command to the main command
	AppLariatCmd.AddCommand(policiessCmd)
}

// cmdListPoliciess returns a list of policiess
func cmdListPoliciess(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	params := &apl.PolicyParams{
		Name:             policiesFilterName,
		PolicyType:       policiesFilterPolicyType,
		PolicyGroup:      policiesFilterPolicyGroup,
		PolicyTemplateID: policiesFilterPolicyTemplateID,
		Return:           policiesFilterReturn,
	}

	output := runListCommand(params, aplSvc.Policies.List)

	if output != nil {
		fields := []string{"ID", "Name", "PolicyType", "PolicyGroup", "CreatedTime"}
		printTableResultsCustom(output.([]apl.Policy), fields)
	}
}

// cmdGetPoliciess gets a specified policies by policies-id
func cmdGetPoliciess(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runGetCommand(args, aplSvc.Policies.Get)

	if output != nil {
		fields := []string{"ID", "Name", "PolicyType", "PolicyGroup", "CreatedTime"}
		printTableResultsCustom(output.(apl.Policy), fields)
	}
}

func cmdCreatePoliciess(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.PolicyCreateInput{}
	runCreateCommand(in, aplSvs.Policies.Create)
}

func cmdUpdatePoliciess(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.PolicyUpdateInput{}
	runUpdateCommand(args, in, aplSvs.Policies.Update)
}

func cmdDeletePoliciess(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	runDeleteCommand(args, aplSvc.Policies.Delete)
}
