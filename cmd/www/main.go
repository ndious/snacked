package main

import (
	"net/http"
	"os"
	"path/filepath"

	cmp "github.com/ndious/snacked/internal/components"
	"github.com/ndious/snacked/internal/components/layout"
	"github.com/ndious/snacked/internal/database"
)

func main() {
	exepath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	os.Setenv("BASEDIR", filepath.Dir(exepath))

	mux := http.NewServeMux()

	mux.HandleFunc("GET /migrate", func(w http.ResponseWriter, r *http.Request) {
		migrations, err := database.Migrate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cmp.Migrations(migrations).Render(r.Context(), w)
	})

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		c := cmp.Hello("Dious")
		layout.Page(c).Render(r.Context(), w)
	})

	http.ListenAndServe(":1337", mux)
}
