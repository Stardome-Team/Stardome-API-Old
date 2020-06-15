package repositories

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

// func TestListPlayer(t *testing.T) {

// 	tests := []struct {
// 		name  string
// 		index int
// 		size  int
// 	}{
// 		{
// 			name:  "List Player Test 1",
// 			index: 1,
// 			size:  5,
// 		},
// 		// {
// 		// 	name:  "List Player Test 2",
// 		// 	index: 1,
// 		// 	size:  10,
// 		// },
// 		// {
// 		// 	name:  "List Player Test 3",
// 		// 	index: 5,
// 		// 	size:  2,
// 		// },
// 	}

// 	for _, test := range tests {

// 		t.Run(test.name, func(t *testing.T) {
// 			mockDb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

// 			if err != nil {
// 				t.Errorf("repositories.ListPlayer. \nError: failed to open a stud database connection | error = %v", err)
// 			}

// 			defer mockDb.Close()

// 			mock.ExpectQuery(`SELECT`).
// 				// WithArgs(string(test.size), string(((test.index - 1) * test.size))).
// 				WillReturnRows(mock.NewRows([]string{"id", "deleted_at"}).
// 					AddRow("playerID", nil).
// 					AddRow("playerID", nil).
// 					AddRow("playerID", nil).
// 					AddRow("playerID", nil).
// 					AddRow("playerID", nil).
// 					AddRow("playerID", nil).
// 					AddRow("playerID", nil).
// 					AddRow("playerID", nil).
// 					AddRow("playerID", nil).
// 					AddRow("playerID", nil).
// 					AddRow("playerID", nil).
// 					AddRow("playerID", nil).
// 					AddRow("playerID", nil).
// 					AddRow("playerID", nil).
// 					AddRow("playerID", nil).
// 					AddRow("playerID", nil).
// 					AddRow("playerID", nil).
// 					AddRow("playerID", nil).
// 					AddRow("playerID", nil).
// 					AddRow("playerID", nil).
// 					AddRow("playerID", nil).
// 					AddRow("playerID", nil).
// 					AddRow("playerID", nil).
// 					AddRow("playerID", nil).
// 					AddRow("playerID", nil))

// 			mock.ExpectBegin()
// 			mock.ExpectCommit()

// 			repo := NewPlayerRepository(func() *gorm.DB {

// 				db, err := gorm.Open("postgres", mockDb)

// 				if err != nil {
// 					t.Errorf("utils.ListPlayer, \n Error: failed to open gorm database connection | error = %v", err)
// 				}

// 				return db
// 			})

// 			results, err := repo.ListPlayers(test.index, test.size)

// 			if err != nil {
// 				t.Errorf("utils.ListPlayer, \n Error: failed to get players | error = %v", err)
// 			}

// 			if results.CurrentItemCount != test.size {
// 				t.Errorf("utils.ListPlayer, \n Error: item counts result = %v ( count = %v | expected = %v)", results.Items, results.CurrentItemCount, test.size)
// 			}
// 		})
// 	}
// }

func TestGetPlayer(t *testing.T) {
	tests := []struct {
		name string
		ID   string
	}{
		{
			name: "Get Player Test 1",
			ID:   "1",
		},
		{
			name: "Get Player Test 2",
			ID:   "2",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockDB, mock, err := sqlmock.New()

			if err != nil {
				t.Errorf("repositories.GetPlayer. \nError: failed to open a stud database connection | error = %v", err)
			}

			defer mockDB.Close()

			mock.MatchExpectationsInOrder(false)
			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "players"  WHERE "players"."deleted_at" IS NULL AND ((id = $1)) ORDER BY "players"."id" ASC LIMIT 1`)).
				WithArgs(string(test.ID)).
				WillReturnRows(mock.NewRows([]string{"id", "user_name", "pass_hash", "email", "display_name", "avatar_url", "avatar_blur_hash", "created_at", "updated_at", "deleted_at"}).
					AddRow(test.ID, "", "", "", "", "", "", time.Now(), nil, nil))

			mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "players"  WHERE "players"."deleted_at" IS NULL AND "players"."id" = $1 AND ((id = $2))`)).
				WithArgs(string(test.ID), string(test.ID)).
				WillReturnRows(mock.NewRows([]string{"count"}).
					AddRow(1))

			mock.ExpectBegin()
			mock.ExpectCommit()
			mock.ExpectClose()

			repo := NewPlayerRepository(func() *gorm.DB {

				db, err := gorm.Open("postgres", mockDB)

				if err != nil {
					t.Errorf("utils.ListPlayer, \n Error: failed to open gorm database connection | error = %v", err)
				}
				return db
			})

			player, err := repo.GetPlayer(test.ID)

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

}
