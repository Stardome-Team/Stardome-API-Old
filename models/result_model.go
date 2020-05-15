package models

// Result :
type Result struct {
	APIVersion *string      `json:"apiVersion,omitempty"`
	Context    *string      `json:"context,omitempty"`
	ID         *string      `json:"id,omitempty"`
	Method     *string      `json:"method,omitempty"`
	Params     *interface{} `json:"params,omitempty"`
	Data
	Error
}
