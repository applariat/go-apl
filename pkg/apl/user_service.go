package apl

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

// UserService is the service object for user operations
type UserService struct {
	sling    *sling.Sling
	endpoint string
}

// NewUsersService return a new userService
func NewUsersService(sling *sling.Sling) *UserService {
	return &UserService{
		sling:    sling,
		endpoint: "users",
	}
}

// User represents a user row
type User struct {
	ID              string      `json:"id"`
	FirstName       string      `json:"first_name"`
	LastName        string      `json:"last_name"`
	UserType        string      `json:"user_type"`
	WorkRole        string      `json:"work_role,omitempty"`
	IsDeleted       bool        `json:"is_deleted"`
	Email           string      `json:"email"`
	EmailVerified   bool        `json:"email_verified"`
	Phone           string      `json:"phone,omitempty"`
	Role            interface{} `json:"role"`
	LastLoginDate   int         `json:"last_login_date"`
	Projects        interface{} `json:"projects,omitempty"`
	CreatedByUserID string      `json:"created_by_user_id,omitempty"`
	LastModified    string      `json:"last_modified"`
	CreatedTime     string      `json:"created_time"`
}

// UserCreateInput is used for the create of users
type UserCreateInput struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserType  string `json:"user_type"`
	WorkRole  string `json:"work_role"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone,omitempty"`
	//Role            interface{} `json:"role"`
	RoleID   string      `json:"role_id"`
	Projects interface{} `json:"projects,omitempty"`
}

// UserUpdateInput is used for the update of users
type UserUpdateInput struct {
	FirstName string      `json:"first_name,omitempty"`
	LastName  string      `json:"last_name,omitempty"`
	UserType  string      `json:"user_type,omitempty"`
	WorkRole  string      `json:"work_role,omitempty"`
	Email     string      `json:"email,omitempty"`
	Password  string      `json:"password,omitempty"`
	Phone     string      `json:"phone,omitempty"`
	Role      interface{} `json:"role,omitempty"`
	Projects  interface{} `json:"projects,omitempty"`
}

// UserParams filter parameters used in list operations
type UserParams struct {
	FirstName string `url:"first_name,omitempty"`
	LastName  string `url:"last_name,omitempty"`
	UserType  string `url:"user_type,omitempty"`
	WorkRole  string `url:"work_role,omitempty"`
	Email     string `url:"email,omitempty"`
}

// List gets a list of users with optional filter params
func (c *UserService) List(params *UserParams) ([]User, *http.Response, error) {
	output := &struct {
		Data []User `json:"data"`
	}{}
	resp, err := doList(c.sling, c.endpoint, params, output)
	return output.Data, resp, err
}

// Get get a user for the id specified
func (c *UserService) Get(id string) (User, *http.Response, error) {
	output := &struct {
		Data User `json:"data"`
	}{}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := doGet(c.sling, path, output)
	return output.Data, resp, err
}

// Create will create a user
func (c *UserService) Create(input *UserCreateInput) (CreateResult, *http.Response, error) {
	return doCreate(c.sling, c.endpoint, input)
}

// Update will update a user for the id specified
func (c *UserService) Update(id string, input *UserUpdateInput) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doUpdate(c.sling, path, input)
}

// Delete will delete the user for the id specified
func (c *UserService) Delete(id string) (ModifyResult, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	return doDelete(c.sling, path)
}
