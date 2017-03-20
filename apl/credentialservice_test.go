package apl_test

import (
	"applariat.io/go-apl/apl"
	"fmt"
	"testing"
)

func TestCredentialService_Create(t *testing.T) {
	aplSvs := apl.NewClient()

	in := &apl.CredentialInput{
		ID:             "chris-test-id",
		Name:           "creds for chris",
		CredentialType: apl.CredentialTypeDocker,
		Credentials: map[string]string{
			"registry_email":    "floopy@glorp.com",
			"registry_password": "urpasswd",
			"registry_user":     "apl_registry",
		},
	}

	out, _, err := aplSvs.Credentials.Create(in)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("PrimaryKey:", out.PrimaryKey)

}

func TestCredentialService_List(t *testing.T) {
	aplSvs := apl.NewClient()

	out, _, err := aplSvs.Credentials.List(nil)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(out)

}

func TestCredentialService_ListByType(t *testing.T) {
	aplSvc := apl.NewClient()

	params := &apl.CredentialParams{
		CredentialType: apl.CredentialTypeDocker,
	}
	out, _, err := aplSvc.Credentials.List(params)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(out)

}

func TestCredentialService_Update(t *testing.T) {
	aplSvc := apl.NewClient()

	in := &apl.CredentialInput{Name: "creds for chris UPDATED!"}
	out, _, err := aplSvc.Credentials.Update("chris-test-id", in)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Updated:", out)

}

func TestCredentialService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.Credentials.Get("chris-test-id")

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Updated Name:", out.Name)
}

func TestCredentialService_Delete(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.Credentials.Delete("chris-test-id")

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Deleted:", out.Result.Deleted)

}
