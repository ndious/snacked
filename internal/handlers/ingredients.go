package handlers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	cmp "github.com/ndious/snacked/internal/components/ingredients"
	"github.com/ndious/snacked/internal/services"
)

func IngredientsRouter(r *http.ServeMux, db *sqlx.DB) (*Router, error) {

    h := IngredientHandler{
        is: services.NewIngredientService(db),
    }

    r.HandleFunc("GET /ingredients", h.Search())

    return &Router{
        r,
        db,
    }, nil
}

type IngredientHandler struct {
    is *services.IngredientService
}

func (self IngredientHandler) Search() func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        searh := r.URL.Query().Get("name")
        ingredients, err := self.is.Search(searh)

        if (err != nil) {
            handleError(err, w)
            return
        }

        cmp.SearchList(ingredients, "0").Render(r.Context(), w)
    }
}

