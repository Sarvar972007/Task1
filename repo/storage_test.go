package repo

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUserById_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer func() {
		if err := db.Close(); err != nil {
			t.Errorf("Error closing mock database: %s", err)
		}
	}()

	storage := &Storage{Db: db}

	rows := sqlmock.NewRows([]string{"id", "username", "email"}).
		AddRow(1, "sarvar", "sarvar@example.com")

	mock.ExpectQuery("SELECT id, username, email FROM users WHERE id = 1 \\$1").
		WithArgs(1).
		WillReturnRows(rows)

	user, err := storage.GetUserById(1)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, 1, user.Id)
	assert.Equal(t, "sarvar", user.UserName)
	assert.Equal(t, "sarvar@example.com", user.Email)

	assert.NoError(t, mock.ExpectationsWereMet())
}
