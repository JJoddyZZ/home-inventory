package mock

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Ping(arg0 *sql.DB) error {
	args := m.Called(arg0)
	return args.Error(0)
}
