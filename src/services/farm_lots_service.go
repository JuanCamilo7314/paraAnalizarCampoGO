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

func CreateFarmLot(farmLotReq models.FarmLotReq) (models.FarmLotReq, error) {
	return repositories.CreateFarmLot(farmLotReq)
}
