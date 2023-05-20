package services

import (
	"AgroXpert-Backend/src/models"
	"AgroXpert-Backend/src/repositories"
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

func GetHistoricHarvestEsimation(FarmLotID string) ([]models.HarvestDetails, error) {
	return repositories.GetHistoricHarvestEsimation(FarmLotID)
}
