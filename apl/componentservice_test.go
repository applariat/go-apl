package apl_test

import (
	"applariat.io/go-apl/apl"
	"fmt"
	"testing"
)

func TestComponentService_List(t *testing.T) {
	aplSvs := apl.NewClient()

	out, _, err := aplSvs.Components.List(nil)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(out)

}

func TestComponentService_ListByType(t *testing.T) {
	aplSvs := apl.NewClient()

	params := &apl.ComponentParams{
		Category: "database",
	}
	out, _, err := aplSvs.Components.List(params)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(out)

}

func TestComponentService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.Components.Get("c-golang")

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(out)
}
