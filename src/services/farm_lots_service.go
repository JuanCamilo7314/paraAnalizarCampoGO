package services

import (
	"AgroXpert-Backend/src/models"
	"AgroXpert-Backend/src/repositories"
)

func GetAllFarmLots() ([]models.FarmLot, error) {
	return repositories.GetAllFarmLots()
}

func GetOneFarmLot(FarmLotID string) (models.FarmLot, error) {
	return repositories.GetOneFarmLot(FarmLotID)
}
