package repositories

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func TestListPlayer(t *testing.T) {

	tests := []struct {
		name  string
		index int
		size  int
	}{
		{
			name:  "List Player Test 1",
			index: 1,
			size:  5,
		},
		{
			name:  "List Player Test 2",
			index: 1,
			size:  10,
		},
		{
			name:  "List Player Test 3",
			index: 5,
			size:  2,
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			mockDb, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

			if err != nil {
				t.Errorf("repositories.ListPlayer. \nError: failed to open a stud database connection | error = %v", err)
			}

			defer mockDb.Close()

			mock.ExpectBegin()
			mock.ExpectQuery("SELECT").
				WithArgs(test.size, test.index).
				WillReturnRows(mock.NewRows([]string{"id", "user_name", "pass_hash", "email", "display_name", "avatar_url", "avatar_blur_hash", "created_at", "updated_at", "deleted_at"}).
					AddRow("playerID", "userName", nil, nil, nil, nil, nil, time.Now(), nil, nil).
					AddRow("playerID", "userName", nil, nil, nil, nil, nil, time.Now(), nil, nil).
					AddRow("playerID", "userName", nil, nil, nil, nil, nil, time.Now(), nil, nil).
					AddRow("playerID", "userName", nil, nil, nil, nil, nil, time.Now(), nil, nil).
					AddRow("playerID", "userName", nil, nil, nil, nil, nil, time.Now(), nil, nil).
					AddRow("playerID", "userName", nil, nil, nil, nil, nil, time.Now(), nil, nil).
					AddRow("playerID", "userName", nil, nil, nil, nil, nil, time.Now(), nil, nil).
					AddRow("playerID", "userName", nil, nil, nil, nil, nil, time.Now(), nil, nil).
					AddRow("playerID", "userName", nil, nil, nil, nil, nil, time.Now(), nil, nil).
					AddRow("playerID", "userName", nil, nil, nil, nil, nil, time.Now(), nil, nil).
					AddRow("playerID", "userName", nil, nil, nil, nil, nil, time.Now(), nil, nil).
					AddRow("playerID", "userName", nil, nil, nil, nil, nil, time.Now(), nil, nil).
					AddRow("playerID", "userName", nil, nil, nil, nil, nil, time.Now(), nil, nil).
					AddRow("playerID", "userName", nil, nil, nil, nil, nil, time.Now(), nil, nil).
					AddRow("playerID", "userName", nil, nil, nil, nil, nil, time.Now(), nil, nil).
					AddRow("playerID", "userName", nil, nil, nil, nil, nil, time.Now(), nil, nil).
					AddRow("playerID", "userName", nil, nil, nil, nil, nil, time.Now(), nil, nil).
					AddRow("playerID", "userName", nil, nil, nil, nil, nil, time.Now(), nil, nil).
					AddRow("playerID", "userName", nil, nil, nil, nil, nil, time.Now(), nil, nil).
					AddRow("playerID", "userName", nil, nil, nil, nil, nil, time.Now(), nil, nil).
					AddRow("playerID", "userName", nil, nil, nil, nil, nil, time.Now(), nil, nil))

			repo := NewPlayerRepository(func() *gorm.DB {

				db, err := gorm.Open("postgres", mockDb)

				if err != nil {
					t.Errorf("utils.ListPlayer, \n Error: failed to open gorm database connection | error = %v", err)
				}

				return db
			})

			results, err := repo.ListPlayers(test.index, test.size)

			if err != nil {
				t.Errorf("utils.ListPlayer, \n Error: failed to get players | error = %v", err)
			}

			if results.CurrentItemCount != test.size {
				t.Errorf("utils.ListPlayer, \n Error: item counts result = %v | expected = %v", results.CurrentItemCount, test.size)
			}
		})
	}
}
