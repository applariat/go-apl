package apl

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

// AuditService is the service object for audit operations
type AuditService struct {
	sling    *sling.Sling
	endpoint string
}

// NewAuditsService return a new AuditService
func NewAuditsService(sling *sling.Sling) *AuditService {
	return &AuditService{
		sling:    sling,
		endpoint: "audits",
	}
}

// Audit represents a audit row
type Audit struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	ResourceID   string `json:"resource_id"`
	Timestamp    int64  `json:"timestamp"`
	Action       string `json:"action"`
	Message      string `json:"message"`
	ResourceType string `json:"resource_type"`
	Params       interface{} `json:"params,omitempty"`
	Result       interface{} `json:"result"`
	Data         interface{} `json:"data,omitempty"`
}

// AuditParams filter parameters used in list operations
type AuditParams struct {
	UserID       string `url:"user_id,omitempty"`
	ResourceID   string `url:"resource_id,omitempty"`
	Action       string `url:"action,omitempty"`
	ResourceType string `url:"resource_type,omitempty"`
}

// List gets a list of audits with optional filter params
func (c *AuditService) List(params *AuditParams) ([]Audit, *http.Response, error) {
	output := &struct{ Data []Audit `json:"data"` }{}
	resp, err := doList(c.sling, c.endpoint, params, output)
	return output.Data, resp, err
}

// Get get a audit for the id specified
func (c *AuditService) Get(id string) (Audit, *http.Response, error) {
	output := &struct{ Data Audit `json:"data"` }{}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := doGet(c.sling, path, output)
	return output.Data, resp, err
}