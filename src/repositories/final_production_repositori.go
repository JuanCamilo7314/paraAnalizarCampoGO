package repositories

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"AgroXpert-Backend/src/database"
	"AgroXpert-Backend/src/models"
)

func GetAllFinalProductions() ([]models.FinalProduction, error) {
	var resultFinalProductions []models.FinalProduction
	var modelFinalProduction models.FinalProduction
	collection := database.Db.GetCollection("FinalProduction")
	filter := bson.M{}

	farms, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("error fiend all final productions: %v", err)
	}

	for farms.Next(context.Background()) {
		err := farms.Decode(&modelFinalProduction)
		if err != nil {
			return nil, fmt.Errorf("error decode final productions: %v", err)
		}

		resultFinalProductions = append(resultFinalProductions, modelFinalProduction)
	}

	return resultFinalProductions, nil
}
