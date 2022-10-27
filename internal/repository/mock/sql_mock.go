package mock

import (
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
)

func NewSQLMockMonitoringPing() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true)) // activating ping monitor so ping() can be traced
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}
