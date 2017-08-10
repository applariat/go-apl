package app

import (
	"fmt"
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

// NewDeploymentsScaleCommand
func NewDeploymentsScaleCommand() *cobra.Command {
	var (
		serviceName      string
		stackComponentID string
		instances        int
	)

	cmd := &cobra.Command{
		Use:   "scale-component [ID]",
		Short: fmt.Sprintf("Scale instances of a component"),
		Long:  "",

		PreRunE: func(cmd *cobra.Command, args []string) error {
			err := checkCommandHasIDInArgs(args, "deployment")
			if err != nil {
				return err
			}
			
			var missingFlags []string

			if stackComponentID == "" {
				missingFlags = append(missingFlags, "--stack-component-id")
			}

			if serviceName == "" {
				missingFlags = append(missingFlags, "--service-name")
			}

			if len(missingFlags) > 0 {
				return fmt.Errorf("Missing required flags: %s", missingFlags)
			}

			return nil
		},

		Run: func(ccmd *cobra.Command, args []string) {
			aplSvc := apl.NewClient()

			in := &apl.DeploymentUpdateInput{
				Command: "override",
				Components: []apl.DeploymentComponent{
					{
						StackComponentID: stackComponentID,
						Services: []apl.Service{
							{
								Name: serviceName,
								Run: apl.Run{
									Instances: instances,
								},
							},
						},
					},
				},
			}

			runUpdateCommand(args, in, aplSvc.Deployments.Update)
		},
	}
	cmd.Flags().IntVar(&instances, "instances", 1, "")
	cmd.Flags().StringVar(&stackComponentID, "stack-component-id", "", "")
	cmd.Flags().StringVar(&serviceName, "service-name", "", "")

	return cmd
}
