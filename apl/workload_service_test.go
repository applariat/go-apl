package apl_test

import (
	"github.com/applariat/go-apl/apl"
	"fmt"
	"testing"
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
//
func TestWorkloadService_List(t *testing.T) {
	aplSvs := apl.NewClient()

	out, _, err := aplSvs.Workloads.List(nil)

	fmt.Println("count:", len(out))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(out)

}

func TestWorkloadService_ListByType(t *testing.T) {
	aplSvc := apl.NewClient()

	params := &apl.WorkloadParams{
		WorkloadType: "wl-level1",
	}
	out, _, err := aplSvc.Workloads.List(params)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(out)

}

func TestWorkloadService_Update(t *testing.T) {
	aplSvc := apl.NewClient()

	in := &apl.WorkloadUpdateInput{Name: "workload for chris UPDATED!"}
	out, _, err := aplSvc.Workloads.Update("chris-test-id", in)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Updated:", out)

}

func TestWorkloadService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.Workloads.Get("chris-test-id")

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Name:", out.Name)
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
