package mock

import (
	"database/sql"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"testing"

	"github.com/jmoiron/sqlx"
)

func MockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *sqlx.DB) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("mock database connection error:%s", err)
	}

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	return mockDB, mock, sqlxDB
}
