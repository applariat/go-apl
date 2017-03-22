package apl

// CreateOutput is the PK of the row just created
type CreateResult struct {
	PrimaryKey string `json:"data"`
}

type ModifyResult struct {
	Skipped   int `json:"skipped"`
	Deleted   int `json:"deleted"`
	Unchanged int `json:"unchanged"`
	Errors    int `json:"errors"`
	Replaced  int `json:"replaced"`
	Inserted  int `json:"inserted"`
} // `json:"data"`


// ModifyOutput is the information returned after a delete/update
type ModifyOutput struct {
	ModifyResult `json:"data"`
}

// CreateInput Used to wrap the data json required by the api
type CreateInput struct {
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

type Releases struct {
	ID string `json:"id"`
	Version int `json:"version"`
	MetaData `json:"meta_data"`
} // `json:"releases"`

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
