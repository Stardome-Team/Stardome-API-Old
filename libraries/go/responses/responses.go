package responses

import (
	"github.com/Blac-Panda/Stardome-API/libraries/go/errors"
)

// Response This is the model for all responses returned
type Response struct {
	APIVersion *string                `json:"apiVersion,omitempty"`
	Context    *string                `json:"context,omitempty"`
	ID         *string                `json:"id,omitempty"`
	Method     *string                `json:"method,omitempty"`
	Params     map[string]interface{} `json:"params,omitempty"`
	Data
	errors.Error
}

// Data is the model for all data returned
type Data struct {
	Data interface{} `json:"data,omitempty"`
}

// DataObject contains all object property for the data model
type DataObject struct {
	Kind             string      `json:"kind,omitempty"`
	Fields           string      `json:"fields,omitempty"`
	ETag             string      `json:"etag,omitempty"`
	CurrentItemCount int         `json:"currentItemCount,omitempty"`
	ItemsPerPage     int         `json:"itemPerPage,omitempty"`
	StartIndex       int         `json:"startIndex,omitempty"`
	TotalItems       int         `json:"totalItems,omitempty"`
	PageIndex        int         `json:"pageIndex,omitempty"`
	TotalPages       int         `json:"totalPages,omitempty"`
	PageLinkTemplate string      `json:"pageLinkTemplate,omitempty"`
	Next             interface{} `json:"next,omitempty"`
	NextLink         string      `json:"nextLink,omitempty"`
	Previous         interface{} `json:"previous,omitempty"`
	PreviousLink     string      `json:"previousLink,omitempty"`
	Self             interface{} `json:"self,omitempty"`
	SelfLink         string      `json:"selfLink,omitempty"`
	Edit             interface{} `json:"edit,omitempty"`
	EditLink         string      `json:"editLink,omitempty"`
	Items            interface{} `json:"items,omitempty"`
}
