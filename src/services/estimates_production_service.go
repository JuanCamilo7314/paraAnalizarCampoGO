package services

import (
	"AgroXpert-Backend/src/models"
	"AgroXpert-Backend/src/repositories"
)

func GetAllEstimatesProductions() ([]models.EstimatesProduct, error) {
	return repositories.GetAllEstimatesProductions()
}

func GetOneEstimatesProduction(estimatesProductionID string) (models.EstimatesProduct, error) {
	return repositories.GetOneEstimatesProduction(estimatesProductionID)
}
