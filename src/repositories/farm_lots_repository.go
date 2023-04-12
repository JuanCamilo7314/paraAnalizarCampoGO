package repositories

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"AgroXpert-Backend/src/database"
	"AgroXpert-Backend/src/models"
)

func GetAllFarmLots() ([]models.FarmLot, error) {
	var resultFarmLot []models.FarmLot
	var modelFarmLot models.FarmLot
	collection := database.Db.GetCollection("FarmLot")
	filter := bson.M{}

	lots, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("error fiend all farm lots: %v", err)
	}

	for lots.Next(context.Background()) {
		err := lots.Decode(&modelFarmLot)
		if err != nil {
			return nil, fmt.Errorf("error decode farm lots: %v", err)
		}

		resultFarmLot = append(resultFarmLot, modelFarmLot)
	}

	return resultFarmLot, nil
}

func GetOneFarmLot(FarmLotID string) (models.FarmLot, error) {
	var modelFarmLot models.FarmLot
	collection := database.Db.GetCollection("FarmLot")

	id, err := primitive.ObjectIDFromHex(FarmLotID)
	if err != nil {
		return models.FarmLot{}, fmt.Errorf("error convert id: %v", err)
	}

	filter := bson.M{"_id": id}
	finalProduction := collection.FindOne(context.Background(), filter)
	err = finalProduction.Decode(&modelFarmLot)
	if err == mongo.ErrNoDocuments {
		return models.FarmLot{}, err
	}

	if err != nil {
		return models.FarmLot{}, fmt.Errorf("error decode farm lot: %v", err)
	}

	return modelFarmLot, nil
}
