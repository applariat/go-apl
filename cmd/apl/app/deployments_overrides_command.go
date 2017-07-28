package app

import (
	"fmt"
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

// NewDeploymentsOverridesCommand
func NewDeploymentsOverridesCommand() *cobra.Command {

	var (
		componentsMap ComponentStringMap
	)

	cmd := &cobra.Command{
		Use:   "override [ID]",
		Short: fmt.Sprintf("Override a component artifact"),
		Long:  "",

		PreRunE: func(cmd *cobra.Command, args []string) error {
			err := checkCommandHasIDInArgs(args, "credential")
			if err != nil {
				return err
			}
			// If there is a file, no other checking is needed
			if isInputFileDefined() {
				return nil
			}

			var missingFlags []string

			if len(componentsMap.Values) <= 0 {
				missingFlags = append(missingFlags, "--component")
			}

			if len(missingFlags) > 0 {
				return fmt.Errorf("Missing required flags: %s", missingFlags)
			}

			return nil
		},

		Run: func(ccmd *cobra.Command, args []string) {
			aplSvc := apl.NewClient()

			in := &apl.DeploymentUpdateInput{}

			if !isInputFileDefined() {

				// Create the []apl.Components
				c := []apl.DeploymentComponent{}
				for _, cmp := range componentsMap.Values {

					artifact := deploymentArtifactFactory(aplSvc, cmp.StackArtifactIDs)

					dc := apl.DeploymentComponent{
						StackComponentID: cmp.StackComponentID,
						Services: []apl.Service{
							{
								Name: cmp.ServiceName,
								Build: apl.Build{
									Artifact: artifact,
								},
								Run: apl.Run{
									Instances: cmp.Instances,
								},
							},
						},
					}

					c = append(c, dc)

				}

				in = &apl.DeploymentUpdateInput{
					Command:    "override",
					Components: c,
				}

			}

			runUpdateCommand(args, in, aplSvc.Deployments.Update)
		},
	}

	cmd.Flags().Var(&componentsMap, "component", componentsMap.Usage())
	return cmd
}
