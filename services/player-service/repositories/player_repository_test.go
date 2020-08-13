package repositories

import (
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/Blac-Panda/Stardome-API/services/player-service/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

type playertest struct {
	description string
	player      models.Player
}

func generatePlayerTests(size int, desc string) []playertest {

	tests := []playertest{}

	for i := 0; i < size; i++ {

		description := strings.Join([]string{desc, "Test", string(i)}, " ")
		id := string(i)
		userName := strings.Join([]string{"player", string(i)}, "_")
		passHash := "pass"
		emailAddress := "test@test.com"
		displayName := strings.Join([]string{"player", string(i)}, "_")
		avatarURL := "URL"
		avatarBlurHash := "hash"
		createdAt := time.Now()

		tests = append(tests, playertest{
			description: description,
			player: models.Player{
				ID:             &id,
				UserName:       &userName,
				PassHash:       &passHash,
				EmailAddress:   &emailAddress,
				DisplayName:    &displayName,
				AvatarURL:      &avatarURL,
				AvatarBlurHash: &avatarBlurHash,
				CreatedAt:      &createdAt,
				UpdatedAt:      nil,
				DeletedAt:      nil,
			},
		})
	}

	return tests
}

func TestGetPlayer(t *testing.T) {

	var tests []playertest = generatePlayerTests(5, "repository.GetPlayer")

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			mockDB, mock, err := sqlmock.New()

			if err != nil {
				t.Errorf("repositories.GetPlayer. \nError: failed to open a stud database connection | error = %v", err)
			}

			defer mockDB.Close()

			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "players"  WHERE "players"."deleted_at" IS NULL AND ((id = $1)) ORDER BY "players"."id" ASC LIMIT 1`)).
				WithArgs(string(*test.player.ID)).
				WillReturnRows(mock.NewRows([]string{"id", "user_name", "pass_hash", "email", "display_name", "avatar_url", "avatar_blur_hash", "created_at", "updated_at", "deleted_at"}).
					AddRow(test.player.ID, test.player.UserName, test.player.PassHash, test.player.EmailAddress, test.player.DisplayName, test.player.AvatarURL, test.player.AvatarBlurHash, test.player.CreatedAt, test.player.UpdatedAt, test.player.DeletedAt))

			mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "players"  WHERE "players"."deleted_at" IS NULL AND "players"."id" = $1 AND ((id = $2))`)).
				WithArgs(string(*test.player.ID), string(*test.player.ID)).
				WillReturnRows(mock.NewRows([]string{"count"}).
					AddRow(1))

			mock.ExpectBegin()
			mock.ExpectCommit()
			mock.ExpectClose()

			repo := NewPlayerRepository(func() *gorm.DB {

				db, err := gorm.Open("postgres", mockDB)

				if err != nil {
					t.Errorf("repositories.GetPlayerByUserName, \n Error: failed to open gorm database connection | error = %v", err)
				}
				return db
			})

			player, err := repo.GetPlayer(*test.player.ID)

			if err != nil {
				t.Error("An error occurred while getting player")
			}

			if player == nil {
				t.Error("An error occurred while getting player")
			}
		})
	}
}

func TestGetPlayerByUserName(t *testing.T) {

	var tests []playertest = generatePlayerTests(5, "repositories.GetPlayerByUserName")

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			mockDB, mock, err := sqlmock.New()

			if err != nil {
				t.Errorf("repositories.GetPlayerByUserName. \nError: failed to open a stud database connection | error = %v", err)
			}

			defer mockDB.Close()

			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "players"  WHERE "players"."deleted_at" IS NULL AND ((user_name = $1)) ORDER BY "players"."id" ASC LIMIT 1`)).
				WithArgs(string(*test.player.UserName)).
				WillReturnRows(mock.NewRows([]string{"id", "user_name", "pass_hash", "email", "display_name", "avatar_url", "avatar_blur_hash", "created_at", "updated_at", "deleted_at"}).
					AddRow(test.player.ID, test.player.UserName, test.player.PassHash, test.player.EmailAddress, test.player.DisplayName, test.player.AvatarURL, test.player.AvatarBlurHash, test.player.CreatedAt, test.player.UpdatedAt, test.player.DeletedAt))

			mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "players"  WHERE "players"."deleted_at" IS NULL AND "players"."id" = $1 AND ((user_name = $2))`)).
				WithArgs(string(*test.player.ID), string(*test.player.UserName)).
				WillReturnRows(mock.NewRows([]string{"count"}).
					AddRow(1))

			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "players"  WHERE "players"."deleted_at" IS NULL AND ((user_name = $1)) ORDER BY "players"."id" ASC LIMIT 1`)).
				WithArgs(string(*test.player.UserName)).
				WillReturnRows(mock.NewRows([]string{"id", "user_name", "pass_hash", "email", "display_name", "avatar_url", "avatar_blur_hash", "created_at", "updated_at", "deleted_at"}).
					AddRow(test.player.ID, test.player.UserName, test.player.PassHash, test.player.EmailAddress, test.player.DisplayName, test.player.AvatarURL, test.player.AvatarBlurHash, test.player.CreatedAt, test.player.UpdatedAt, test.player.DeletedAt))

			mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "players"  WHERE "players"."deleted_at" IS NULL AND "players"."id" = $1 AND ((user_name = $2))`)).
				WithArgs(string(*test.player.ID), string(*test.player.UserName)).
				WillReturnRows(mock.NewRows([]string{"count"}).
					AddRow(1))

			mock.ExpectBegin()
			mock.ExpectCommit()
			mock.ExpectClose()

			repo := NewPlayerRepository(func() *gorm.DB {

				db, err := gorm.Open("postgres", mockDB)

				if err != nil {
					t.Errorf("repositories.GetPlayerByUserName, \n Error: failed to open gorm database connection | error = %v", err)
				}
				return db
			})

			player, err := repo.GetPlayerByUserName(*test.player.UserName)

			if err != nil {
				t.Error("repositories.GetPlayerByUserName, An error occurred while getting player")
			}

			if player == nil {
				t.Error("repositories.GetPlayerByUserName, An error occurred while getting player")
			}
		})
	}
}

func TestCreatePlayer(t *testing.T) {

	var tests []playertest = generatePlayerTests(5, "repositories.CreatePlayer")

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			mockDB, mock, err := sqlmock.New()

			if err != nil {
				t.Errorf("repositories.CreatePlayer. \nError: failed to open a stud database connection | error = %v", err)
			}
			defer mockDB.Close()

			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "players" WHERE "players"."deleted_at" IS NULL AND (( user_name = $1 )) ORDER BY "players"."id" ASC LIMIT 1`)).
				WithArgs(string(*test.player.UserName)).
				WillReturnRows(mock.NewRows([]string{"id", "user_name", "pass_hash", "email", "display_name", "avatar_url", "avatar_blur_hash", "created_at", "updated_at", "deleted_at"}))

			mock.ExpectBegin()

			mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "players" ("id","user_name","pass_hash","email","display_name","avatar_url","avatar_blur_hash","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING "players"."id"`)).
				WithArgs(&test.player.ID, &test.player.UserName, &test.player.PassHash, &test.player.EmailAddress, &test.player.DisplayName, &test.player.AvatarURL, &test.player.AvatarBlurHash, &test.player.CreatedAt, &test.player.UpdatedAt, &test.player.DeletedAt).
				WillReturnRows(mock.NewRows([]string{"id"}).
					AddRow(test.player.ID))

			mock.ExpectCommit()

			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "players"  WHERE "players"."deleted_at" IS NULL AND ((id = $1)) ORDER BY "players"."id" ASC LIMIT 1`)).
				WithArgs(string(*test.player.ID)).
				WillReturnRows(mock.NewRows([]string{"id", "user_name", "pass_hash", "email", "display_name", "avatar_url", "avatar_blur_hash", "created_at", "updated_at", "deleted_at"}).
					AddRow(test.player.ID, test.player.UserName, test.player.PassHash, test.player.EmailAddress, test.player.DisplayName, test.player.AvatarURL, test.player.AvatarBlurHash, test.player.CreatedAt, test.player.UpdatedAt, test.player.DeletedAt))

			mock.ExpectClose()

			repo := NewPlayerRepository(func() *gorm.DB {

				db, err := gorm.Open("postgres", mockDB)

				if err != nil {
					t.Errorf("repositories.CreatePlayer. \nError: failed to open a gorm database connection | error = %v", err)
				}
				return db
			})

			player, err := repo.CreatePlayer(&test.player)

			if err != nil {
				t.Errorf("repositories.CreatePlayer. \nError: create player failed | error = %v", err)
			}

			if player == nil {
				t.Errorf("repositories.CreatePlayer. \nError: create player failed")

			}
		})
	}
}
