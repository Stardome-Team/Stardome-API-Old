package models

// Pagination :
type Pagination struct {
	StartIndex       *int
	ItemsPerPage     *int
	TotalItems       *int
	CurrentItemCount int
	Items            interface{}
}

// Token ;
type Token struct {
	Token     *string `json:"token"`
	ExpiresOn int64   `json:"expiresOn"`
	Type      string  `json:"type"`
}
