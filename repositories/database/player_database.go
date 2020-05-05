package database

import (
	"github.com/Blac-Panda/Stardome-API/models"
	"github.com/Blac-Panda/Stardome-API/repositories"
	"github.com/Blac-Panda/Stardome-API/utils"
	"github.com/jinzhu/gorm"
)

type repo struct {
	db func() *gorm.DB
}

// NewPlayerRepository :
func NewPlayerRepository(db func() *gorm.DB) repositories.PlayerRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) ListPlayers() ([]*models.Player, error) {
	var db *gorm.DB = r.db()
	defer db.Close()

	if db == nil {
		return nil, utils.ErrorInternalError
	}

	var players []*models.Player = []*models.Player{}

	db.Find(&players)

	return players, nil
}

func (r *repo) GetPlayer(id string) (*models.Player, error) {
	var db *gorm.DB = r.db()

	if db == nil {
		return nil, utils.ErrorInternalError
	}

	var player *models.Player = &models.Player{}

	db.First(&player)

	return player, nil
}

func (r *repo) CreatePlayer(p *models.Player) (*models.Player, error) {
	var db *gorm.DB = r.db()
	defer db.Close()

	if db == nil {
		// TODO: Log Errors
		return nil, utils.ErrorInternalError
	}

	if !db.First(&models.Player{}, " user_name = ? ", &p.UserName).RecordNotFound() {
		return nil, utils.ErrorPlayerAlreadyExist
	}

	if err := db.Create(p).Error; err != nil {
		// TODO: Log Errors
		return nil, utils.ErrorPlayerCreationFailed
	}

	return p, nil
}

func (r *repo) UpdatePlayer() {

}

func (r *repo) ModifyPlayer() {

}

func (r *repo) DeletePlayer() {

}
