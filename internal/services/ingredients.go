package services

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/ndious/snacked/internal/models"
)

type IngredientService struct {
    db *sqlx.DB
}

func NewIngredientService(db *sqlx.DB) *IngredientService {
    return &IngredientService {
        db: db,
    }
}

type IngredientInterfaceService interface {
    GetRecipeIngredients(recipeId int)(ingredients []models.Ingredient)
    Get(id int)(ingredient models.Ingredient, err error)
    Create(recipeId int, name string, quantity float32, unit models.Units)(ingredient models.Ingredient, err error)
    Update(id int, name, quantity string)(ingredient models.Ingredient, err error)
}

func (is IngredientService) GetRecipeIngredient(recipeId, ingredientId int)(ingredient models.RecipeIngredient, err error) {
    err = is.db.Select(&ingredient, `
    SELECT ri.id AS id, i.name AS name, ri.quantity AS quantity, i.unit AS unit
    FROM recipe_ingredients AS ri
    INNER JOIN ingredients as i ON ri.ingredient_id = i.id
    WHERE ri.recipe_id == $1 AND ingredient_id == $2;
    `, recipeId, ingredientId)

    return ingredient, err
}

func (is IngredientService) Create(ingredient models.RecipeIngredient)(models.RecipeIngredient, error) {
    query := `
    INSERT INTO recipe_ingredients
        (recipe_id, ingredient_id, quantity)
    VALUES
        ($1, $2, $3)
    RETURNING id;
    `
    err := is.db.QueryRow(query, ingredient.RecipeID, ingredient.IngredientID, ingredient.Quantity).Scan(&ingredient.ID)

    return ingredient, err
}

func (is IngredientService) Update(ig models.RecipeIngredient)(models.RecipeIngredient, error) {
    ingredient, err1 := is.GetRecipeIngredient(ig.RecipeID, ig.IngredientID)
    query := `
    UPDATE recipe_ingredients
    SET quantity = :quantity
    WHERE id == :id;
    `

    ingredient.Quantity = ig.Quantity

    _, err2 := is.db.NamedExec(query, ingredient)

    return ingredient, errors.Join(err1, err2)
}

func (is IngredientService) Search(name string)(ingredients []models.Ingredient, err error) {
    fmt.Println(name)
    err = is.db.Select(&ingredients, "SELECT * FROM ingredients WHERE name ILIKE $1", "%"+name+"%")
    fmt.Println(ingredients)

    return ingredients, err
}

func (is IngredientService) GetIngredient(ingredientId int)(ingredient models.Ingredient, err error) {
    err = is.db.Get(&ingredient, "SELECT * FROM ingredients WHERE id = $1", ingredientId)
    fmt.Println(err, ingredient, ingredientId)

    return ingredient, err
}

func (is IngredientService) BuildRecipeIngredient(recipeId, ingredientId int)(models.RecipeIngredient, error) {
    
    ingredient, err := is.GetIngredient(ingredientId)
    if (err != nil) {
        return models.RecipeIngredient{}, err
    }

    return models.RecipeIngredient{
        RecipeID: recipeId,
        IngredientID: ingredientId,
        Name: ingredient.Name,
        Unit: ingredient.Unit,
        Quantity: .0,
    }, nil
}
