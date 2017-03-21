package apl_test

import (
	"applariat.io/go-apl/apl"
	"fmt"
	"testing"
)

func TestStackVersionService_Create(t *testing.T) {
	aplSvs := apl.NewClient()

	in := &apl.StackVersionInput{
		ID:             "chris-test-id",
	}

	out, _, err := aplSvs.StackVersions.Create(in)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("PrimaryKey:", out.PrimaryKey)

}

func TestStackVersionService_List(t *testing.T) {
	aplSvs := apl.NewClient()

	out, _, err := aplSvs.StackVersions.List(nil)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(out)

}

func TestStackVersionService_ListByType(t *testing.T) {
	aplSvc := apl.NewClient()

	params := &apl.StackVersionParams{
		StackID: "25",
	}
	out, _, err := aplSvc.StackVersions.List(params)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(out)

}

func TestStackVersionService_Update(t *testing.T) {
	aplSvc := apl.NewClient()

	in := &apl.StackVersionInput{Name: "creds for chris UPDATED!"}
	out, _, err := aplSvc.StackVersions.Update("chris-test-id", in)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Updated:", out)

}

func TestStackVersionService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.StackVersions.Get("chris-test-id")

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(out)
}

func TestStackVersionService_Delete(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.StackVersions.Delete("chris-test-id")

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Deleted:", out.Result.Deleted)

}
