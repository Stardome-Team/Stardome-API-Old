package services

import (
	"errors"

	"github.com/Blac-Panda/Stardome-API/models"
	"github.com/Blac-Panda/Stardome-API/repositories"
)

// PlayerService :
type PlayerService interface {
	ListPlayers()
	GetPlayer()
	CreatePlayer() (*models.Player, error)
	UpdatePlayer()
	ModifyPlayer()
	DeletePlayer()
}

type service struct {
	repository repositories.PlayerRepository
}

// NewPlayerService :
func NewPlayerService(r repositories.PlayerRepository) PlayerService {
	return &service{
		repository: r,
	}
}

func (s *service) ListPlayers() {
}

func (s *service) GetPlayer() {}

func (s *service) CreatePlayer() (*models.Player, error) {
	player, err := s.repository.CreatePlayer()

	if err != nil {
		return nil, errors.New("Not yet implemented")
	}
	return player, nil
}

func (s *service) UpdatePlayer() {}

func (s *service) ModifyPlayer() {}

func (s *service) DeletePlayer() {}
