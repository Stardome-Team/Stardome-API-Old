package repositories

import (
	"github.com/Blac-Panda/Stardome-API/models"
	"github.com/Blac-Panda/Stardome-API/utils"
	"github.com/fatih/structs"
	"github.com/jinzhu/gorm"
)

// PlayerRepository :
type PlayerRepository interface {
	ListPlayers(index, size int) (*models.Pagination, error)
	GetPlayer(id string) (*models.Player, error)
	GetPlayerByUserName(username string) (*models.Player, error)
	CreatePlayer(p *models.Player) (*models.Player, error)
	UpdatePlayer(p *models.Player) (*models.Player, error)
	ModifyPlayer(id string, p map[string]interface{}) (*models.Player, error)
	DeletePlayer(id string) error
}
type repo struct {
	db func() *gorm.DB
}

// NewPlayerRepository :
func NewPlayerRepository(db func() *gorm.DB) PlayerRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) ListPlayers(index, size int) (*models.Pagination, error) {
	var db *gorm.DB = r.db()
	defer db.Close()

	if db == nil {
		return nil, utils.ErrorInternalError
	}

	var players []*models.Player = []*models.Player{}

	var count int = 0

	db.Offset((index - 1) * size).Limit(size).Find(&players).Count(&count)

	return &models.Pagination{
		StartIndex:       &index,
		ItemsPerPage:     &size,
		TotalItems:       &count,
		CurrentItemCount: len(players),
		Items:            &players,
	}, nil
}

func (r *repo) GetPlayer(id string) (*models.Player, error) {
	var db *gorm.DB = r.db()

	if db == nil {
		return nil, utils.ErrorInternalError
	}

	var player *models.Player = &models.Player{}

	var count int = 0

	if db.Where("id = ?", id).First(&player).Count(&count); count == 0 {
		return nil, utils.ErrorPlayerNotFound
	}

	return player, nil
}

func (r *repo) GetPlayerByUserName(username string) (*models.Player, error) {
	var db *gorm.DB = r.db()

	if db == nil {
		return nil, utils.ErrorInternalError
	}

	var player *models.Player = &models.Player{}

	var count int = 0

	if db.Where("user_name = ?", username).First(&player).Count(&count); count == 0 {
		return nil, utils.ErrorPlayerNotFound
	}

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

	var player *models.Player = &models.Player{}

	db.Where("id = ?", p.ID).First(&player)

	return player, nil
}

func (r *repo) UpdatePlayer(p *models.Player) (*models.Player, error) {
	var db *gorm.DB = r.db()
	defer db.Close()

	if db == nil {
		// TODO: Log Errors
		return nil, utils.ErrorInternalError
	}

	if db.First(&models.Player{}, " id = ? ", &p.ID).RecordNotFound() {
		// TODO: Log Errors
		return nil, utils.ErrorPlayerNotFound
	}

	mp := structs.New(p)

	if rows := db.Model(&models.Player{}).Updates(mp.Map()).RowsAffected; rows == 0 {
		return nil, utils.ErrorPlayerUpdateFailed
	}

	var player *models.Player = &models.Player{}

	db.Where("id = ?", p.ID).First(&player)

	return player, nil
}

func (r *repo) ModifyPlayer(id string, p map[string]interface{}) (*models.Player, error) {
	var db *gorm.DB = r.db()
	defer db.Close()

	if db == nil {
		// TODO: Log Errors
		return nil, utils.ErrorInternalError
	}

	if db.First(&models.Player{}, " id = ? ", id).RecordNotFound() {
		// TODO: Log Errors
		return nil, utils.ErrorPlayerNotFound
	}

	if rows := db.Model(&models.Player{}).Omit("id", "user_name", "pass_hash", "created_at").Where(" id = ? ", id).Updates(p).RowsAffected; rows == 0 {
		return nil, utils.ErrorPlayerUpdateFailed
	}

	var player *models.Player = &models.Player{}

	db.Where("id = ?", id).First(&player)

	return player, nil
}

func (r *repo) DeletePlayer(id string) error {
	var db *gorm.DB = r.db()
	defer db.Close()

	if db == nil {
		// TODO: Log Errors
		return utils.ErrorInternalError
	}

	if db.First(&models.Player{}, " id = ? ", id).RecordNotFound() {
		// TODO: Log Errors
		return utils.ErrorPlayerNotFound
	}

	if rows := db.Where("id = ?", id).Delete(&models.Player{}).RowsAffected; rows == 0 {
		return utils.ErrorPlayerDeleteFailed
	}

	return nil
}
