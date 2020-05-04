package repositories

import "github.com/Blac-Panda/Stardome-API/models"

// PlayerRepository :
type PlayerRepository interface {
	ListPlayers()
	GetPlayer()
	CreatePlayer(p *models.Player) (*models.Player, error)
	UpdatePlayer()
	ModifyPlayer()
	DeletePlayer()
}
