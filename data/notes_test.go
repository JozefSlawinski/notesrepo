package data

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	// "gitlab.mfb.io/user/graphql_server/models"
)

// func TestInit(t *testing.T) {
// 	var err error
// 	Db, err = sql.Open("postgres", "dbname=noteas user=postgres password=123 port=5432 sslmode=disable")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return
// }

func TestInnerCreateNote(t *testing.T) {

	// Creates sqlmock database connection and a mock to manage expectations.
	database, mock, err := sqlmock.New()

	dateCreated := time.Now()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Closes the database and prevents new queries from starting.

	// Here we are creating rows in our mocked database.
	rows := sqlmock.NewRows([]string{"id", "uuid", "title", "content", "version", "created", "modified"}).
		AddRow(1, "4db2668e-bc73-4260-67d8-330bf33dbc86", "testTitle", "testContent", 1, dateCreated, dateCreated)

	mock.ExpectPrepare("insert into notes .* returning .*").WillBeClosed()
	mock.ExpectQuery("insert into notes .* returning .*").
		WithArgs(sqlmock.AnyArg(), "testTitle", "testContent", 1, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(rows)

	note, err := InnerCreateNote(database, "testTitle", "testContent")

	if err != nil {
		t.Errorf("error '%s' ", err)
	}
	defer database.Close()
	// Here we just construction our expecting result.
	noteExp := Note{
		Id:       1,
		Uuid:     "4db2668e-bc73-4260-67d8-330bf33dbc86",
		Title:    "testTitle",
		Content:  "testContent",
		Version:  1,
		Created:  dateCreated,
		Modified: dateCreated,
	}

	assert.Equal(t, noteExp, note)
}

func TestInnerEditNote(t *testing.T) {

	// Creates sqlmock database connection and a mock to manage expectations.
	database, mock, err := sqlmock.New()

	dateCreated := time.Now()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	// Closes the database and prevents new queries from starting.
	rows := sqlmock.NewRows([]string{"id", "uuid", "title", "content", "version", "created", "modified"}).
		AddRow(1, "4db2668e-bc73-4260-67d8-330bf33dbc86", "testTitle", "testContent", 1, dateCreated, dateCreated)

	mock.ExpectQuery("SELECT .* FROM notes WHERE .*").
		WithArgs("4db2668e-bc73-4260-67d8-330bf33dbc86", "1").
		WillReturnRows(rows)

	rowsEdited := sqlmock.NewRows([]string{"id", "uuid", "title", "content", "version", "created", "modified"}).
		AddRow(1, "4db2668e-bc73-4260-67d8-330bf33dbc86", "testTitle1", "testContent", 1, dateCreated, dateCreated)

	mock.ExpectPrepare("insert into notes .* returning .*").WillBeClosed()
	mock.ExpectQuery("insert into notes .* returning .*").
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(rowsEdited)

	note, err := InnerEditNote(database, "4db2668e-bc73-4260-67d8-330bf33dbc86", "testTitle1", "testContent", "1")

	if err != nil {
		t.Errorf("error '%s' ", err)
	}
	defer database.Close()
	// Here we just construction our expecting result.
	noteExp := Note{
		Id:       1,
		Uuid:     "4db2668e-bc73-4260-67d8-330bf33dbc86",
		Title:    "testTitle1",
		Content:  "testContent",
		Version:  1,
		Created:  dateCreated,
		Modified: dateCreated,
	}

	assert.Equal(t, noteExp, note)
}
