package services

import (
	"AgroXpert-Backend/src/models"
	"AgroXpert-Backend/src/repositories"
)

func GetAllEstimatesProductions() ([]models.EstimateModel, error) {
	return repositories.GetAllEstimatesProductions()
}

func GetOneEstimatesProduction(estimatesProductionID string) (models.EstimateModel, error) {
	return repositories.GetOneEstimatesProduction(estimatesProductionID)
}
