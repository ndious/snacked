package handlers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
)

type Router struct {
    router *http.ServeMux
    db *sqlx.DB
}

func handleError(err error, w http.ResponseWriter) {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(err.Error()))
}

