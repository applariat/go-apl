package apl_test

import (
	"fmt"
	"github.com/applariat/go-apl/apl"
	"testing"
)

var (
	testLocDeployId     string
	testLocDeployFilter = "gke"
)

//func TestLocDeployService_Create(t *testing.T) {
//
//	testLocDeployId = "loc-artifact-test-id"
//
//	aplSvs := apl.NewClient()
//
//	in := &apl.LocDeployCreateInput{
//		ID: testLocDeployId,
//		Name: "LocDeploy Test",
//		LocDeploysType: testLocDeployFilter,
//	}
//
//	out, _, err := aplSvs.LocDeploys.Create(in)
//
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	fmt.Println("New LocDeploy ID:", out.PrimaryKey)
//
//}

func TestLocDeployService_List(t *testing.T) {
	aplSvs := apl.NewClient()

	out, _, err := aplSvs.LocDeploys.List(nil)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount > 0 {
		testLocDeployId = out[0].ID
		fmt.Printf("LocDeploy found %d rows\n", rowCount)
	} else {
		t.Fatal("No LocDeploy rows found")
	}

}

func TestLocDeployService_ListByType(t *testing.T) {
	aplSvc := apl.NewClient()

	params := &apl.LocDeployParams{
		LocDeploysType: testLocDeployFilter,
	}
	out, _, err := aplSvc.LocDeploys.List(params)

	if err != nil {
		t.Fatal(err)
	}
	rowCount := len(out)
	if rowCount == 0 {
		t.Fatal("No LocDeploy rows found for filter", testLocDeployFilter)
	}

	fmt.Printf("LocDeploy filtered found %d rows for filter \"%s\"\n", rowCount, testLocDeployFilter)

}

//func TestLocDeployService_Update(t *testing.T) {
//	aplSvc := apl.NewClient()
//
//	in := &apl.LocDeployUpdateInput{Name: "name UPDATED!"}
//	out, _, err := aplSvc.LocDeploys.Update(testLocDeployId, in)
//
//	if err != nil {
//		t.Fatal(err)
//	}
//	fmt.Println("Skipped:", out.Skipped)
//	fmt.Println("Deleted:", out.Deleted)
//	fmt.Println("Unchanged:", out.Unchanged)
//	fmt.Println("Replaced:", out.Replaced)
//	fmt.Println("Errors:", out.Errors)
//}

func TestLocDeployService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.LocDeploys.Get(testLocDeployId)

	if err != nil {
		t.Fatal(err)
	}

	if out == (apl.LocDeploy{}) {
		t.Fatal("No LocDeploy found for ID", testLocDeployId)
	}

	fmt.Println("LocDeploy found for ID", testLocDeployId)

}

/* Dangerous */

//func TestLocDeployService_Delete(t *testing.T) {
//	aplSvc := apl.NewClient()
//	out, _, err := aplSvc.LocDeploys.Delete(testLocDeployId)
//
//	if err != nil {
//		t.Fatal(err)
//	}
//	fmt.Println("Skipped:", out.Skipped)
//	fmt.Println("Deleted:", out.Deleted)
//	fmt.Println("Unchanged:", out.Unchanged)
//	fmt.Println("Replaced:", out.Replaced)
//	fmt.Println("Errors:", out.Errors)
//}
