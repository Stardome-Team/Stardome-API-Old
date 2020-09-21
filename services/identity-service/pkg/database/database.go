package database

import (
	"fmt"

	"github.com/Stardome-Team/Stardome-API/services/identity-service/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB represents a DB connection that can be used to run SQL queries.
type DB struct {
	db *gorm.DB
}

// New returns a new DB connection that wraps the given gorm.DB instance.
func New(db *gorm.DB) *DB {
	return &DB{db: db}
}

// DB returns the gorm.DB wrapped by this object.
func (db *DB) DB() *gorm.DB {
	return db.db
}

// OpenConnection creates a connection to the database
func OpenConnection(cfg *config.Config) (*DB, error) {
	var dns string = fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", cfg.Database.Host, cfg.Database.Port, cfg.Database.Name, cfg.Database.Username, cfg.Database.Password)

	connection, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	db := New(connection)

	return db, nil
}
