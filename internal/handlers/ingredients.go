package handlers

import (
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	cmp "github.com/ndious/snacked/internal/components/ingredients"
	"github.com/ndious/snacked/internal/models"
	"github.com/ndious/snacked/internal/services"
)

func IngredientsRouter(r *http.ServeMux, db *sqlx.DB) (*Router, error) {

	h := IngredientHandler{
		is: services.NewIngredientService(db),
	}

	r.HandleFunc("GET /ingredients", h.Search())

	r.HandleFunc("POST /recipes/{recipe_id}/ingredients", h.Create())

	return &Router{
		r,
		db,
	}, nil
}

type IngredientHandler struct {
	is *services.IngredientService
}

// Search returns a handler function that searches for ingredients based on the provided name.
//
// It takes in the http.ResponseWriter and http.Request as parameters.
// It returns a function that takes in the http.ResponseWriter and http.Request as parameters.
func (ih IngredientHandler) Search() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("name")
		ingredients, err := ih.is.Search(search)

		if err != nil {
			handleError(err, w)
			return
		}

		cmp.SearchList(ingredients, "0", search).Render(r.Context(), w)
	}
}

// Create handles the creation of a new ingredient.
//
// It takes a writer and a request as parameters and returns a function.
func (ih IngredientHandler) Create() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		recipeId, err := strconv.Atoi(r.PathValue("recipe_id"))
		if err != nil {
			handleError(err, w)
			return
		}

		name := r.FormValue("name")
		ingredient, err := ih.is.CreateIngredient(name, models.Gram)

		if err != nil {
			handleError(err, w)
			return
		}

		recipeIngredient, err := ih.is.BuildRecipeIngredient(recipeId, ingredient.ID)
		if err != nil {
			handleError(err, w)
			return
		}

		cmp.FormCreate(recipeIngredient).Render(r.Context(), w)
	}
}
