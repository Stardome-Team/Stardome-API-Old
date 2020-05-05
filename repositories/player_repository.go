package repositories

import "github.com/Blac-Panda/Stardome-API/models"

// PlayerRepository :
type PlayerRepository interface {
	ListPlayers() ([]*models.Player, error)
	GetPlayer()
	CreatePlayer(p *models.Player) (*models.Player, error)
	UpdatePlayer()
	ModifyPlayer()
	DeletePlayer()
}
