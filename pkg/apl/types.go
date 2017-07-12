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
	Config  string `json:"config,omitempty"`
	Image   string `json:"image,omitempty"`
	Data    string `json:"data,omitempty"`
	Builder string `json:"builder,omitempty"`
} // `json:"artifacts"`

// Build
type Build struct {
	Artifact  `json:"artifacts"`
	BuildVars `json:"buildvars,omitempty"`
} // `json:"build,omitempty"`

// Service
type Service struct {
	ComponentServiceID string `json:"component_service_id"`
	Build              `json:"build,omitempty"`
	Run                `json:"run,omitempty"`
}

// BuildVars
type BuildVars []map[string]string // `json:"buildvars,omitempty"`

// EnvVars
type EnvVars []map[string]string // `json:"envvars,omitempty"`

// Run
type Run struct {
	Instances int `json:"instances,omitempty"`
} // `json:"run,omitempty"`

//// HealthProbe ...
//type HealthProbe   struct {
//	Delay   interface{} `json:"delay,omitempty"`
//	URL     interface{} `json:"url,omitempty"`
//	Type    interface{} `json:"type,omitempty"`
//	Port    interface{} `json:"port,omitempty"`
//	Timeout interface{} `json:"timeout,omitempty"`
//} // `json:"health_probe"`
//
//// Defaults ...
//type Defaults struct {
//	// TODO: Figure out why this is an int and not string
//	//ReserveMemory   string `json:"reserve_memory,omitempty"`
//	Instances       int    `json:"instances,omitempty"`
//	StorageSize     int    `json:"storage_size,omitempty"`
//	ClientPort      int    `json:"client_port,omitempty"`
//	LimitMemory     string `json:"limit_memory,omitempty"`
//	StorageType     string `json:"storage_type,omitempty"`
//	HealthCheckType string `json:"health_check_type,omitempty"`
//	ServicePort     int    `json:"service_port,omitempty"`
//	ServiceName     string `json:"service_name,omitempty"`
//	ReserveCPU      string `json:"reserve_cpu,omitempty"`
//	ExternalNetwork bool   `json:"external_network,omitempty"`
//	LimitCPU        string `json:"limit_cpu,omitempty"`
//	HealthProbe     `json:"health_probe,omitempty"`
//} // `json:"defaults"`
