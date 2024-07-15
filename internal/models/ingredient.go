package models

type Units string

const (
    Kilogram Units = "kg"
    Gram           = "g"
    Miligram       = "mg"
    Liter          = "l"
    Milliliter     = "ml"
    Drop           = "dp"
    Tablespoon     = "tbs"
    Teaspoon       = "tsp"
    pintch         = "pt"
)

type Ingredient struct {
    ID int
    Name string
    Unit Units
}

type StepIngredient struct {
    ID int
    RecipeID int
    Name string
    Quantity float32
    Unit Units
}

type RecipeIngredient struct {
    ID int
    RecipeID int
    IngredientID int
    Name string
    Quantity float32
    Unit Units
}
