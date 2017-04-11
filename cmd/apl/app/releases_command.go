package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
)

var (
	releaseFilterName           string
	releaseFilterVersion        string
	releaseFilterStackID        string
	releaseFilterStackVersionID string
	releaseFilterProjectID      string
	releaseFilterLocImageID     string
	releaseFilterBuildStatus    string

	releasesCmd       = createListCommand(cmdListReleases, "releases", "")
	releasesGetCmd    = createGetCommand(cmdGetReleases, "release", "")
	releasesCreateCmd = createCreateCommand(cmdCreateReleases, "release", "")
	releasesDeleteCmd = createDeleteCommand(cmdDeleteReleases, "release", "")
)

func init() {

	// command flags
	releasesCmd.Flags().StringVar(&releaseFilterName, "name", "", "Filter releases by name")
	releasesCmd.Flags().StringVar(&releaseFilterVersion, "version", "", "Filter releases by version")
	releasesCmd.Flags().StringVar(&releaseFilterStackID, "stack-id", "", "Filter releases by stack_id")
	releasesCmd.Flags().StringVar(&releaseFilterStackVersionID, "stack-version-id", "", "Filter releases by stack_version_id")
	releasesCmd.Flags().StringVar(&releaseFilterProjectID, "project-id", "", "Filter releases by project_id")
	releasesCmd.Flags().StringVar(&releaseFilterLocImageID, "loc-image-id", "", "Filter releases by loc_image_id")
	releasesCmd.Flags().StringVar(&releaseFilterBuildStatus, "build-status", "", "Filter releases by build_status")

	// add sub commands
	releasesCmd.AddCommand(releasesGetCmd)
	releasesCmd.AddCommand(releasesCreateCmd)
	releasesCmd.AddCommand(releasesDeleteCmd)

	// Add this command to the main command
	AppLariatCmd.AddCommand(releasesCmd)
}

// cmdListReleases returns a list of releases
func cmdListReleases(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	params := &apl.ReleaseParams{
		Name:           releaseFilterName,
		Version:        releaseFilterVersion,
		StackID:        releaseFilterStackID,
		StackVersionID: releaseFilterStackVersionID,
		ProjectID:      releaseFilterProjectID,
		LocImageID:     releaseFilterLocImageID,
		BuildStatus:    releaseFilterBuildStatus,
	}

	output := runListCommand(params, aplSvc.Releases.List)

	if output != nil {
		fields := []string{"ID", "Version", "BuildStatus", "CreatedTime"}
		printTableResultsCustom(output.([]apl.Release), fields)
	}
}

// cmdGetReleases gets a specified release by release-id
func cmdGetReleases(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runGetCommand(args, aplSvc.Releases.Get)

	if output != nil {
		fields := []string{"ID", "Version", "BuildStatus", "CreatedTime"}
		printTableResultsCustom(output.(apl.Release), fields)
	}
}

func cmdCreateReleases(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.ReleaseCreateInput{}
	runCreateCommand(in, aplSvs.Releases.Create)
}

func cmdDeleteReleases(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()
	runDeleteCommand(args, aplSvc.Releases.Delete)
}
