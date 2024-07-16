package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	cmp "github.com/ndious/snacked/internal/components/ingredients"
	"github.com/ndious/snacked/internal/models"
	"github.com/ndious/snacked/internal/services"
)

func RecipeIngredientsRouter(r *http.ServeMux, db *sqlx.DB) (*Router, error) {

	h := RecipeIngredientHandler{
		rs: services.NewRecipeService(db),
		is: services.NewIngredientService(db),
	}

	r.HandleFunc("GET /recipes/{recipe_id}/ingredients/new", h.FormSearch())

	r.HandleFunc("GET /recipes/{recipe_id}/ingredients/search", h.Search())

	r.HandleFunc("GET /recipes/{recipe_id}/ingredients/{ingredient_id}/new", h.New())

	r.HandleFunc("POST /recipes/{recipe_id}/ingredients/{ingredient_id}", h.Create())

	r.HandleFunc("PUT /recipes/{recipe_id}/ingredients/{ingredient_id}", h.Update())

	r.HandleFunc("DELETE /recipes/{recipe_id}/ingredients/{ingredient_id}", h.Delete())

	return &Router{
		r,
		db,
	}, nil
}

type RecipeIngredientHandler struct {
	rs *services.RecipeService
	is *services.IngredientService
}

func (rih RecipeIngredientHandler) getRecipeIngredientIds(r *http.Request) (recipeId, ingredientId int, err error) {
	var err1, err2 error

	recipeId, err1 = strconv.Atoi(r.PathValue("recipe_id"))

	ingredientId, err2 = strconv.Atoi(r.PathValue("ingredient_id"))

	return recipeId, ingredientId, errors.Join(err1, err2)
}

func (rih RecipeIngredientHandler) FormSearch() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		recipeId, err := strconv.Atoi(r.PathValue("recipe_id"))
		if err != nil {
			handleError(err, w)
			return
		}

		ingredient := models.RecipeIngredient{
			RecipeID: recipeId,
		}

		cmp.FormNew(ingredient).Render(r.Context(), w)
	}
}

func (rih RecipeIngredientHandler) New() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		recipeId, ingredientId, err := rih.getRecipeIngredientIds(r)
		if err != nil {
			handleError(err, w)
			return
		}

		ingredient, err := rih.is.BuildRecipeIngredient(recipeId, ingredientId)
		if err != nil {
			handleError(err, w)
			return
		}

		cmp.FormCreate(ingredient).Render(r.Context(), w)
	}
}

func (rih RecipeIngredientHandler) Create() func(w http.ResponseWriter, r *http.Request) {
	return rih.handleSave(func(ingredient models.RecipeIngredient) (models.RecipeIngredient, error) {
		fmt.Println("Create", ingredient)
		return rih.is.Create(ingredient)
	})
}

func (rih RecipeIngredientHandler) Update() func(w http.ResponseWriter, r *http.Request) {
	return rih.handleSave(func(ingredient models.RecipeIngredient) (models.RecipeIngredient, error) {
		fmt.Println("Update", ingredient)
		return rih.is.Update(ingredient)
	})
}

func (rih RecipeIngredientHandler) handleSave(save func(ingredient models.RecipeIngredient) (models.RecipeIngredient, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		recipeId, ingredientId, err := rih.getRecipeIngredientIds(r)
		if err != nil {
			handleError(err, w)
			return
		}

		quantity, err := rih.getQuantity(r)
		if err != nil {
			handleError(err, w)
			return
		}

		ingredient, err := rih.is.BuildRecipeIngredient(recipeId, ingredientId)
		if err != nil {
			handleError(err, w)
			return
		}

		ingredient.Quantity = quantity

		ing, err := save(ingredient)
		if err != nil {
			handleError(err, w)
			return
		}

		cmp.FormUpdate(ing).Render(r.Context(), w)
	}
}

func (rih RecipeIngredientHandler) Delete() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (rih RecipeIngredientHandler) getQuantity(r *http.Request) (float32, error) {
	quantity, err := strconv.ParseFloat(r.FormValue("quantity"), 32)

	return float32(quantity), err
}

func (rih RecipeIngredientHandler) Search() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		recipeId := r.PathValue("recipe_id")
		search := r.URL.Query().Get("name")
		ingredients, err := rih.is.Search(search)

		if err != nil {
			handleError(err, w)
			return
		}

		cmp.SearchList(ingredients, recipeId, search).Render(r.Context(), w)
	}
}
