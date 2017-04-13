package app

import (
	"github.com/spf13/cobra"
)

var (
	printerType string

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

			err = checkInputFileExists()
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
	AppLariatCmd.AddCommand(NewAuditsCommand())
	AppLariatCmd.AddCommand(NewComponentsCommand())
	AppLariatCmd.AddCommand(NewCredentialsCommand())
	AppLariatCmd.AddCommand(NewDeploymentsCommand())
	AppLariatCmd.AddCommand(NewEventsCommand())
	AppLariatCmd.AddCommand(NewGenerateDocumentationCommand())
	AppLariatCmd.AddCommand(NewLocArtifactsCommand())
	AppLariatCmd.AddCommand(NewLocDeploysCommand())
	AppLariatCmd.AddCommand(NewOrgsCommand())
	AppLariatCmd.AddCommand(NewPoliciesCommand())
	AppLariatCmd.AddCommand(NewPolicyResultsCommand())
	AppLariatCmd.AddCommand(NewPolicySchedulesCommand())
	AppLariatCmd.AddCommand(NewProjectRolesCommand())
	AppLariatCmd.AddCommand(NewProjectsCommand())
	AppLariatCmd.AddCommand(NewReleasesCommand())
	AppLariatCmd.AddCommand(NewRolesCommand())
	AppLariatCmd.AddCommand(NewStackArtifactsCommand())
	AppLariatCmd.AddCommand(NewStackComponentsCommand())
	AppLariatCmd.AddCommand(NewStackVersionsCommand())
	AppLariatCmd.AddCommand(NewStacksCommand())
	AppLariatCmd.AddCommand(NewTypesCommand())
	AppLariatCmd.AddCommand(NewUsersCommand())
	AppLariatCmd.AddCommand(NewWorkloadsCommand())

}
