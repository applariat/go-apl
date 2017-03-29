package apl_test

import (
	"github.com/applariat/go-apl/apl"
	"fmt"
	"testing"
)

var (
	testEventId string
	testEventFilter string
)

func TestEventService_Create(t *testing.T) {

	testEventFilter = "cluster"
	aplSvs := apl.NewClient()

	in := &apl.EventCreateInput{
		ObjectType: testEventFilter,
		ObjectName: "floopy1",
		Message: "this event was created by a test.",
		EventType: "no_action",
		Source: "api",
	}

	out, _, err := aplSvs.Events.Create(in)

	if err != nil {
		t.Fatal(err)
	}

	testEventId = out.Data.(string)
	fmt.Println("New Event ID:", testEventId)

}

func TestEventService_List(t *testing.T) {
	aplSvs := apl.NewClient()

	out, _, err := aplSvs.Events.List(nil)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount > 0 {
		fmt.Printf("Event found %d rows\n", rowCount)
	} else {
		t.Fatal("No Event rows found")
	}

}

func TestEventService_ListByType(t *testing.T) {
	aplSvc := apl.NewClient()

	params := &apl.EventParams{
		ObjectType: testEventFilter,
	}
	out, _, err := aplSvc.Events.List(params)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount == 0 {
		t.Fatal("No Event rows found for filter", testEventFilter)
	}

	fmt.Printf("Event filtered found %d rows for filter \"%s\"\n", rowCount, testEventFilter)

}

func TestEventService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.Events.Get(testEventId)

	if err != nil {
		t.Fatal(err)
	}

	if out == (apl.Event{}) {
		t.Fatal("No Event found for ID", testEventId)
	}

	fmt.Println("Event found for ID", testEventId)

}

