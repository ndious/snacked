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

func (self RecipeIngredientHandler) getRecipeIngredientIds(r *http.Request)(recipeId, ingredientId int, err error) {
    var err1, err2 error
    
    recipeId, err1 = strconv.Atoi(r.PathValue("recipe_id"))

    ingredientId, err2 = strconv.Atoi(r.PathValue("ingredient_id"))

    return recipeId, ingredientId, errors.Join(err1, err2)
}

func (self RecipeIngredientHandler) FormSearch() func(w http.ResponseWriter, r *http.Request)  {
    return func(w http.ResponseWriter, r *http.Request) {
        recipeId, err := strconv.Atoi(r.PathValue("recipe_id"))
        if (err != nil) {
            handleError(err, w)
            return
        }

        ingredient := models.RecipeIngredient{
            RecipeID: recipeId,
        }

        cmp.FormNew(ingredient).Render(r.Context(), w)        
    }
}

func (self RecipeIngredientHandler) New() func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        recipeId, ingredientId, err := self.getRecipeIngredientIds(r)
        if (err != nil) {
            handleError(err, w)
            return
        }

        ingredient, err := self.is.BuildRecipeIngredient(recipeId, ingredientId)
        if (err != nil) {
            handleError(err, w)
            return
        }

        cmp.FormCreate(ingredient).Render(r.Context(), w)
    }
}

func (self RecipeIngredientHandler) Create() func(w http.ResponseWriter, r *http.Request)  {
    return self.handleSave(func (ingredient models.RecipeIngredient)(models.RecipeIngredient, error) {
        fmt.Println("Create", ingredient)
        return self.is.Create(ingredient)
    })
}

func (self RecipeIngredientHandler) Update() func(w http.ResponseWriter, r *http.Request)  {
    return self.handleSave(func (ingredient models.RecipeIngredient)(models.RecipeIngredient, error) {
        fmt.Println("Update", ingredient)
        return self.is.Update(ingredient)
    })
}

func (self RecipeIngredientHandler) handleSave(save func(ingredient models.RecipeIngredient)(models.RecipeIngredient, error)) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        recipeId, ingredientId, err := self.getRecipeIngredientIds(r)
        if (err != nil) {
            handleError(err, w)
            return
        }

        quantity, err := self.getQuantity(r)
        if (err != nil) {
            handleError(err, w)
            return
        }

        ingredient, err := self.is.BuildRecipeIngredient(recipeId, ingredientId)
        if (err != nil) {
            handleError(err, w)
            return
        }

        ingredient.Quantity = quantity
        
        ing, err := save(ingredient)
        if (err != nil) {
            handleError(err, w)
            return
        }

        cmp.FormUpdate(ing).Render(r.Context(), w)
    }
}

func (self RecipeIngredientHandler) Delete() func(w http.ResponseWriter, r *http.Request)  {
    return func(w http.ResponseWriter, r *http.Request) {
    }
}

func (self RecipeIngredientHandler) getQuantity(r *http.Request)(float32, error) {
    quantity, err := strconv.ParseFloat(r.FormValue("quantity"), 32)

    return float32(quantity), err
}

func (self RecipeIngredientHandler) Search() func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        recipeId := r.PathValue("recipe_id")
        searh := r.URL.Query().Get("name")
        ingredients, err := self.is.Search(searh)

        if (err != nil) {
            handleError(err, w)
            return
        }

        cmp.SearchList(ingredients, recipeId).Render(r.Context(), w)
    }
}

