package apl_test

import (
	"github.com/applariat/go-apl/apl"
	"fmt"
	"testing"
)

var (
	testOrgId string
)

func TestOrgService_List(t *testing.T) {
	aplSvs := apl.NewClient()

	out, _, err := aplSvs.Orgs.List()

	if err != nil {
		t.Fatal(err)
	}

	if out == (apl.Org{}) {
		t.Fatal("No Org found")
	}

	testOrgId = out.ID

	fmt.Println("Org found for ID", testOrgId)


}

func TestOrgService_Update(t *testing.T) {
	aplSvc := apl.NewClient()

	in := &apl.OrgUpdateInput{NumOfEmployees: "99"}
	out, _, err := aplSvc.Orgs.Update(testOrgId, in)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Skipped:", out.Skipped)
	fmt.Println("Deleted:", out.Deleted)
	fmt.Println("Unchanged:", out.Unchanged)
	fmt.Println("Replaced:", out.Replaced)
	fmt.Println("Errors:", out.Errors)
}

func TestOrgService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.Orgs.Get(testOrgId)

	if err != nil {
		t.Fatal(err)
	}

	if out == (apl.Org{}) {
		t.Fatal("No Org found for ID", testOrgId)
	}

	fmt.Println("Org found for ID", testOrgId)

}

