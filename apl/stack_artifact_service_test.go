package apl_test

import (
	"github.com/applariat/go-apl/apl"
	"fmt"
	"testing"
)


func TestStackArtifactService_Create(t *testing.T) {
	aplSvs := apl.NewClient()

	in := &apl.StackArtifactCreateInput{
		ID:             "chris-test-id",
		Name: "Chris Test Zip",
		ArtifactName:           "Chris/chris.zip",
		LocArtifactID: "la-gs-apl",
		ProjectID: "p-mobile-apps-apl",
		StackID: "f5cdb3c7-992f-4109-8c97-21b43edf5254",
		Version: "1.1.1",
		Package: "archive",
	}

	out, _, err := aplSvs.StackArtifacts.Create(in)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("PrimaryKey:", out.PrimaryKey)

}

func TestStackArtifactService_List(t *testing.T) {
	aplSvs := apl.NewClient()

	out, _, err := aplSvs.StackArtifacts.List(nil)

	fmt.Println("count:", len(out))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(out)

}

func TestStackArtifactService_ListByType(t *testing.T) {
	aplSvc := apl.NewClient()

	params := &apl.StackArtifactParams{
		StackID: "86235abc-e2a9-4086-81df-e4965faf6aa2",
	}
	out, _, err := aplSvc.StackArtifacts.List(params)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(out)

}

func TestStackArtifactService_Update(t *testing.T) {
	aplSvc := apl.NewClient()

	in := &apl.StackArtifactUpdateInput{Name: "stack artifact UPDATED!"}
	out, _, err := aplSvc.StackArtifacts.Update("chris-test-id", in)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Updated:", out)

}

func TestStackArtifactService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.StackArtifacts.Get("chris-test-id")

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Name:", out.Name)
}

func TestStackArtifactService_Delete(t *testing.T) {
	aplSvc := apl.NewClient()
	out, _, err := aplSvc.StackArtifacts.Delete("chris-test-id")

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Skipped:", out.Skipped)
	fmt.Println("Deleted:", out.Deleted)
	fmt.Println("Unchanged:", out.Unchanged)
	fmt.Println("Replaced:", out.Replaced)
	fmt.Println("Errors:", out.Errors)
}
