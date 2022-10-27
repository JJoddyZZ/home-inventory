package repository

import (
	"database/sql"
	"errors"
	"testing"

	m "github.com/JJoddyZZ/home-inventory/internal/repository/mock"
	"github.com/stretchr/testify/assert"
)

func Test_NewStorage(t *testing.T) {
	st := NewStorage(&sql.DB{})
	assert.NotNil(t, st)
}

func Test_Ping(t *testing.T) {
	t.Run("ping returns error", func(t *testing.T) {
		db, mock := m.NewSQLMockMonitoringPing()
		mock.ExpectPing().WillReturnError(errors.New("some error"))
		repo := &Storage{
			client: db,
		}
		err := repo.Ping()
		assert.Error(t, err)
	})
	t.Run("ping succeeds", func(t *testing.T) {
		db, mock := m.NewSQLMockMonitoringPing()
		mock.ExpectPing()
		repo := &Storage{
			client: db,
		}
		err := repo.Ping()
		assert.NoError(t, err)
	})
}
