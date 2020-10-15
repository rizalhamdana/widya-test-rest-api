package helper

// Parameters data structure
type Parameters struct {
	CategoryIDs []int
	Query       string
	Limit       int64
	Offset      int64
	Status      string
	Sort        string
	OrderBy     string
	Channel     string
	Filter      interface{}
}

// ResponseDetail data structure
type ResponseDetail struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ResponseDetail data structure
type ResponseDelete struct {
	Message string `json:"message"`
}

// ResponseList data structure
type ResponseList struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    Meta        `json:"meta,omitempty"`
}

// Meta data structure
type Meta struct {
	Count  int         `json:"count"`
	Limit  int         `json:"limit,omitempty"`
	Offset int         `json:"offset,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

// ResponseDetailOutput function for contruct output response detail
func ResponseDetailOutput(message string, data interface{}) ResponseDetail {
	res := ResponseDetail{
		Message: message,
		Data:    data,
	}
	return res
}

// ResponseDetailOutput function for contruct output response detail
func ResponseDeleteOutput(message string) ResponseDelete {
	res := ResponseDelete{
		Message: message,
	}
	return res
}
