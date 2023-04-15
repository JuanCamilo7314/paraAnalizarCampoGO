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

		fmt.Printf("Estimates of Production: %v", modelEstimatesProduction)
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
