package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var (
	releaseParams           apl.ReleaseParams
	releaseName             string
	releaseStackID          string
	releaseStackVersionID   string
	releaseStackComponentID string
	releaseServiceName      string
	releaseStackArtifactID  string
)

// NewReleasesCommand Creates a cobra command for Releases
func NewReleasesCommand() *cobra.Command {

	cmd := createListCommand(cmdListReleases, "releases", "")
	cmd.Flags().StringVar(&releaseParams.Name, "name", "", "Filter releases by name")
	cmd.Flags().StringVar(&releaseParams.Version, "version", "", "Filter releases by version")
	cmd.Flags().StringVar(&releaseParams.StackID, "stack-id", "", "Filter releases by stack_id")
	cmd.Flags().StringVar(&releaseParams.StackVersionID, "stack-version-id", "", "Filter releases by stack_version_id")
	cmd.Flags().StringVar(&releaseParams.ProjectID, "project-id", "", "Filter releases by project_id")
	cmd.Flags().StringVar(&releaseParams.LocImageID, "loc-image-id", "", "Filter releases by loc_image_id")
	cmd.Flags().StringVar(&releaseParams.BuildStatus, "build-status", "", "Filter releases by build_status")

	// Get
	getCmd := createGetCommand(cmdGetReleases, "release", "")
	cmd.AddCommand(getCmd)

	// Create
	createCmd := createCreateCommand(cmdCreateReleases, "release", "")
	//createCmd := &cobra.Command{
	//	Use:   "create",
	//	Short: fmt.Sprintf("Create a release"),
	//	Long:  "",
	//	Run:   cmdCreateReleases,
	//	PreRunE: func(cmd *cobra.Command, args []string) error {
	//		var missingFlags []string
	//
	//		if releaseName == "" {
	//			missingFlags = append(missingFlags, "--name")
	//		} else {
	//			// sanitize name, must be dns friendly
	//			releaseName = subdomainSafe(releaseName)
	//		}
	//
	//		if releaseStackID == "" {
	//			missingFlags = append(missingFlags, "--stack-id")
	//		}
	//
	//		if releaseStackVersionID == "" {
	//			missingFlags = append(missingFlags, "--stack-version-id")
	//		}
	//
	//		if releaseStackComponentID == "" {
	//			missingFlags = append(missingFlags, "--stack-component-id")
	//		}
	//
	//		if releaseServiceName == "" {
	//			missingFlags = append(missingFlags, "--service-name")
	//		}
	//
	//		if releaseStackArtifactID == "" {
	//			missingFlags = append(missingFlags, "--stack-artifact-id")
	//		}
	//
	//		if len(missingFlags) > 0 {
	//			return fmt.Errorf("Missing required flags: %s", missingFlags)
	//		}
	//
	//		return nil
	//	},
	//}
	//
	//createCmd.Flags().StringVar(&releaseName, "name", "", "")
	//createCmd.Flags().StringVar(&releaseStackID, "stack-id", "", "")
	//createCmd.Flags().StringVar(&releaseStackVersionID, "stack-version-id", "", "")
	//createCmd.Flags().StringVar(&releaseStackComponentID, "stack-component-id", "", "")
	//createCmd.Flags().StringVar(&releaseServiceName, "service-name", "", "")
	//createCmd.Flags().StringVar(&releaseStackArtifactID, "stack-artifact-id", "", "")
	cmd.AddCommand(createCmd)

	// Delete
	deleteCmd := createDeleteCommand(cmdDeleteReleases, "release", "")
	cmd.AddCommand(deleteCmd)

	return cmd
}

// cmdListReleases returns a list of releases
func cmdListReleases(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	output := runListCommand(&releaseParams, aplSvc.Releases.List)
	if output != nil {
		fields := []string{"ID", "StackID", "Version", "CreatedTime"}
		printTableResultsCustom(output.([]apl.Release), fields)
	}
}

// cmdGetReleases gets a specified release by release-id
func cmdGetReleases(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	output := runGetCommand(args, aplSvc.Releases.Get)
	if output != nil {
		fields := []string{"ID", "StackID", "Version", "CreatedTime"}
		printTableResultsCustom(output.(apl.Release), fields)
	}
}

func cmdCreateReleases(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	in := &apl.ReleaseCreateInput{}

	//if !isInputFileDefined() {
	//	//artifact := artifactFactory(aplSvc, releaseStackArtifactID)
	//
	//	components := []map[string]interface{} {
	//			StackComponentID string                    `json:"stack_component_id,omitempty"`
	//			Name             string                    `json:"name,omitempty"`
	//			Services         []struct {
	//				Name             string `json:"name,omitempty"`
	//				ReleaseArtifacts `json:"release,omitempty"`
	//			}, `json:"services,omitempty"`
	//	}
	//
	//	in = &apl.ReleaseCreateInput{
	//		MetaData: &apl.MetaData{
	//			DisplayName: releaseName,
	//		},
	//		StackID:        releaseStackID,
	//		StackVersionID: releaseStackVersionID,
	//		//Components: []apl.ReleaseComponent{
	//		//	{
	//		//		StackComponentID: releaseStackComponentID,
	//		//		Services: []apl.ReleaseComponentService{
	//		//			{
	//		//				Name: releaseServiceName,
	//		//				ReleaseArtifacts: apl.ReleaseArtifacts{
	//		//					Artifact: apl.ReleaseArtifacts{
	//		//						Code:
	//		//					}
	//		//					Artifact: artifact,
	//		//				},
	//		//			},
	//		//		},
	//		//	},
	//		//},
	//	}
	//}

	//printYAML(in)

	runCreateCommand(in, aplSvc.Releases.Create)
}

func cmdDeleteReleases(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	runDeleteCommand(args, aplSvc.Releases.Delete)
}
