package apl_test

import (
	"fmt"
	"github.com/applariat/go-apl/apl"
	"testing"
)

var (
	testFormId     string
	testFormFilter string
)

func TestFormService_List(t *testing.T) {
	aplSvs := apl.NewClient()

	out, _, err := aplSvs.Forms.List(nil)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount > 0 {
		testFormId = out[0].ID
		testFormFilter = out[0].Name
	} else {
		t.Fatal("No Form rows found")
	}

	fmt.Printf("Form found %d rows\n", rowCount)

}

func TestFormService_ListByType(t *testing.T) {
	aplSvc := apl.NewClient()

	params := &apl.FormParams{
		Name: testFormFilter,
	}
	out, _, err := aplSvc.Forms.List(params)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount == 0 {
		t.Fatal("No Form rows found for filter", testFormFilter)
	}

	fmt.Printf("Form filtered found %d rows for filter \"%s\"\n", rowCount, testFormFilter)

}

func TestFormService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.Forms.Get(testFormId)

	if err != nil {
		t.Fatal(err)
	}

	if out == (apl.Form{}) {
		t.Fatal("No Form found for ID", testFormId)
	}

	fmt.Println("Form found for ID", testFormId)
}
