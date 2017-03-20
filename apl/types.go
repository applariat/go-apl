package apl

// CreateOutput is the PK of the row just created
type CreateOutput struct {
	PrimaryKey string `json:"data"`
}

// ModifyOutput is the information returned after a delete/update
type ModifyOutput struct {
	Result struct {
		Skipped   int `json:"skipped"`
		Deleted   int `json:"deleted"`
		Unchanged int `json:"unchanged"`
		Errors    int `json:"errors"`
		Replaced  int `json:"replaced"`
		Inserted  int `json:"inserted"`
	} `json:"data"`
}

// CreateInput Used to wrap the data json required by the api
type CreateInput struct {
	Data interface{} `json:"data"`
}

// CreatedByUser used as a join in many results
type CreatedByUser struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
