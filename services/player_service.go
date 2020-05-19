package services

import (
	"net/http"
	"time"

	"github.com/Blac-Panda/Stardome-API/models"
	"github.com/Blac-Panda/Stardome-API/repositories"
	"github.com/Blac-Panda/Stardome-API/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

// PlayerService :
type PlayerService interface {
	ListPlayers(index, size int) (*models.Pagination, *models.ErrorParsing)
	GetPlayer(id string) (*models.Player, *models.ErrorParsing)
	CreatePlayer(pr *models.PlayerRegistration) (*models.Player, *models.ErrorParsing)
	UpdatePlayer(id string, player *models.Player) (*models.Player, *models.ErrorParsing)
	ModifyPlayer(id string, player map[string]interface{}) (*models.Player, *models.ErrorParsing)
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

func (s *service) ListPlayers(index, size int) (*models.Pagination, *models.ErrorParsing) {
	players, err := s.repository.ListPlayers(index, size)

	if err != nil {
		return nil, &models.ErrorParsing{
			Error:      err,
			Type:       gin.ErrorTypePublic,
			Metadata:   http.StatusText(http.StatusInternalServerError),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return players, nil
}

func (s *service) GetPlayer(id string) (*models.Player, *models.ErrorParsing) {
	player, err := s.repository.GetPlayer(id)

	if err != nil {
		return nil, &models.ErrorParsing{
			Error:      err,
			Type:       gin.ErrorTypePublic,
			Metadata:   http.StatusText(http.StatusNotFound),
			StatusCode: http.StatusNotFound,
		}
	}

	return player, nil
}

func (s *service) CreatePlayer(pr *models.PlayerRegistration) (*models.Player, *models.ErrorParsing) {

	guid := xid.New()
	passwordHash, err := utils.GenerateHashFromPassword(pr.Password)

	if err != nil {
		// TODO: Log Error
		return nil, &models.ErrorParsing{
			Error:      utils.ErrorEncryptionFailed,
			Type:       gin.ErrorTypePublic,
			Metadata:   http.StatusText(http.StatusPreconditionFailed),
			StatusCode: http.StatusPreconditionFailed,
		}
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
		return nil, &models.ErrorParsing{
			Error:      err,
			Type:       gin.ErrorTypePublic,
			Metadata:   http.StatusText(http.StatusConflict),
			StatusCode: http.StatusConflict,
		}
	}
	return player, nil
}

func (s *service) UpdatePlayer(id string, player *models.Player) (*models.Player, *models.ErrorParsing) {

	if id != *player.ID {
		return nil, &models.ErrorParsing{
			Error:      utils.ErrorRequestIDMismatch,
			Type:       gin.ErrorTypePublic,
			Metadata:   utils.ReasonIDMismatch,
			StatusCode: http.StatusBadRequest,
		}
	}

	player, err := s.repository.UpdatePlayer(player)

	if err != nil {
		return nil, &models.ErrorParsing{
			Error:      err,
			Type:       gin.ErrorTypePublic,
			Metadata:   http.StatusText(http.StatusNotFound),
			StatusCode: http.StatusNotFound,
		}
	}

	return player, nil
}

func (s *service) ModifyPlayer(id string, p map[string]interface{}) (*models.Player, *models.ErrorParsing) {

	player, err := s.repository.ModifyPlayer(id, p)

	if err != nil {
		return nil, &models.ErrorParsing{
			Error:      err,
			Type:       gin.ErrorTypePublic,
			Metadata:   http.StatusText(http.StatusNotFound),
			StatusCode: http.StatusNotFound,
		}
	}

	return player, nil
}

func (s *service) DeletePlayer() {}
