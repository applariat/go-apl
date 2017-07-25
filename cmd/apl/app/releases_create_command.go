package app

import (
	"fmt"
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

func NewReleasesCreateCommand() *cobra.Command {
	var (
		releaseName           string
		releaseStackID        string
		releaseStackVersionID string
		componentsMap         ComponentStringMap
	)
	cmd := &cobra.Command{
		Use:   "create",
		Short: fmt.Sprintf("Create a release"),
		Long:  "",

		PreRunE: func(cmd *cobra.Command, args []string) error {

			// If there is a file, no other checking is needed
			if isInputFileDefined() {
				return nil
			}

			var missingFlags []string

			if releaseName == "" {
				missingFlags = append(missingFlags, "--name")
			} else {
				// sanitize name, must be dns friendly
				releaseName = subdomainSafe(releaseName)
			}

			if releaseStackID == "" {
				missingFlags = append(missingFlags, "--stack-id")
			}

			if releaseStackVersionID == "" {
				missingFlags = append(missingFlags, "--stack-version-id")
			}

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
			in := &apl.ReleaseCreateInput{}

			if !isInputFileDefined() {

				c := []apl.ReleaseOverrideComponent{}
				for _, cmp := range componentsMap.Values {

					artifact := releaseArtifactFactory(aplSvc, cmp.StackArtifactIDs)

					roc := apl.ReleaseOverrideComponent{
						Name:             cmp.ServiceName,
						StackComponentID: cmp.StackComponentID,
						Services: []apl.ReleaseOverrideService{
							{
								Name: cmp.ServiceName,
								Release: apl.ReleaseOverrideRelease{
									Artifacts: artifact,
								},
							},
						},
					}

					c = append(c, roc)
				}

				in = &apl.ReleaseCreateInput{
					MetaData: apl.MetaData{
						DisplayName: releaseName,
					},
					StackID:        releaseStackID,
					StackVersionID: releaseStackVersionID,
					Components:     c,
				}

			}

			runCreateCommand(in, aplSvc.Releases.Create)
		},
	}

	cmd.Flags().StringVar(&releaseName, "name", "", "")
	cmd.Flags().StringVar(&releaseStackID, "stack-id", "", "")
	cmd.Flags().StringVar(&releaseStackVersionID, "stack-version-id", "", "")
	cmd.Flags().Var(&componentsMap, "component", componentsMap.Usage())

	return cmd

}

// releaseArtifactFactory fetches the type and builds the struct
func releaseArtifactFactory(aplSvc *apl.Client, stackArtifactIDs []string) apl.ReleaseOverrideArtifact {

	artifact := apl.ReleaseOverrideArtifact{}

	for _, saID := range stackArtifactIDs {

		base := &apl.ReleaseOverrideArtifactBase{
			StackArtifactID: saID,
		}

		sa, _, err := aplSvc.StackArtifacts.Get(saID)
		if err != nil {
			panic(err.Error())
		}

		switch sa.StackArtifactType {
		case "code":
			artifact.Code = base
		case "image":
			artifact.Image = base
		case "builder":
			artifact.Builder = base
		default:
			panic(fmt.Errorf("Unsupported StackArtifactType %s", sa.StackArtifactType))
		}

	}
	return artifact
}
