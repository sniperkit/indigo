package bulk

type Document struct {
	Id     string                 `json:"id,omitempty"`
	Fields map[string]interface{} `json:"fields,omitempty"`
}

type Request struct {
	Method   string   `json:"method,omitempty"`
	Document Document `json:"document,omitempty"`
}

type Resource struct {
	BatchSize int32     `json:"batch_size,omitempty"`
	Requests  []Request `json:"requests,omitempty"`
}
