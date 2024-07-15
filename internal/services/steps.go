package services

import "github.com/ndious/snacked/internal/models"


type StepService interface {
    Get(id string)(step models.Step)
    GetIngredients(id string)(ingredients []models.StepIngredient)
    Create(description, duration string)(step models.Step)
    Update(id, description, duration string)(step models.Step)
}

