package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var (
	auditFilterUserID       string
	auditFilterResourceID   string
	auditFilterAction       string
	auditFilterResourceType string

	auditsCmd    = createListCommand(cmdListAudits, "audits", "")
	auditsGetCmd = createGetCommand(cmdGetAudits, "audit", "")
)

func init() {

	// command flags
	auditsCmd.Flags().StringVar(&auditFilterUserID, "user-id", "", "Filter audits by user-id")
	auditsCmd.Flags().StringVar(&auditFilterResourceID, "resource-id", "", "Filter audits by resource-id")
	auditsCmd.Flags().StringVar(&auditFilterAction, "action", "", "Filter audits by action")
	auditsCmd.Flags().StringVar(&auditFilterResourceType, "resource-type", "", "Filter audits by resource-type")

	// add sub commands
	auditsCmd.AddCommand(auditsGetCmd)

	// Add this command to the main command
	AppLariatCmd.AddCommand(auditsCmd)
}

// cmdListAudits returns a list of audits
func cmdListAudits(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	params := &apl.AuditParams{
		UserID:       auditFilterUserID,
		ResourceID:   auditFilterResourceID,
		Action:       auditFilterAction,
		ResourceType: auditFilterResourceType,
	}

	output := runListCommand(params, aplSvc.Audits.List)

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
