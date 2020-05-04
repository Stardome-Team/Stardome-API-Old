package services

import (
	"time"

	"github.com/Blac-Panda/Stardome-API/models"
	"github.com/Blac-Panda/Stardome-API/repositories"
	"github.com/Blac-Panda/Stardome-API/utils"
	"github.com/rs/xid"
)

// PlayerService :
type PlayerService interface {
	ListPlayers()
	GetPlayer()
	CreatePlayer(pr *models.PlayerRegistration) (*models.Player, error)
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

func (s *service) CreatePlayer(pr *models.PlayerRegistration) (*models.Player, error) {

	guid := xid.New()
	passwordHash, err := utils.GenerateHashFromPassword(pr.Password)

	if err != nil {
		// TODO: Log Error
		return nil, err
	}

	playerID := guid.String()
	time := time.Now()

	var newPlayer models.Player = models.Player{
		ID:        &playerID,
		UserName:  &pr.UserName,
		PassHash:  &passwordHash,
		CreatedAt: &time,
		UpdatedAt: &time,
	}

	player, err := s.repository.CreatePlayer(&newPlayer)

	if err != nil {
		return nil, err
	}
	return player, nil
}

func (s *service) UpdatePlayer() {}

func (s *service) ModifyPlayer() {}

func (s *service) DeletePlayer() {}
