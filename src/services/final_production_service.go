package services

import (
	"AgroXpert-Backend/src/models"
	"AgroXpert-Backend/src/repositories"
)

func GetAllFinalProductions() ([]models.FinalProduction, error) {
	return repositories.GetAllFinalProductions()
}

func GetOneFinalProduction(finalProductionID string) (models.FinalProduction, error) {
	return repositories.GetOneFinalProduction(finalProductionID)
}
