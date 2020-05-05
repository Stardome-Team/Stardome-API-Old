package repositories

import "github.com/Blac-Panda/Stardome-API/models"

// PlayerRepository :
type PlayerRepository interface {
	ListPlayers() ([]*models.Player, error)
	GetPlayer(id string) (*models.Player, error)
	CreatePlayer(p *models.Player) (*models.Player, error)
	UpdatePlayer()
	ModifyPlayer()
	DeletePlayer()
}
