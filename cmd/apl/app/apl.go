package app

import (
	"github.com/spf13/cobra"
)

var (
	printerType string
	inputFile   string

	// AppLariatCmd The appLariat (apl) Command Line Interface is a unified tool to manage your appLariat service.
	// You can control all appLariat services from the command line and automate them through scripts.
	AppLariatCmd = &cobra.Command{
		Use:   "apl",
		Short: "apl",
		Long:  `The appLariat (apl) Command Line Interface is a unified tool to manage your appLariat service. You can control all appLariat services from the command line and automate them through scripts.`,
		PersistentPreRunE: func(ccmd *cobra.Command, args []string) error {

			err := checkPrinterType()
			if err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {

	// persistent flags, globals
	AppLariatCmd.PersistentFlags().StringVarP(&printerType, "output", "o", "table", "Output format: json|yaml")

	// Commands

	// Completely hidden
	//AppLariatCmd.AddCommand(NewAuditsCommand())
	//AppLariatCmd.AddCommand(NewPoliciesCommand())
	//AppLariatCmd.AddCommand(NewPolicyResultsCommand())
	//AppLariatCmd.AddCommand(NewPolicySchedulesCommand())
	//AppLariatCmd.AddCommand(NewProjectRolesCommand())
	//AppLariatCmd.AddCommand(NewCredentialsCommand())
	//AppLariatCmd.AddCommand(NewWorkloadsCommand())
	//AppLariatCmd.AddCommand(NewRolesCommand())
	//AppLariatCmd.AddCommand(NewStackComponentsCommand())
	//AppLariatCmd.AddCommand(NewTypesCommand())

	// Get only
	AppLariatCmd.AddCommand(NewOrgsCommand())
	AppLariatCmd.AddCommand(NewComponentsCommand())
	AppLariatCmd.AddCommand(NewLocArtifactsCommand())
	AppLariatCmd.AddCommand(NewLocDeploysCommand())
	AppLariatCmd.AddCommand(NewProjectsCommand())
	AppLariatCmd.AddCommand(NewUsersCommand())
	AppLariatCmd.AddCommand(NewEventsCommand())

	// More than Get
	AppLariatCmd.AddCommand(NewDeploymentsCommand())
	AppLariatCmd.AddCommand(NewReleasesCommand())
	AppLariatCmd.AddCommand(NewStacksCommand())
	AppLariatCmd.AddCommand(NewStackArtifactsCommand())
	AppLariatCmd.AddCommand(NewStackVersionsCommand())

	// Utility
	AppLariatCmd.AddCommand(NewGenerateDocumentationCommand())

}
