// +build integration

package tests

import (
	"fmt"
	"github.com/applariat/go-apl/pkg/apl"
	"testing"
)

var (
	testDeploymentId     = "deployment-test-id"
	testDeploymentFilter = "deployment"
)

func TestDeploymentService_Create(t *testing.T) {

	// TODO: Fix TestDeploymentService_Create!
	t.SkipNow()

	in := &apl.DeploymentCreateInput{
		ID:   testDeploymentId,
		Name: "Deployment Test",
	}

	out, _, err := aplClient.Deployments.Create(in)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("New Deployment ID:", out.Data)

}

func TestDeploymentService_List(t *testing.T) {

	out, _, err := aplClient.Deployments.List(nil)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount > 0 {
		fmt.Printf("Deployment found %d rows\n", rowCount)
	} else {
		t.Fatal("No Deployment rows found")
	}

}

func TestDeploymentService_ListByType(t *testing.T) {

	// TODO: Fix TestDeploymentService_ListByType!
	t.SkipNow()

	aplSvc := apl.NewClient()

	params := &apl.DeploymentParams{
		Name: "deployment",
	}
	out, _, err := aplSvc.Deployments.List(params)

	if err != nil {
		t.Fatal(err)
	}
	rowCount := len(out)
	if rowCount == 0 {
		t.Fatal("No Deployment rows found for filter", testDeploymentFilter)
	}

	fmt.Printf("Deployment filtered found %d rows for filter \"%s\"\n", rowCount, testDeploymentFilter)

}

func TestDeploymentService_Update(t *testing.T) {

	// TODO: Fix TestDeploymentService_Update!
	t.SkipNow()

	aplSvc := apl.NewClient()

	in := &apl.DeploymentUpdateInput{Name: "stack artifact UPDATED!"}
	out, _, err := aplSvc.Deployments.Update(testDeploymentId, in)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Skipped:", out.Skipped)
	fmt.Println("Deleted:", out.Deleted)
	fmt.Println("Unchanged:", out.Unchanged)
	fmt.Println("Replaced:", out.Replaced)
	fmt.Println("Errors:", out.Errors)

}

func TestDeploymentService_Get(t *testing.T) {

	// TODO: Fix TestDeploymentService_Get!
	t.SkipNow()

	aplSvc := apl.NewClient()

	out, _, err := aplSvc.Deployments.Get(testDeploymentId)

	if err != nil {
		t.Fatal(err)
	}

	if out == (apl.Deployment{}) {
		t.Fatal("No Deployment found for ID", testDeploymentId)
	}

	fmt.Println("Deployment found for ID", testDeploymentId)

}

func TestDeploymentService_Delete(t *testing.T) {

	// TODO: Fix TestDeploymentService_Delete!
	t.SkipNow()

	aplSvc := apl.NewClient()
	out, _, err := aplSvc.Deployments.Delete(testDeploymentId)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Skipped:", out.Skipped)
	fmt.Println("Deleted:", out.Deleted)
	fmt.Println("Unchanged:", out.Unchanged)
	fmt.Println("Replaced:", out.Replaced)
	fmt.Println("Errors:", out.Errors)
}
