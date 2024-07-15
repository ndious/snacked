package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/ndious/snacked/internal/components/layout"
	cmp "github.com/ndious/snacked/internal/components/recipes"
	"github.com/ndious/snacked/internal/services"
)

func RecipesRouter(r *http.ServeMux, db *sqlx.DB) (*Router, error) {
    rs := services.NewRecipeService(db)

    h := RecipeHandler{
        rs: rs,
    }

    r.HandleFunc("GET /recipes/new", h.new())

    r.HandleFunc("GET /recipes", h.GetAll())

    r.HandleFunc("POST /recipes", h.Create())

    r.HandleFunc("GET /recipes/{id}", h.Get())

    r.HandleFunc("PUT /recipes/{id}/update", h.Update())

    r.HandleFunc("GET /recipes/{id}/edit", h.Edit())

    return &Router{
        r,
        db,
    }, nil
}

type RecipeHandler struct {
    rs *services.RecipeService
}

func (rh RecipeHandler) new() func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        c := cmp.FormNew()
        layout.Page(c).Render(r.Context(), w)        
    }
}

func (rh RecipeHandler) GetAll() func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        recipes, err := rh.rs.GetAll()

        if (err != nil) {
            handleError(err, w)
            return 
        }

        c := cmp.Index(recipes)
        layout.Page(c).Render(r.Context(), w)
    }
}

func (rh RecipeHandler) Get() func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        id, err := strconv.Atoi(r.PathValue("id"))
        if (err != nil) {
            handleError(err, w)
            return
        }

        recipe, err := rh.rs.Get(id)
        if (err != nil) {
            handleError(err, w)
            return
        }
        fmt.Println(recipe)

        ingredients, err := rh.rs.GetIngredients(id)
        if (err != nil) {
            handleError(err, w)
            return
        }
        fmt.Println(ingredients)

		c := cmp.Show(recipe, ingredients)
		layout.Page(c).Render(r.Context(), w)
    }
}

func (rh RecipeHandler) Create() func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {

        title := r.FormValue("title")
        recipe, err := rh.rs.Create(title)

        if (err != nil) {
            handleError(err, w)
            return
        }
        
        http.Redirect(w, r, fmt.Sprintf("/recipes/%d/edit", recipe.ID), http.StatusFound)
    }
}

func (rh RecipeHandler) Edit() func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        id, err := strconv.Atoi(r.PathValue("id"))
        if (err != nil) {
            handleError(err, w)
            return
        }

        recipe, err := rh.rs.Get(id)
        if (err != nil) {
            handleError(err, w)
            return
        }

        ingredients, err := rh.rs.GetIngredients(id)
        if (err != nil) {
            handleError(err, w)
            return
        }

		c := cmp.FormUpdate(recipe, ingredients)
		layout.Page(c).Render(r.Context(), w)
    }
}


func (rh RecipeHandler) Update() func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {

        title := r.FormValue("title")
        description := r.FormValue("description")

        id, err := strconv.Atoi(r.PathValue("id"))
        if (err != nil) {
            handleError(err, w)
            return
        }

        recipe, err := rh.rs.Update(id, title, description)
        if (err != nil) {
            handleError(err, w)
            return
        }

        ingredients, err := rh.rs.GetIngredients(id)
        if (err != nil) {
            handleError(err, w)
            return
        }

        cmp.FormUpdate(recipe, ingredients).Render(r.Context(), w)
    }
}

