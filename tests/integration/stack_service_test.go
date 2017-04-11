// +build integration

package tests

import (
	"fmt"
	"github.com/applariat/go-apl/pkg/apl"
	"testing"
)

var (
	testStackId     = "test-stack-id"
	testStackFilter = "acme-test"
)

func TestStackService_Create(t *testing.T) {

	// TODO: Fix TestStackService_Create!
	t.SkipNow()

	in := &apl.StackCreateInput{
		ID:   testStackId,
		Name: testStackFilter,
		//StackVersions: "[]",
		//Project: map[string]string{},
		//StackArtifacts: "[]",
	}

	out, _, err := aplClient.Stacks.Create(in)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("PrimaryKey:", out.Data)

}

func TestStackService_List(t *testing.T) {

	out, _, err := aplClient.Stacks.List(nil)

	fmt.Println("count:", len(out))
	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount == 0 {
		t.Fatal("No Stack rows found")
	}

	fmt.Printf("Stack found %d rows\n", rowCount)

}

func TestStackService_ListByType(t *testing.T) {
	aplSvc := apl.NewClient()

	params := &apl.StackParams{
		Name: testStackFilter,
	}
	out, _, err := aplSvc.Stacks.List(params)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount == 0 {
		t.Fatal("No Stack rows found for filter", testStackFilter)
	}

	fmt.Printf("Stack filtered found %d rows for filter \"%s\"\n", rowCount, testStackFilter)

}

func TestStackService_Update(t *testing.T) {

	// TODO: Fix TestStackService_Update!
	t.SkipNow()

	aplSvc := apl.NewClient()

	in := &apl.StackUpdateInput{Name: testStackFilter}
	out, _, err := aplSvc.Stacks.Update(testStackId, in)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Skipped:", out.Skipped)
	fmt.Println("Deleted:", out.Deleted)
	fmt.Println("Unchanged:", out.Unchanged)
	fmt.Println("Replaced:", out.Replaced)
	fmt.Println("Errors:", out.Errors)

}

func TestStackService_Get(t *testing.T) {

	// TODO: Fix TestStackService_Get!
	t.SkipNow()

	aplSvc := apl.NewClient()

	out, _, err := aplSvc.Stacks.Get(testStackId)

	if err != nil {
		t.Fatal(err)
	}
	if out == (apl.Stack{}) {
		t.Fatal("No Stack found for ID", testStackId)
	}

	fmt.Println("Stack found for ID", testStackId)
}

func TestStackService_Delete(t *testing.T) {

	// TODO: Fix TestStackService_Delete!
	t.SkipNow()

	aplSvc := apl.NewClient()
	out, _, err := aplSvc.Stacks.Delete(testStackId)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Skipped:", out.Skipped)
	fmt.Println("Deleted:", out.Deleted)
	fmt.Println("Unchanged:", out.Unchanged)
	fmt.Println("Replaced:", out.Replaced)
	fmt.Println("Errors:", out.Errors)
}
