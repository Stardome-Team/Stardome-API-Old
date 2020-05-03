package database

import (
	"errors"
	"fmt"

	"github.com/Blac-Panda/Stardome-API/models"
	"github.com/Blac-Panda/Stardome-API/repositories"
	"github.com/jinzhu/gorm"
)

type repo struct {
	db *gorm.DB
}

// NewPlayerRepository :
func NewPlayerRepository(db *gorm.DB) repositories.PlayerRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) ListPlayers() {

}

func (r *repo) GetPlayer() {

}

func (r *repo) CreatePlayer() (*models.Player, error) {
	fmt.Println("\n\n\n Create Plater Repository \n\n\n.")

	return nil, errors.New("Not yet implemented")
}

func (r *repo) UpdatePlayer() {

}

func (r *repo) ModifyPlayer() {

}

func (r *repo) DeletePlayer() {

}
