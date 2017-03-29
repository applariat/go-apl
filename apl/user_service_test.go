package apl_test

import (
	"github.com/applariat/go-apl/apl"
	"fmt"
	"testing"
	"encoding/base64"
)

func TestUserService_Create(t *testing.T) {
	aplSvs := apl.NewClient()

	data := []byte("transfloopiglorpsnarfle")
	passwd := base64.StdEncoding.EncodeToString(data)

	in := &apl.UserCreateInput{
		ID: "chris-test-id",
		FirstName: "TestFirstName",
		LastName: "TestLastName",
		Email: "testing@example.com",
		Password: passwd,
		UserType: "user",
		WorkRole: "Unknown",
		RoleId: "27a6f6fa-0209-495b-95f0-1266dd14b399",
	}

	out, _, err := aplSvs.Users.Create(in)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("PrimaryKey:", out.PrimaryKey)

}

func TestUserService_List(t *testing.T) {
	aplSvs := apl.NewClient()

	out, _, err := aplSvs.Users.List(nil)

	fmt.Println("count:", len(out))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(out)

}

func TestUserService_ListByType(t *testing.T) {
	aplSvc := apl.NewClient()

	params := &apl.UserParams{
		UserType: "user",
	}
	out, _, err := aplSvc.Users.List(params)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(out)

}

func TestUserService_Update(t *testing.T) {
	aplSvc := apl.NewClient()

	in := &apl.UserUpdateInput{FirstName: "UPDATED!"}
	out, _, err := aplSvc.Users.Update("chris-test-id", in)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Updated:", out)

}

func TestUserService_Get(t *testing.T) {
	aplSvc := apl.NewClient()

	out, _, err := aplSvc.Users.Get("chris-test-id")

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("FirstName:", out.FirstName)
}

func TestUserService_Delete(t *testing.T) {
	aplSvc := apl.NewClient()
	out, _, err := aplSvc.Users.Delete("chris-test-id")

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Skipped:", out.Skipped)
	fmt.Println("Deleted:", out.Deleted)
	fmt.Println("Unchanged:", out.Unchanged)
	fmt.Println("Replaced:", out.Replaced)
	fmt.Println("Errors:", out.Errors)
}
