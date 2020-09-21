package auth

import "github.com/Blac-Panda/Stardome-API/services/identity-service/pkg/log"

type service struct {
	repo   Repository
	logger log.Logger
}

// Service contains interfaces for authentication services
type Service interface {
}

// NewService creates new authentication service
func NewService(r Repository, l log.Logger) Service {
	return &service{repo: r, logger: l}
}
