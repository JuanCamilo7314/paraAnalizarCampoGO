package services

import (
	"AgroXpert-Backend/src/models"
	"AgroXpert-Backend/src/repositories"
	"fmt"
)

func GetAllHarvests() ([]models.Harvest, error) {
	return repositories.GetAllHarvests()
}

func GetOneHarvest(HarvestID string) (models.Harvest, error) {
	return repositories.GetOneHarvest(HarvestID)
}

func CreateHarvest(harvestReq models.CreateHarvest) (models.CreateHarvest, error) {
	return repositories.CreateHarvest(harvestReq)
}

func GetHistoricHarvestEsimation(FarmLotID string) ([]models.HistoricHarvest, error) {
	var historic []models.HistoricHarvest
	var estimates []models.EstimateModel
	var finalProduction models.FinalProduction

	harvests, err := repositories.GetHarvestsByFarmLotID(FarmLotID)

	if err != nil {
		return []models.HistoricHarvest{}, err
	}

	for _, harvest := range harvests {
		idsEstimates := models.ReqIdsEstimates{
			Ids: harvest.Estimates,
		}

		estimates, err = repositories.GetEstimatesPerHarvest(idsEstimates)
		if err != nil {
			return []models.HistoricHarvest{}, err
		}

		finalProduction, err = repositories.GetOneFinalProduction(harvest.SummaryFinalProduction.Hex())
		if err != nil {
			return []models.HistoricHarvest{}, err
		}

		fmt.Printf("finalProduction: %+v\n", finalProduction)

		historic = append(historic, models.HistoricHarvest{
			Harvest:         harvest,
			Estimates:       estimates,
			FinalProduction: finalProduction,
		})
	}
	return historic, err
}
