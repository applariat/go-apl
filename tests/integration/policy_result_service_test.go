// +build integration

package tests

import (
	"fmt"
	"github.com/applariat/go-apl/apl"
	"testing"
)

var (
	testPolicyResultId     string
	testPolicyResultFilter = "p-mobile-apps-apl"
)

func TestPolicyResultService_Create(t *testing.T) {

	in := &apl.PolicyResultCreateInput{
		PolicyID:  "po-lease-termination-apl",
		ProjectID: testPolicyResultFilter,
	}

	out, _, err := aplClient.PolicyResults.Create(in)

	if err != nil {
		t.Fatal(err)
	}
	testPolicyResultId = out.Data.(string)
	fmt.Println("New PolicyResult ID:", testPolicyResultId)

}

func TestPolicyResultService_List(t *testing.T) {

	out, _, err := aplClient.PolicyResults.List(nil)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount > 0 {
		fmt.Printf("PolicyResult found %d rows\n", rowCount)
	} else {
		t.Fatal("No PolicyResult rows found")
	}

}

func TestPolicyResultService_ListByType(t *testing.T) {
	aplSvc := apl.NewClient()

	params := &apl.PolicyResultParams{
		ProjectID: testPolicyResultFilter,
	}
	out, _, err := aplSvc.PolicyResults.List(params)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount == 0 {
		t.Fatal("No PolicyResult rows found for filter", testPolicyResultFilter)
	}

	fmt.Printf("PolicyResult filtered found %d rows for filter \"%s\"\n", rowCount, testPolicyResultFilter)

}

func TestPolicyResultService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.PolicyResults.Get(testPolicyResultId)

	if err != nil {
		t.Fatal(err)
	}

	if out == (apl.PolicyResult{}) {
		t.Fatal("No PolicyResult found for ID", testPolicyResultId)
	}

	fmt.Println("PolicyResult found for ID", testPolicyResultId)

}
