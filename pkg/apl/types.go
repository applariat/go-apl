package apl

// CreateResult is the PK of the row just created
// Or it's a more complex struct of results.
type CreateResult struct {
	Data interface{} `json:"data"`
}

// ModifyResult is the information returned after a delete/update
type ModifyResult struct {
	Skipped   int `json:"skipped"`
	Deleted   int `json:"deleted"`
	Unchanged int `json:"unchanged"`
	Errors    int `json:"errors"`
	Replaced  int `json:"replaced"`
	Inserted  int `json:"inserted"`
} // `json:"data"`

// ModifyOutput wraps a ModifyResult with 'data'
type ModifyOutput struct {
	ModifyResult `json:"data"`
}

// CreateInput Used to wrap the data json required by the api
type CreateInput struct {
	Data interface{} `json:"data"`
}

// UpdateInput Used to wrap the data json required by the api
type UpdateInput struct {
	Data interface{} `json:"data"`
}

// ListParams Used to wrap params for function passing
type ListParams struct {
	Data interface{} `json:"data"`
}

// CreatedByUser used as a join in many results
type CreatedByUser struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
} // `json:"created_by_user"`

// MetaData ...
type MetaData struct {
	Link        string `json:"link,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Description string `json:"description,omitempty"`
	Icon        string `json:"icon,omitempty"`
} // `json:"meta_data"`

// Releases ...
type Releases struct {
	ID       string `json:"id"`
	Version  int    `json:"version"`
	MetaData `json:"meta_data"`
} // `json:"releases"`

// Artifact
type Artifact struct {
	Code    interface{} `json:"code,omitempty"`
	Config  string      `json:"config,omitempty"`
	Image   string      `json:"image,omitempty"`
	Data    string      `json:"data,omitempty"`
	Builder string      `json:"builder,omitempty"`
} // `json:"artifacts"`

// Build
type Build struct {
	Artifact  `json:"artifacts"`
	BuildVars `json:"buildvars,omitempty"`
} // `json:"build,omitempty"`

// Overrides
type Overrides struct {
	Build `json:"build,omitempty"`
} // `json:"overrides,omitempty"`

// Service
type Service struct {
	ComponentServiceID string      `json:"component_service_id,omitempty"`
	Name               string      `json:"name,omitempty"`
	Build              interface{} `json:"build,omitempty"`
	Run                interface{} `json:"run,omitempty"`
	Overrides          interface{} `json:"overrides,omitempty"`
}

// BuildVars
type BuildVars []map[string]string // `json:"buildvars,omitempty"`

// EnvVars
type EnvVars []map[string]string // `json:"envvars,omitempty"`

// Run
type Run struct {
	Instances int `json:"instances,omitempty"`
} // `json:"run,omitempty"`
