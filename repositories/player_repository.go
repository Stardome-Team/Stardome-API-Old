package repositories

import "github.com/Blac-Panda/Stardome-API/models"

// PlayerRepository :
type PlayerRepository interface {
	ListPlayers()
	GetPlayer()
	CreatePlayer() (*models.Player, error)
	UpdatePlayer()
	ModifyPlayer()
	DeletePlayer()
}
