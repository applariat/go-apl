// +build integration

package tests

import (
	"fmt"
	"github.com/applariat/go-apl/pkg/apl"
	"testing"
)

var (
	testPolicyId     = "policy-test-id"
	testPolicyFilter = "scheduled"
)

func TestPolicyService_Create(t *testing.T) {

	in := &apl.PolicyCreateInput{
		ID:         testPolicyId,
		Name:       "policy test name",
		PolicyType: testPolicyFilter,

		PolicyGroup:      "cluster_management",
		PolicyTemplateID: "pt-auto-scale-cluster",
		Return:           "OP1",
		Inputs:           "{}",
		Assets:           "{}",
		Actions:          "{}",
		Operations:       "{}",
		Attributes:       "{}",
		Constants:        "{}",
	}

	out, _, err := aplClient.Policies.Create(in)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("New Policy ID:", out.Data)

}

func TestPolicyService_List(t *testing.T) {

	out, _, err := aplClient.Policies.List(nil)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount > 0 {
		fmt.Printf("Policy found %d rows\n", rowCount)
	} else {
		t.Fatal("No Policy rows found")
	}

}

func TestPolicyService_ListByType(t *testing.T) {
	aplSvc := apl.NewClient()

	params := &apl.PolicyParams{
		PolicyType: testPolicyFilter,
	}
	out, _, err := aplSvc.Policies.List(params)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount == 0 {
		t.Fatal("No Policy rows found for filter", testPolicyFilter)
	}

	fmt.Printf("Policy filtered found %d rows for filter \"%s\"\n", rowCount, testPolicyFilter)

}

func TestPolicyService_Update(t *testing.T) {
	aplSvc := apl.NewClient()

	in := &apl.PolicyUpdateInput{Name: "name UPDATED!"}
	out, _, err := aplSvc.Policies.Update(testPolicyId, in)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Skipped:", out.Skipped)
	fmt.Println("Deleted:", out.Deleted)
	fmt.Println("Unchanged:", out.Unchanged)
	fmt.Println("Replaced:", out.Replaced)
	fmt.Println("Errors:", out.Errors)
}

func TestPolicyService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.Policies.Get(testPolicyId)

	if err != nil {
		t.Fatal(err)
	}

	if out == (apl.Policy{}) {
		t.Fatal("No Policy found for ID", testPolicyId)
	}

	fmt.Println("Policy found for ID", testPolicyId)

}

func TestPolicyService_Delete(t *testing.T) {
	aplSvc := apl.NewClient()
	out, _, err := aplSvc.Policies.Delete(testPolicyId)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Skipped:", out.Skipped)
	fmt.Println("Deleted:", out.Deleted)
	fmt.Println("Unchanged:", out.Unchanged)
	fmt.Println("Replaced:", out.Replaced)
	fmt.Println("Errors:", out.Errors)
}
