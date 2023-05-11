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

func PostNewFinalProduction(finalProductionReq models.FinalProduction, idHarvest string) (models.FinalProduction, error) {

	newFinalProduction, err := repositories.PostNewFinalProduction(finalProductionReq)

	if err != nil {
		return models.FinalProduction{}, err
	}

	err = repositories.UpdateSummaryFinalProduction(idHarvest, newFinalProduction.ID)

	if err != nil {
		return models.FinalProduction{}, err
	}

	return newFinalProduction, nil
}
