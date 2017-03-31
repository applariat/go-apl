package apl_test

import (
	"fmt"
	"github.com/applariat/go-apl/apl"
	"testing"
)

var (
	testAuditId     string
	testAuditFilter string
)

func TestAuditService_List(t *testing.T) {
	aplSvs := apl.NewClient()

	out, _, err := aplSvs.Audits.List(nil)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount > 0 {
		testAuditId = out[0].ID
		testAuditFilter = out[0].ResourceType
	} else {
		t.Fatal("No audit rows found")
	}

	fmt.Printf("Audit found %d rows\n", rowCount)

}

func TestAuditService_ListByType(t *testing.T) {
	aplSvs := apl.NewClient()

	params := &apl.AuditParams{
		ResourceType: testAuditFilter,
	}
	out, _, err := aplSvs.Audits.List(params)

	if err != nil {
		t.Fatal(err)
	}

	rowCount := len(out)
	if rowCount == 0 {
		t.Fatal("No Audit rows found for filter", testAuditFilter)
	}

	fmt.Printf("Audit filtered found %d rows for filter \"%s\"\n", rowCount, testAuditFilter)

}

func TestAuditService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.Audits.Get(testAuditId)

	if err != nil {
		t.Fatal(err)
	}

	if out == (apl.Audit{}) {
		t.Fatal("No Audit found for ID", testAuditId)
	}

	fmt.Println("Audit found for ID", testAuditId)
}
