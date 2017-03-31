package apl_test

import (
	"fmt"
	"github.com/applariat/go-apl/apl"
	"testing"
)

var (
	testStackComponentId     = "test-stack-component-id"
	testStackComponentFilter = "c-mongodb"
)

func TestStackComponentService_Create(t *testing.T) {

	// TODO: Fix TestStackComponentService_Create!
	t.SkipNow()

	aplSvs := apl.NewClient()

	in := &apl.StackComponentCreateInput{
		ID: testStackComponentId,
	}

	out, _, err := aplSvs.StackComponents.Create(in)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("PrimaryKey:", out.Data)

}

func TestStackComponentService_List(t *testing.T) {
	aplSvs := apl.NewClient()

	out, _, err := aplSvs.StackComponents.List(nil)

	fmt.Println("count:", len(out))
	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount == 0 {
		t.Fatal("No StackComponent rows found")
	}

	fmt.Printf("StackComponent found %d rows\n", rowCount)

}

func TestStackComponentService_ListByType(t *testing.T) {
	aplSvc := apl.NewClient()

	params := &apl.StackComponentParams{
		ComponentID: testStackComponentFilter,
	}
	out, _, err := aplSvc.StackComponents.List(params)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount == 0 {
		t.Fatal("No StackComponent rows found for filter", testStackComponentFilter)
	}

	fmt.Printf("StackComponent filtered found %d rows for filter \"%s\"\n", rowCount, testStackComponentFilter)

}

func TestStackComponentService_Update(t *testing.T) {

	// TODO: Fix TestStackComponentService_Update!
	t.SkipNow()

	aplSvc := apl.NewClient()

	in := &apl.StackComponentUpdateInput{Name: "stack component UPDATED!"}
	out, _, err := aplSvc.StackComponents.Update(testStackComponentId, in)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Skipped:", out.Skipped)
	fmt.Println("Deleted:", out.Deleted)
	fmt.Println("Unchanged:", out.Unchanged)
	fmt.Println("Replaced:", out.Replaced)
	fmt.Println("Errors:", out.Errors)

}

func TestStackComponentService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.StackComponents.Get(testStackComponentId)

	if err != nil {
		t.Fatal(err)
	}
	if out == (apl.StackComponent{}) {
		t.Fatal("No StackComponent found for ID", testStackComponentId)
	}

	fmt.Println("StackComponent found for ID", testStackComponentId)
}

func TestStackComponentService_Delete(t *testing.T) {

	// TODO: Fix TestStackComponentService_Delete!
	t.SkipNow()

	aplSvc := apl.NewClient()
	out, _, err := aplSvc.StackComponents.Delete(testStackComponentId)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Skipped:", out.Skipped)
	fmt.Println("Deleted:", out.Deleted)
	fmt.Println("Unchanged:", out.Unchanged)
	fmt.Println("Replaced:", out.Replaced)
	fmt.Println("Errors:", out.Errors)
}
