package app

import (
	"github.com/applariat/go-apl/pkg/apl"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	eventFilterName       string
	eventFilterObjectType string
	eventFilterSource     string
	eventFilterActive     string

	eventsCmd       = createListCommand(cmdListEvents, "events", "")
	eventsGetCmd    = createGetCommand(cmdGetEvents, "event", "")
	eventsCreateCmd = createCreateCommand(cmdCreateEvents, "event", "")
)

func init() {

	// command flags
	eventsCmd.Flags().StringVar(&eventFilterObjectType, "object-type", "", "Filter events by object_type")
	eventsCmd.Flags().StringVar(&eventFilterSource, "source", "", "Filter events by source")
	eventsCmd.Flags().StringVar(&eventFilterName, "object-name", "", "Filter events by object_name")

	// add sub commands
	eventsCmd.AddCommand(eventsGetCmd)
	eventsCmd.AddCommand(eventsCreateCmd)

	// Add this command to the main command
	AppLariatCmd.AddCommand(eventsCmd)
}

// cmdListEvents returns a list of events
func cmdListEvents(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	params := &apl.EventParams{
		ObjectName: eventFilterName,
		ObjectType: eventFilterObjectType,
		Source:     eventFilterSource,
	}

	if eventFilterActive != "" {
		if efa, err := strconv.ParseBool(eventFilterActive); err == nil {
			params.Active = efa
		} else {
			ccmd.Usage()
			return
		}

	}

	output := runListCommand(params, aplSvc.Events.List)

	if output != nil {
		fields := []string{"ID", "ObjectName", "ObjectType", "Active"}
		printTableResultsCustom(output.([]apl.Event), fields)
	}
}

// cmdGetEvents gets a specified event by event-id
func cmdGetEvents(ccmd *cobra.Command, args []string) {
	aplSvc := apl.NewClient()

	output := runGetCommand(args, aplSvc.Events.Get)

	if output != nil {
		fields := []string{"ID", "ObjectName", "ObjectType", "Active"}
		printTableResultsCustom(output.(apl.Event), fields)
	}
}

func cmdCreateEvents(ccmd *cobra.Command, args []string) {
	aplSvs := apl.NewClient()
	in := &apl.EventCreateInput{}
	runCreateCommand(in, aplSvs.Events.Create)
}
