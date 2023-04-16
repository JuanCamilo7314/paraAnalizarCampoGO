package repositories

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"AgroXpert-Backend/src/database"
	"AgroXpert-Backend/src/models"
)

func GetAllEstimatesProductions() ([]models.EstimateModel, error) {
	var resultEstimatesProductions []models.EstimateModel
	var modelEstimatesProduction models.EstimateModel
	collection := database.Db.GetCollection("Estimates")
	filter := bson.M{}

	estimatesProductions, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("error fiend all estimates of productions: %v", err)
	}

	for estimatesProductions.Next(context.Background()) {
		err := estimatesProductions.Decode(&modelEstimatesProduction)
		if err != nil {
			return nil, fmt.Errorf("error decode estimates productions: %v", err)
		}

		resultEstimatesProductions = append(resultEstimatesProductions, modelEstimatesProduction)
	}

	return resultEstimatesProductions, nil
}

func GetOneEstimatesProduction(estimatesProductionID string) (models.EstimateModel, error) {
	var modelEstimatesProduction models.EstimateModel
	collection := database.Db.GetCollection("Estimates")

	id, err := primitive.ObjectIDFromHex(estimatesProductionID)
	if err != nil {
		return models.EstimateModel{}, fmt.Errorf("error convert id: %v", err)
	}

	filter := bson.M{"_id": id}
	finalProduction := collection.FindOne(context.Background(), filter)
	err = finalProduction.Decode(&modelEstimatesProduction)
	if err == mongo.ErrNoDocuments {
		return models.EstimateModel{}, err
	}

	if err != nil {
		return models.EstimateModel{}, fmt.Errorf("error decode estimates production: %v", err)
	}

	return modelEstimatesProduction, nil
}

func CreateNewEstimation(newEstimation models.EstimateModel) (models.EstimateModel, error) {
	collection := database.Db.GetCollection("Estimates")

	mapNewEstimacion := bson.M{
		"date":                 newEstimation.Date,
		"numTrees":             newEstimation.NumTrees,
		"totalFruitsEstimates": newEstimation.TotalFruitsEstimates,
		"averageFruits":        newEstimation.AverageFruits,
		"estimatedProduction":  newEstimation.EstimatedProduction,
		"treesAssessed":        newEstimation.TreesAssessed,
	}

	result, err := collection.InsertOne(context.Background(), mapNewEstimacion)
	if err != nil {
		return models.EstimateModel{}, fmt.Errorf("error insert new estimation: %v", err)
	}

	newEstimation.ID = result.InsertedID.(primitive.ObjectID)

	return newEstimation, nil
}

func GetEstimatesPerHarvest(reqIds models.ReqIdsEstimates) ([]models.EstimateModel, error) {
	var resultEstimatesProductions []models.EstimateModel
	var estimateProduction models.EstimateModel
	collection := database.Db.GetCollection("Estimates")
	filter := bson.M{"_id": bson.M{"$in": reqIds.Ids}}

	estimatesProductions, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("error fiend all estimates of productions: %v", err)
	}

	for estimatesProductions.Next(context.Background()) {
		err := estimatesProductions.Decode(&estimateProduction)
		if err != nil {
			return nil, fmt.Errorf("error decode estimates productions: %v", err)
		}

		estimateFound, err := deepCopyObject(estimateProduction)
		if err != nil {
			return nil, fmt.Errorf("error deep copy object: %v", err)
		}

		resultEstimatesProductions = append(resultEstimatesProductions, estimateFound)
	}

	fmt.Printf("Result: %v", resultEstimatesProductions)
	return resultEstimatesProductions, nil
}

func deepCopyObject(estimate models.EstimateModel) (models.EstimateModel, error) {
	bytesEstimate, err := json.Marshal(estimate)
	if err != nil {
		return models.EstimateModel{}, fmt.Errorf("error marshal: %v", err)
	}

	var copyEstimates models.EstimateModel
	err = json.Unmarshal(bytesEstimate, &copyEstimates)
	if err != nil {
		return models.EstimateModel{}, fmt.Errorf("error unmarshal: %v", err)
	}

	return copyEstimates, nil
}
