package apl

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

// FormService is the service object for form operations
type FormService struct {
	sling    *sling.Sling
	endpoint string
}

// NewFormsService return a new FormService
func NewFormsService(sling *sling.Sling) *FormService {
	return &FormService{
		sling:    sling,
		endpoint: "forms",
	}
}

// Form represents a form row
type Form struct {
	ID            string      `json:"id,omitempty"`
	Name          string      `json:"name"`
	AllowMultiple bool        `json:"allowMultiple,omitempty"`
	Form          interface{} `json:"form"`
	Model         interface{} `json:"model"`
	Schema        interface{} `json:"schema"`
	LastModified  string      `json:"last_modified"`
	CreatedTime   string      `json:"created_time"`
}

// FormParams filter parameters used in list operations
type FormParams struct {
	Name string `url:"name,omitempty"`
}

// List gets a list of forms with optional filter params
func (c *FormService) List(params *FormParams) ([]Form, *http.Response, error) {
	output := &struct {
		Data []Form `json:"data"`
	}{}
	resp, err := doList(c.sling, c.endpoint, params, output)
	return output.Data, resp, err
}

// Get get a form for the id specified
func (c *FormService) Get(id string) (Form, *http.Response, error) {
	output := &struct {
		Data Form `json:"data"`
	}{}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := doGet(c.sling, path, output)
	return output.Data, resp, err
}
