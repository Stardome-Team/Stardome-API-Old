package models

// Pagination :
type Pagination struct {
	StartIndex       *int
	ItemsPerPage     *int
	TotalItems       *int
	CurrentItemCount int
	Items            interface{}
}
