package repositories

import (
	"AgroXpert-Backend/src/database"
	"AgroXpert-Backend/src/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetGraphsHarvest(FarmLotID string) ([]models.HarvestGraph, error) {
	var resultHarvestGraph []models.HarvestGraph
	var productionGraph models.ProductionGraph
	var modelHarvest models.Harvest
	collectionHarvest := database.Db.GetCollection("Harvest")

	id, err := primitive.ObjectIDFromHex(FarmLotID)
	if err != nil {
		return nil, fmt.Errorf("error convert id: %v", err)

	}

	filter := bson.M{"idFarmLot": id}
	harvest, err := collectionHarvest.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("error find harvest: %v", err)
	}

	for harvest.Next(context.Background()) {
		err := harvest.Decode(&modelHarvest)
		if err != nil {
			return nil, err
		}

		productionGraph, err = GetGraphProduction(modelHarvest.SummaryFinalProduction)
		if err != nil {
			return nil, err
		}

		var harvestGraph models.HarvestGraph
		harvestGraph.Type = modelHarvest.Type
		harvestGraph.Year = productionGraph.Date.Year()
		harvestGraph.TotalProduction = productionGraph.TotalProduction

		resultHarvestGraph = append(resultHarvestGraph, harvestGraph)
	}

	return resultHarvestGraph, nil
}

func GetGraphProduction(FinalProduction primitive.ObjectID) (models.ProductionGraph, error) {
	var productionGraph models.ProductionGraph
	collectionProduction := database.Db.GetCollection("FinalProduction")

	filter := bson.M{"_id": FinalProduction}
	projection := bson.M{"totalProduction": 1, "date": 1}
	production := collectionProduction.FindOne(context.Background(), filter, options.FindOne().SetProjection(projection))

	err := production.Decode(&productionGraph)
	if err != nil {
		return models.ProductionGraph{}, fmt.Errorf("error decode production: %v", err)
	}

	return productionGraph, nil
}
