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

func CreateEstimate(estimateReq models.ReqEstimate) (models.EstimateModel, error) {
	newEstimate := models.EstimateModel{}

	infFarmLot, err := GetOneFarmLot(estimateReq.IdFarmLot)
	if err != nil {
		return newEstimate, err
	}

	newEstimate.CreateEstimation(estimateReq.TreesAssessed, infFarmLot)

	estimateCreated, err := repositories.CreateNewEstimation(newEstimate)
	if err != nil {
		return estimateCreated, err
	}

	return estimateCreated, nil
}

func GetEstimatesPerHarvest(reqIds models.ReqIdsEstimates) ([]models.EstimateModel, error) {
	return repositories.GetEstimatesPerHarvest(reqIds)
}
