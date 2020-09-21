package auth

import (
	"github.com/Blac-Panda/Stardome-API/services/identity-service/pkg/database"
	"github.com/Blac-Panda/Stardome-API/services/identity-service/pkg/log"
)

type repository struct {
	db     *database.DB
	logger log.Logger
}

// Repository contains interfaces for authentication services
type Repository interface {
	FindPlayerBy(username string) error
}

// NewRepository creates a new instance for authentication repository
func NewRepository(db *database.DB, l log.Logger) Repository {
	return &repository{db: db, logger: l}
}

func (r *repository) FindPlayerBy(username string) error {

	return nil
}
