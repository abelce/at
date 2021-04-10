package jsonapi

type JsonapiDocument struct {
	Errors []JsonapiError `json:"errors,omitempty"`
}
type JsonapiError struct {
	Code   int                    `json:"code,omitempty"`
	Detail string                 `json:"detail"`
	Meta   map[string]interface{} `json:"meta"`
}
