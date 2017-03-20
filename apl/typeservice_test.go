package apl_test

import (
	"applariat.io/go-apl/apl"
	"fmt"
	"testing"
)

func TestTypeService_List(t *testing.T) {
	aplSvs := apl.NewClient()

	out, _, err := aplSvs.Types.List()

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(out)

}

func TestTypeService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.Types.Get("loc_deploys_type")

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(out)
}
