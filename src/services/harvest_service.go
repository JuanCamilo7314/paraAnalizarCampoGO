package services

import (
	"AgroXpert-Backend/src/models"
	"AgroXpert-Backend/src/repositories"

	"go.mongodb.org/mongo-driver/mongo"
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

	if err == nil {
		for i := 0; i < len(harvests); i++ {

			if harvests[i].Estimates == nil {
				historic = append(historic, models.HistoricHarvest{
					Harvest: harvests[i],
				})
			} else {

				idsEstimates := models.ReqIdsEstimates{
					Ids: harvests[i].Estimates,
				}

				estimates, err = repositories.GetEstimatesPerHarvest(idsEstimates)

				if err != nil {
					estimates = []models.EstimateModel{}
				}

				finalProduction, err = repositories.GetOneFinalProduction(string(harvests[i].SummaryFinalProduction.Hex()))

				if err != nil {
					finalProduction = models.FinalProduction{}
				}

				historic = append(historic, models.HistoricHarvest{
					Harvest:         harvests[i],
					Estimates:       estimates,
					FinalProduction: finalProduction,
				})
			}

		}
		return historic, err
	} else if err == mongo.ErrNoDocuments {
		historic = []models.HistoricHarvest{}
	}

	return historic, nil
}
