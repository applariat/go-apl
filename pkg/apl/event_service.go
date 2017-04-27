package apl

import (
	"fmt"
	"github.com/dghubble/sling"
	"net/http"
)

// EventService is the service object for event operations
type EventService struct {
	sling    *sling.Sling
	endpoint string
}

// NewEventsService return a new eventService
func NewEventsService(sling *sling.Sling) *EventService {
	return &EventService{
		sling:    sling,
		endpoint: "events",
	}
}

// Event represents a event row
type Event struct {
	ID              string `json:"id"`
	EventType       string `json:"event_type"`
	ObjectType      string `json:"object_type"`
	ObjectName      string `json:"object_name"`
	UpdateData      string `json:"update_data,omitempty"`
	Source          string `json:"source"`
	Message         string `json:"message"`
	Active          bool   `json:"active"`
	CreatedByUserID string `json:"created_by_user_id"`
	CreatedTime     string `json:"created_time"`
	LastModified    string `json:"last_modified"`
}

// EventCreateInput is used for the create of events
type EventCreateInput struct {
	ID         string `json:"id,omitempty"`
	EventType  string `json:"event_type"`
	ObjectType string `json:"object_type"`
	ObjectName string `json:"object_name"`
	UpdateData string `json:"update_data,omitempty"`
	Source     string `json:"source"`
	Message    string `json:"message"`
}

// EventParams filter parameters used in list operations
type EventParams struct {
	ObjectType      string `url:"object_type,omitempty"`
	ObjectName      string `url:"object_name,omitempty"`
	Active          bool   `url:"active,omitempty"`
	Source          string `url:"message,omitempty"`
	CreatedByUserID string `url:"created_by_user_id,omitempty"`
}

// List gets a list of events with optional filter params
func (c *EventService) List(params *EventParams) ([]Event, *http.Response, error) {
	output := &struct {
		Data []Event `json:"data"`
	}{}
	resp, err := doList(c.sling, c.endpoint, params, output)
	return output.Data, resp, err
}

// Get get a event for the id specified
func (c *EventService) Get(id string) (Event, *http.Response, error) {
	output := &struct {
		Data Event `json:"data"`
	}{}
	path := fmt.Sprintf("%s/%s", c.endpoint, id)
	resp, err := doGet(c.sling, path, output)
	return output.Data, resp, err
}

// Create will create a event
func (c *EventService) Create(input *EventCreateInput) (CreateResult, *http.Response, error) {
	return doCreate(c.sling, c.endpoint, input)
}
