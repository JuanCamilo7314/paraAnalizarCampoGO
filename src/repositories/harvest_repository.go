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
	var modelHarvest models.Harvest

	idHarvestUpdate, err := primitive.ObjectIDFromHex(idHarvest)
	if err != nil {
		return fmt.Errorf("error convert id: %v", err)
	}
	fmt.Printf("idHarvestUpdate: %v", idHarvestUpdate)
	filter := bson.M{"_id": idHarvestUpdate}
	update := bson.M{
		"$addToSet": bson.M{"estimates": idNewEstimate},
	}

	// find, err := GetOneHarvest(idHarvest)
	// if err != nil {
	// 	fmt.Printf("error find harvest: %v,", err)
	// } else {
	// 	fmt.Printf("find: %v", find)
	// }

	harvest := collection.FindOne(context.Background(), filter)
	err = harvest.Decode(&modelHarvest)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("error decode harvest: %v", err)
	}

	if err != nil {
		fmt.Printf("error decode harvest: %v", err)
	}

	fmt.Printf("modelHarvest: %v", modelHarvest)

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("error update estimates harvest : %v ", err)
	}

	fmt.Printf("Documents modified: %v", result.UpsertedID)

	return nil
}
