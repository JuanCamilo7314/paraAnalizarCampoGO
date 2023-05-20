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
	fmt.Println("harvestID: ", HarvestID)
	var modelHarvest models.Harvest
	collection := database.Db.GetCollection("Harvest")

	id, err := primitive.ObjectIDFromHex(HarvestID)
	if err != nil {
		return models.Harvest{}, fmt.Errorf("error convert id: %v", err)
	}

	filter := bson.M{"_id": id}
	harvest := collection.FindOne(context.Background(), filter)
	err = harvest.Decode(&modelHarvest)
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

func CreateHarvest(harvestReq models.CreateHarvest) (models.CreateHarvest, error) {
	collection := database.Db.GetCollection("Harvest")

	mapNewFarmLot := bson.M{
		"type":                   harvestReq.Type,
		"idFarmLot":              harvestReq.IDFarmLot,
		"evaluationStartDate":    harvestReq.EvaluationStartDate + "Z",
		"evaluationEndDate":      harvestReq.EvaluationEndDate + "Z",
		"summaryFinalProduction": nil,
		"estimates":              []primitive.ObjectID{},
	}

	result, err := collection.InsertOne(context.Background(), mapNewFarmLot)
	if err != nil {
		return models.CreateHarvest{}, fmt.Errorf("error insert farm lot: %v", err)
	}

	id := result.InsertedID.(primitive.ObjectID)
	harvestReq.ID = id

	return harvestReq, nil
}

func UpdateSummaryFinalProduction(idHarvest string, idFinalProduction primitive.ObjectID) error {
	collection := database.Db.GetCollection("Harvest")

	idHarvestUpdate, err := primitive.ObjectIDFromHex(idHarvest)
	if err != nil {
		return fmt.Errorf("error convert id: %v", err)
	}

	filter := bson.M{"_id": idHarvestUpdate}
	update := bson.M{"$set": bson.M{"summaryFinalProduction": idFinalProduction}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("error update summary final production: %v", err)
	}

	return nil
}

func UpdateEstimatesHarvest(idHarvest string, idNewEstimate primitive.ObjectID) error {
	collection := database.Db.GetCollection("Harvest")

	idHarvestUpdate, err := primitive.ObjectIDFromHex(idHarvest)
	if err != nil {
		return fmt.Errorf("error convert id: %v", err)
	}

	filter := bson.M{"_id": idHarvestUpdate}
	update := bson.M{
		"$push": bson.M{"estimates": idNewEstimate},
	}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("error update estimates harvest : %v ", err)
	}

	return nil
}

func GetHistoricHarvestEsimation(FarmLotID string) ([]models.HarvestDetails, error) {
	var resultHarvest []models.HarvestDetails
	var modelHarvestDetails models.HarvestDetails

	collection := database.Db.GetCollection("Harvest")
	id, err := primitive.ObjectIDFromHex(FarmLotID)

	pipelineHistoric := []bson.M{
		{
			"$match": bson.M{
				"idFarmLot": id,
			},
		},
		{
			"$lookup": bson.M{
				"from": "Estimates",
				"let":  bson.M{"idsEstimates": "$estimates"},
				"pipeline": []bson.M{
					{"$match": bson.M{
						"$expr": bson.M{
							"$in": []string{"$_id", "$$idsEstimates"},
						},
					},
					},
				},
				"as": "estimates",
			},
		},
		{
			"$lookup": bson.M{
				"from":         "FinalProduction",
				"localField":   "summaryFinalProduction",
				"foreignField": "_id",
				"as":           "summaryFinalProduction",
			},
		},

		{
			"$unwind": "$summaryFinalProduction",
		},
	}

	historic, err := collection.Aggregate(context.Background(), pipelineHistoric)
	err = historic.Decode(&modelHarvestDetails)

	if err == mongo.ErrNoDocuments {
		return []models.HarvestDetails{}, err
	}

	for historic.Next(context.Background()) {
		var lookup models.HarvestDetails
		err := historic.Decode(&lookup)
		if err != nil {
			return nil, err
		}
		resultHarvest = append(resultHarvest, lookup)
	}

	if err := historic.Err(); err != nil {
		return nil, err
	}

	return resultHarvest, nil

}
