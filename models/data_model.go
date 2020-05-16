package models

// Data ;
type Data struct {
	Data *DataObject `json:"data,omitempty"`
}

// DataObject ;
type DataObject struct {
	Kind             string      `json:"kind,omitempty"`
	Fields           string      `json:"fields,omitempty"`
	ETag             string      `json:"etag,omitempty"`
	ID               string      `json:"id,omitempty"`
	Update           string      `json:"update,omitempty"`
	Deleted          bool        `json:"delete,omitempty"`
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
	*Player          `json:",omitempty"`
}
