package services

import (
	"AgroXpert-Backend/src/models"
	"AgroXpert-Backend/src/repositories"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllHarvests() ([]models.Harvest, error) {
	return repositories.GetAllHarvests()
}

func GetOneHarvest(HarvestID string) (models.Harvest, error) {
	return repositories.GetOneHarvest(HarvestID)
}

func GetHistoricHarvestEsimation(FarmLotID string) ([]models.HistoricHarvest, error) {
	var historic []models.HistoricHarvest
	var estimates []models.EstimateModel
	var finalProduction models.FinalProduction

	harvests, err := repositories.GetHarvestsByFarmLotID(FarmLotID)

	if err != mongo.ErrNoDocuments {
		for i := 0; i < len(harvests); i++ {

			fmt.Printf("Estimaciones: %v", harvests[i].Estimates)

			idsEstimates := models.ReqIdsEstimates{
				Ids: harvests[i].Estimates,
			}

			estimates, err = repositories.GetEstimatesPerHarvest(idsEstimates)

			if err == mongo.ErrNoDocuments {
				print("No hay estimaciones registradas")
				return historic, err
			}

			finalProduction, err = repositories.GetOneFinalProduction(string(harvests[i].SummaryFinalProduction.Hex()))

			if err == mongo.ErrNoDocuments {
				print("No hay produccion final registrada")
				return historic, err
			}

			historic = append(historic, models.HistoricHarvest{
				Harvest:         harvests[i],
				Estimates:       estimates,
				FinalProduction: finalProduction,
			})

		}
		return historic, err
	}
	print("No hay cosechas registradas")
	return historic, nil
}