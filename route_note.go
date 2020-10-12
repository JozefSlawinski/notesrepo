package main

import (
	"net/http"

	"./data"
)

func newNote(writer http.ResponseWriter, request *http.Request) {

	generateHTML(writer, nil, "layout", "navbar", "new.note")

}

// POST /note/create
// Create the user account
func createNote(writer http.ResponseWriter, request *http.Request) {
	title := request.PostFormValue("title")
	content := request.PostFormValue("content")
	if _, err := data.CreateNote(title, content); err != nil {
		danger(err, "Cannot create note")
	}
	http.Redirect(writer, request, "/", 302)
}

// GET /thread/read
// Show the details of the thread, including the posts and the form to write a post
func readNote(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	version := vals.Get("version")
	note, err := data.NoteByUUIDVersion(uuid, version)
	if err != nil {
		error_message(writer, request, "Cannot read note")
	} else {
		generateHTML(writer, &note, "layout", "navbar", "note")

	}
}

func deleteNote(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	_, err := data.DeleteNote(uuid)
	if err != nil {
		error_message(writer, request, "Cannot delete note")
	} else {
		index(writer, request)

	}
}

// POST /signup
// Create the user account
func editNote(writer http.ResponseWriter, request *http.Request) {
	uuid := request.PostFormValue("uuid")
	title := request.PostFormValue("title")
	content := request.PostFormValue("content")
	version := request.PostFormValue("version")
	if _, err := data.EditNote(uuid, title, content, version); err != nil {
		danger(err, "Cannot edit note")
	}
	http.Redirect(writer, request, "/", 302)
}
