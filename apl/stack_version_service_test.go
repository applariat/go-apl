package apl_test

import (
	"fmt"
	"github.com/applariat/go-apl/apl"
	"testing"
)

var (
	testStackVersionId = "test-stack-version-id"
)

func TestStackVersionService_Create(t *testing.T) {

	// TODO: Fix TestStackVersionService_Create!
	t.SkipNow()

	aplSvs := apl.NewClient()

	in := &apl.StackVersionCreateInput{
		StackID: testStackVersionId,
	}

	out, _, err := aplSvs.StackVersions.Create(in)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("PrimaryKey:", out.Data)

}

func TestStackVersionService_List(t *testing.T) {
	aplSvs := apl.NewClient()

	out, _, err := aplSvs.StackVersions.List()

	fmt.Println("count:", len(out))
	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount == 0 {
		t.Fatal("No StackVersion rows found")
	}

	fmt.Printf("StackVersion found %d rows\n", rowCount)

}

func TestStackVersionService_Get(t *testing.T) {

	// TODO: Fix TestStackVersionService_Get!
	t.SkipNow()

	aplSvc := apl.NewClient()

	out, _, err := aplSvc.StackVersions.Get(testStackVersionId)

	if err != nil {
		t.Fatal(err)
	}
	if out == (apl.StackVersion{}) {
		t.Fatal("No StackVersion found for ID", testStackVersionId)
	}

	fmt.Println("StackVersion found for ID", testStackVersionId)
}

func TestStackVersionService_Delete(t *testing.T) {

	// TODO: Fix TestStackVersionService_Delete!
	t.SkipNow()

	aplSvc := apl.NewClient()
	out, _, err := aplSvc.StackVersions.Delete(testStackVersionId)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Skipped:", out.Skipped)
	fmt.Println("Deleted:", out.Deleted)
	fmt.Println("Unchanged:", out.Unchanged)
	fmt.Println("Replaced:", out.Replaced)
	fmt.Println("Errors:", out.Errors)
}
