package services

import (
	"AgroXpert-Backend/src/models"
	"AgroXpert-Backend/src/repositories"
)

func GetGraphsHarvest(FarmLotID string) ([]models.HarvestGraph, error) {
	return repositories.GetGraphsHarvest(FarmLotID)
}
