package app

import (
	"fmt"
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

// NewDeploymentsCreateCommand
func NewDeploymentsCreateCommand() *cobra.Command {

	var (
		name         string
		releaseID    string
		locDeployID  string
		workloadType string
		//instances     int
		componentsMap ComponentStringMap
	)

	cmd := &cobra.Command{
		Use:   "create",
		Short: fmt.Sprintf("Create a deployment"),
		Long:  "",
		PreRunE: func(cmd *cobra.Command, args []string) error {

			// If there is a file, no other checking is needed
			if isInputFileDefined() {
				return nil
			}

			var missingFlags []string

			if name == "" {
				missingFlags = append(missingFlags, "--name")
			} else {
				// sanitize name, must be dns friendly
				name = subdomainSafe(name)
			}

			if releaseID == "" {
				missingFlags = append(missingFlags, "--release-id")
			}

			if locDeployID == "" {
				missingFlags = append(missingFlags, "--loc-deploy-id")
			}

			// Should this be required? I don't think so.
			//if len(componentsMap.Values) < 1 {
			//	missingFlags = append(missingFlags, "--component")
			//}

			if len(missingFlags) > 0 {
				return fmt.Errorf("Missing required flags: %s", missingFlags)
			}

			return nil
		},

		Run: func(ccmd *cobra.Command, args []string) {
			aplSvc := apl.NewClient()

			// this function will use the file or command line args for input.
			in := &apl.DeploymentCreateInput{}

			if !isInputFileDefined() {

				// Create the []apl.Components
				c := []apl.DeploymentComponent{}
				for _, cmp := range componentsMap.Values {

					// build artifact
					artifact := deploymentArtifactFactory(aplSvc, cmp.StackArtifactIDs)

					dc := apl.DeploymentComponent{
						StackComponentID: cmp.StackComponentID,
						Services: []apl.Service{
							{
								Name: cmp.ServiceName,
								Overrides: apl.Overrides{
									Build: apl.Build{
										Artifact: artifact,
									},
								},
								//Run: apl.Run{
								//	Instances: cmp.Instances,
								//},
							},
						},
					}
					c = append(c, dc)

				}

				in = &apl.DeploymentCreateInput{
					Name:        name,
					LocDeployID: locDeployID,
					ReleaseID:   releaseID,
					Components:  c,
				}

				if workloadType != "" {
					in.WorkloadType = workloadType
				}

			}

			runCreateCommand(in, aplSvc.Deployments.Create)
		},
	}
	addInputFileFlag(cmd)
	cmd.Flags().StringVar(&name, "name", "", "")
	cmd.Flags().StringVar(&workloadType, "workload-type", "", "")
	cmd.Flags().StringVar(&releaseID, "release-id", "", "")
	cmd.Flags().StringVar(&locDeployID, "loc-deploy-id", "", "")
	cmd.Flags().Var(&componentsMap, "component", componentsMap.Usage())

	//cmd.Flags().IntVar(&instances, "instances", 1, "")

	return cmd
}
