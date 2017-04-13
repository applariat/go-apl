package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var auditParams apl.AuditParams

// NewAuditsCommand Creates a cobra command for Audits
func NewAuditsCommand() *cobra.Command {

	cmd := createListCommand(cmdListAudits, "audits", "")
	getCmd := createGetCommand(cmdGetAudits, "audit", "")

	// command flags
	cmd.Flags().StringVar(&auditParams.UserID, "user-id", "", "Filter audits by user-id")
	cmd.Flags().StringVar(&auditParams.ResourceID, "resource-id", "", "Filter audits by resource-id")
	cmd.Flags().StringVar(&auditParams.Action, "action", "", "Filter audits by action")
	cmd.Flags().StringVar(&auditParams.ResourceType, "resource-type", "", "Filter audits by resource-type")

	// add sub commands
	cmd.AddCommand(getCmd)

	return cmd
}

// cmdListAudits returns a list of audits
func cmdListAudits(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runListCommand(&auditParams, aplSvc.Audits.List)

	if output != nil {
		fields := []string{"ID", "Action", "ResourceID", "ResourceType"}
		printTableResultsCustom(output.([]apl.Audit), fields)
	}
}

// cmdGetAudits gets a specified audit by audit-id
func cmdGetAudits(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runGetCommand(args, aplSvc.Audits.Get)

	if output != nil {
		fields := []string{"ID", "Action", "ResourceID", "ResourceType"}
		printTableResultsCustom(output.(apl.Audit), fields)
	}
}
