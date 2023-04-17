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

func GetAllHarvests() ([]models.Harvest, error) {
	var resultHarvest []models.Harvest
	var modelHarvest models.Harvest
	collection := database.Db.GetCollection("Harvest")
	filter := bson.M{}

	harvest, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("error fiend all harvest: %v", err)
	}

	for harvest.Next(context.Background()) {
		err := harvest.Decode(&modelHarvest)
		if err != nil {
			return nil, fmt.Errorf("error decode harvest: %v", err)
		}

		resultHarvest = append(resultHarvest, modelHarvest)
	}

	return resultHarvest, nil
}

func GetOneHarvest(HarvestID string) (models.Harvest, error) {
	var modelHarvest models.Harvest
	collection := database.Db.GetCollection("Harvest")

	id, err := primitive.ObjectIDFromHex(HarvestID)
	if err != nil {
		return models.Harvest{}, fmt.Errorf("error convert id: %v", err)
	}

	filter := bson.M{"_id": id}
	Harvest := collection.FindOne(context.Background(), filter)
	err = Harvest.Decode(&modelHarvest)
	if err == mongo.ErrNoDocuments {
		return models.Harvest{}, err
	}

	if err != nil {
		return models.Harvest{}, fmt.Errorf("error decode farm lot: %v", err)
	}

	return modelHarvest, nil
}

func GetHarvestsByFarmLotID(FarmLotID string) ([]models.Harvest, error) {
	var resultHarvest []models.Harvest
	var modelHarvest models.Harvest
	collection := database.Db.GetCollection("Harvest")

	id, err := primitive.ObjectIDFromHex(FarmLotID)
	if err != nil {
		return nil, fmt.Errorf("error convert id: %v", err)
	}

	filter := bson.M{"idFarmLot": id}
	harvest, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("error fiend all harvest: %v", err)
	}

	for harvest.Next(context.Background()) {
		err := harvest.Decode(&modelHarvest)
		if err != nil {
			return nil, fmt.Errorf("error decode harvest: %v", err)
		}
		resultHarvest = append(resultHarvest, modelHarvest)
	}

	return resultHarvest, nil
}
