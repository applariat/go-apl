package apl_test

import (
	"github.com/applariat/go-apl/apl"
	"fmt"
	"testing"
)

var (
	testWorkloadId = "wl-level2-apl"
	testWorkloadFilter = "wl-level1"
)

//func TestWorkloadService_Create(t *testing.T) {
//	aplSvs := apl.NewClient()
//
//	in := &apl.WorkloadCreateInput{
//		ID:             "chris-test-id",
//		Name:           "creds for chris",
//		WorkloadType: 	"wl-level1",
//		Description:    "The Test Description",
//		LeaseType: 		"temporary",
//		MaxLeasePeriodDays: 7,
//		Priority: 			300,
//		QualityOfService: "best_effort",
//	}
//
//	out, _, err := aplSvs.Workloads.Create(in)
//
//	if err != nil {
//		t.Fatal(err)
//	}
//	fmt.Println("PrimaryKey:", out.PrimaryKey)
//
//}

func TestWorkloadService_List(t *testing.T) {
	aplSvs := apl.NewClient()

	out, _, err := aplSvs.Workloads.List(nil)

	fmt.Println("count:", len(out))
	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount > 0 {
		fmt.Printf("Workload found %d rows\n", rowCount)
	} else {
		t.Fatal("No Workload rows found")
	}


}

func TestWorkloadService_ListByType(t *testing.T) {
	aplSvc := apl.NewClient()

	params := &apl.WorkloadParams{
		WorkloadType: testWorkloadFilter,
	}
	out, _, err := aplSvc.Workloads.List(params)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount == 0 {
		t.Fatal("No Workload rows found for filter", testWorkloadFilter)
	}

	fmt.Printf("Workload filtered found %d rows for filter \"%s\"\n", rowCount, testWorkloadFilter)

}

//func TestWorkloadService_Update(t *testing.T) {
//	aplSvc := apl.NewClient()
//
//	in := &apl.WorkloadUpdateInput{Name: "UPDATED!"}
//	out, _, err := aplSvc.Workloads.Update(testWorkloadId, in)
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

func TestWorkloadService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.Workloads.Get(testWorkloadId)

	if err != nil {
		t.Fatal(err)
	}

	if out == (apl.Workload{}) {
		t.Fatal("No Workload found for ID", testWorkloadId)
	}

	fmt.Println("Workload found for ID", testWorkloadId)

}

//func TestWorkloadService_Delete(t *testing.T) {
//	aplSvc := apl.NewClient()
//	out, _, err := aplSvc.Workloads.Delete("chris-test-id")
//
//	if err != nil {
//		t.Fatal(err)
//	}
//  fmt.Println("Skipped:", out.Skipped)
//  fmt.Println("Deleted:", out.Deleted)
//  fmt.Println("Unchanged:", out.Unchanged)
//  fmt.Println("Replaced:", out.Replaced)
//  fmt.Println("Errors:", out.Errors)
//}
