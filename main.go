package main

import (
	"net/http"
	"time"

	"./data"
)

func main() {
	p("NotesRepo", version(), "started at", config.Address)

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	data.ConnectDb()
	//
	// all route patterns matched here
	// route handler functions defined in other files
	//

	// index
	mux.HandleFunc("/", index)

	// // defined in route_note.go
	mux.HandleFunc("/note/new", newNote)
	mux.HandleFunc("/note/create", createNote)
	mux.HandleFunc("/note/delete", deleteNote)
	mux.HandleFunc("/note/edit", editNote)
	mux.HandleFunc("/note/read", readNote)

	// starting up the server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}

func index(writer http.ResponseWriter, request *http.Request) {
	notes, err := data.Notes()
	if err != nil {
		error_message(writer, request, "Cannot get notes")
	} else {

		generateHTML(writer, notes, "layout", "navbar", "index")

	}
}

// GET /err?msg=
// shows the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()

	generateHTML(writer, vals.Get("msg"), "layout", "navbar", "error")

}
