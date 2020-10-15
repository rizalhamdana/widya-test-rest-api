package share

// Parameters data structure
type Parameters struct {
	Query   string
	Limit   int
	Offset  int
	Sort    string
	OrderBy string
	All     bool
	Search  string
	Filter  interface{}
}

// ResultRepository struct
type ResultRepository struct {
	Result interface{}
	Count  int
	Error  error
}

// List Result
type ListResult struct {
	Meta *Meta         `json:"meta"`
	Data []interface{} `json:"data"`
}

// Meta data structure
type Meta struct {
	Offset int64 `json:"offest"`
	Limit  int64 `json:"limit"`
	Count  int64 `json:"count"`
}
