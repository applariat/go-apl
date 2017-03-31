// +build integration

package tests

import (
	"fmt"
	"github.com/applariat/go-apl/apl"
	"testing"
)

var (
	testPolicyScheduleId     = "ps-lease-term-sched-apl"
	testPolicyScheduleFilter = "loc_deploy"
)

func TestPolicyScheduleService_Create(t *testing.T) {

	// TODO: Fix TestPolicyScheduleService_Create!
	t.SkipNow()

	in := &apl.PolicyScheduleCreateInput{
		ID:           testPolicyScheduleId,
		Name:         "policy-schedule-test-name",
		PolicyID:     "po-lease-termination-apl",
		ResourceID:   "ld-gke-deploy",
		ResourceType: testPolicyScheduleFilter,
		Schedule: map[string]interface{}{
			"crontab":            "0 0 02 * * ?",
			"iterations":         1,
			"schedule_frequency": "day",
			"schedule_value":     2,
			"timezone":           "America/Los_Angeles",
			"value_type":         "hour_of_day",
		},
		Inputs: map[string]string{"lease_type": "temporary"},
	}

	out, _, err := aplClient.PolicySchedules.Create(in)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("New PolicySchedule ID:", out.Data)

}

func TestPolicyScheduleService_List(t *testing.T) {

	out, _, err := aplClient.PolicySchedules.List(nil)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount > 0 {
		fmt.Printf("PolicySchedule found %d rows\n", rowCount)
	} else {
		fmt.Printf("No PolicySchedule rows found")
	}

}

func TestPolicyScheduleService_ListByType(t *testing.T) {
	aplSvc := apl.NewClient()

	params := &apl.PolicyScheduleParams{
		ResourceType: testPolicyScheduleFilter,
	}
	out, _, err := aplSvc.PolicySchedules.List(params)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount == 0 {
		fmt.Printf("No PolicySchedule rows found for filter", testPolicyScheduleFilter)
	}

	fmt.Printf("PolicySchedule filtered found %d rows for filter \"%s\"\n", rowCount, testPolicyScheduleFilter)

}

func TestPolicyScheduleService_Update(t *testing.T) {

	// TODO: Fix TestPolicyScheduleService_Update!
	t.SkipNow()

	aplSvc := apl.NewClient()

	in := &apl.PolicyScheduleUpdateInput{Name: "gke-cluster-lease-termination-schedule"}
	out, _, err := aplSvc.PolicySchedules.Update(testPolicyScheduleId, in)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Skipped:", out.Skipped)
	fmt.Println("Deleted:", out.Deleted)
	fmt.Println("Unchanged:", out.Unchanged)
	fmt.Println("Replaced:", out.Replaced)
	fmt.Println("Errors:", out.Errors)
}

func TestPolicyScheduleService_Get(t *testing.T) {

	// TODO: Fix TestPolicyScheduleService_Get!
	t.SkipNow()

	aplSvc := apl.NewClient()

	out, _, err := aplSvc.PolicySchedules.Get(testPolicyScheduleId)

	if err != nil {
		t.Fatal(err)
	}

	if out == (apl.PolicySchedule{}) {
		t.Fatal("No PolicySchedule found for ID", testPolicyScheduleId)
	}

	fmt.Println("PolicySchedule found for ID", testPolicyScheduleId)

}

func TestPolicyScheduleService_Delete(t *testing.T) {

	// TODO: Fix TestPolicyScheduleService_Delete!
	t.SkipNow()

	aplSvc := apl.NewClient()
	out, _, err := aplSvc.PolicySchedules.Delete(testPolicyScheduleId)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Skipped:", out.Skipped)
	fmt.Println("Deleted:", out.Deleted)
	fmt.Println("Unchanged:", out.Unchanged)
	fmt.Println("Replaced:", out.Replaced)
	fmt.Println("Errors:", out.Errors)
}
