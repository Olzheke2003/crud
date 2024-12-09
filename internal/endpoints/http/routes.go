package http

import (
	"crud/internal/endpoints/http/create"
	"crud/internal/endpoints/http/delete"
	"crud/internal/endpoints/http/read"
	"crud/internal/endpoints/http/update"
	"net/http"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/create", create.Handler())
	mux.HandleFunc("/read/", read.Handler())
	mux.HandleFunc("/update/", update.Handler())
	mux.HandleFunc("/delete/", delete.Handler())
	return mux
}
