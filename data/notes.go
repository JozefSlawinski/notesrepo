package data

import (
	"database/sql"
	"strconv"
	"time"
)

type Note struct {
	Id       int
	Uuid     string
	Title    string
	Content  string
	Version  int
	Created  time.Time
	Modified time.Time
}

func (note *Note) CreatedAtDate() string {
	return note.Created.Format("Jan 2, 2006 at 3:04pm")
}

func (note *Note) ModifiedAtDate() string {
	return note.Modified.Format("Jan 2, 2006 at 3:04pm")
}

func (note *Note) VersionStr() string {
	return strconv.Itoa(note.Version)
}

// Get a note by the UUID
func NoteByUUID(uuid string) (conv Note, err error) {
	conv = Note{}
	err = Db.QueryRow("SELECT id, uuid, title, content, version, created, modified FROM notes WHERE uuid = $1", uuid).
		Scan(&conv.Id, &conv.Uuid, &conv.Title, &conv.Content, &conv.Version, &conv.Created, &conv.Modified)
	return
}

// Get a note by the UUID and version
func NoteByUUIDVersion(uuid, version string) (conv Note, err error) {
	return InnerNoteByUUIDVersion(Db, uuid, version)
}

func InnerNoteByUUIDVersion(database *sql.DB, uuid, version string) (conv Note, err error) {
	conv = Note{}
	err = database.QueryRow("SELECT id, uuid, title, content, version, created, modified FROM notes WHERE uuid = $1 AND version = $2", uuid, version).
		Scan(&conv.Id, &conv.Uuid, &conv.Title, &conv.Content, &conv.Version, &conv.Created, &conv.Modified)
	return
}

// Get the notes with highest version
func Notes() (notes []Note, err error) {
	return InnerNotes(Db)
}

func InnerNotes(database *sql.DB) (notes []Note, err error) {
	rows, err := database.Query("SELECT distinct on(uuid) id, uuid, title, content, version, created, modified FROM notes ORDER BY uuid, version DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Note{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Title, &conv.Content, &conv.Version, &conv.Created, &conv.Modified); err != nil {
			return
		}
		notes = append(notes, conv)
	}
	rows.Close()
	return
}

// Create a new note
func CreateNote(title, content string) (conv Note, err error) {
	return InnerCreateNote(Db, title, content)
}

func InnerCreateNote(database *sql.DB, title, content string) (conv Note, err error) {
	statement := "insert into notes (uuid, title, content, version, created, modified) values ($1, $2, $3, $4, $5, $6) returning id, uuid, title, content, version, created, modified"
	stmt, err := database.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), title, content, 1, time.Now(), time.Now()).Scan(&conv.Id, &conv.Uuid, &conv.Title, &conv.Content, &conv.Version, &conv.Created, &conv.Modified)
	return
}

// Edit the note
func EditNote(uuid, title, content, version string) (conv Note, err error) {
	return InnerEditNote(Db, uuid, title, content, version)
}

func InnerEditNote(database *sql.DB, uuid, title, content, version string) (conv Note, err error) {
	statement := "insert into notes (uuid, title, content, version, created, modified) values ($1, $2, $3, $4, $5, $6) returning id, uuid, title, content, version, created, modified"
	note, err := InnerNoteByUUIDVersion(database, uuid, version)
	if err != nil {
		return
	}

	stmt, err := database.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(uuid, title, content, note.Version+1, note.Created, time.Now()).Scan(&conv.Id, &conv.Uuid, &conv.Title, &conv.Content, &conv.Version, &conv.Created, &conv.Modified)
	return
}

// List the notes with previous versions
func (note *Note) NotesPreviousVersions() (notes []Note, err error) {
	return note.InnerNotesPreviousVersions(Db)
}

func (note *Note) InnerNotesPreviousVersions(database *sql.DB) (notes []Note, err error) {
	rows, err := database.Query("SELECT id, uuid, title, content, version, created, modified FROM notes WHERE uuid = $1  ORDER BY version DESC", note.Uuid)
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Note{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Title, &conv.Content, &conv.Version, &conv.Created, &conv.Modified); err != nil {
			return
		}
		notes = append(notes, conv)
	}
	rows.Close()
	return
}

//Delete the note
func DeleteNote(uuid string) (conv Note, err error) {
	return InnerDeleteNote(Db, uuid)
}

func InnerDeleteNote(database *sql.DB, uuid string) (conv Note, err error) {
	statement := "DELETE FROM notes WHERE uuid = $1"
	stmt, err := database.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(uuid)
	return
}
