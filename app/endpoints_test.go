package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/Cside/jsondiff"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func (a *api) assertJSON(actual []byte, data interface{}, t *testing.T) {
	expected, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}

	if diff := jsondiff.Diff(expected, actual); diff != "" {
		t.Errorf("the expected json: %s is different from actual %s", expected, actual)
	}
}

func TestIndexHandler(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected when opening a stub databace connection", err)
	}
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	app := &api{sqlxDB}

	req, err := http.NewRequest("GET", "http://loclalhost/api/", nil)
	if err != nil {
		t.Fatalf("an error %s was not expected while creating request", err)
	}
	w := httptest.NewRecorder()

	// now := time.Now()
	rows := sqlmock.NewRows([]string{"id", "username", "email", "encryped_password", "avater"}).
	AddRow(1, "test1", "abc@gmail.com", sql.NullString{"test.password", true}, "testavatar").
	AddRow(2, "test2", "abc@gmail.com", sql.NullString{"test.password", true}, "testavatar")
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id,username,email,encryped_password,avater FROM users`)).WillReturnRows(rows)
	app.IndexHandler(w, req)

	if w.Code != 200 {
		t.Fatalf("expected status code to be 200, but got: %d", w.Code)
	}

	data := users{
		{ID: 1, UserName: "test1", Email: "abc@gmail.com", EncrypedPassword: sql.NullString{"test.password", true}, Avater: "testavatar"},
		{ID: 2, UserName: "test2", Email: "abc@gmail.com", EncrypedPassword: sql.NullString{"test.password", true}, Avater: "testavatar"},
	}

	app.assertJSON(w.Body.Bytes(), data, t)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
