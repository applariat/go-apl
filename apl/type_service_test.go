package apl_test

import (
	"github.com/applariat/go-apl/apl"
	"fmt"
	"testing"
)

func TestTypeService_List(t *testing.T) {
	aplSvs := apl.NewClient()

	out, _, err := aplSvs.Types.List()

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount == 0 {
		t.Fatal("No Type rows found")
	}

	fmt.Printf("Type filtered found %d rows", rowCount)


}

func TestTypeService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.Types.Get("credential_type")

	if err != nil {
		t.Fatal(err)
	}

	if out == (apl.Type{}) {
		t.Fatal("No Type found for ID", "credential_type")
	}

	fmt.Println("Type found for ID", "credential_type")

}
