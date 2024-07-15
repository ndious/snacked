package services

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/ndious/snacked/internal/models"
)

type RecipeService struct {
    db *sqlx.DB
}

type recipeService interface {
    CetAll()(recipes []models.Recipe, err error)
    Get(id int)(recipe models.Recipe, err error)
    GetSteps(recipeId int)(steps []models.Step, err error)
    GetIngredients(recipeId int)(ingredients []models.RecipeIngredient, err error)
    Create(name int)(recipe models.Recipe, err error)
    Update(id int, title, description string)(recipe models.Recipe, err error)
}

func NewRecipeService(db *sqlx.DB) *RecipeService {
    return &RecipeService {
        db: db,
    }
}

func (rs RecipeService) GetAll()(recipes []models.Recipe, err error) {
    recipes = []models.Recipe{}
    err = rs.db.Select(&recipes, "SELECT * FROM recipes")

    return recipes, err
}

func (rs RecipeService) Get(id int)(recipe models.Recipe, err error) {
    err = rs.db.Get(&recipe, "SELECT * FROM recipes WHERE id=$1", id)

    return recipe, err
}

func (rs RecipeService) Create(title string)(recipe models.Recipe, err error) {
    id := 0
    err = rs.db.QueryRow("INSERT into recipes (title) VALUES ($1) RETURNING id", title).Scan(&id)
    recipe.ID = id

    return recipe, err
}

func (rs RecipeService) Update(id int, title, description string)(recipe models.Recipe, err error) {
    recipe = models.Recipe{
        ID: id,
        Title: title,
        Description: sql.NullString{
            String: description,
        },
    }

    _, err = rs.db.NamedExec(`
    UPDATE recipes 
    SET title = :title, description = :description
    WHERE id = :id
    `, recipe)

    return recipe, err
}

func (rs RecipeService) GetIngredients(recipeId int)(ingredients []models.RecipeIngredient, err error) {
    err = rs.db.Select(&ingredients, `
    SELECT ri.id AS id, i.name AS name, ri.quantity AS quantity, i.unit AS unit
    FROM recipe_ingredients AS ri
    INNER JOIN ingredients as i ON ri.ingredient_id = i.id
    WHERE ri.recipe_id = $1
    `, recipeId)

    return ingredients, err
}

