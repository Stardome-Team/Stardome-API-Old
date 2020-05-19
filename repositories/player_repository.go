package repositories

import "github.com/Blac-Panda/Stardome-API/models"

// PlayerRepository :
type PlayerRepository interface {
	ListPlayers(index, size int) (*models.Pagination, error)
	GetPlayer(id string) (*models.Player, error)
	CreatePlayer(p *models.Player) (*models.Player, error)
	UpdatePlayer(p *models.Player) (*models.Player, error)
	ModifyPlayer(id string, p map[string]interface{}) (*models.Player, error)
	DeletePlayer(id string) error
}
