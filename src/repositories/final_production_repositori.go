package repositories

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"AgroXpert-Backend/src/database"
	"AgroXpert-Backend/src/models"
	"AgroXpert-Backend/src/utils"
)

func GetAllFinalProductions() ([]models.FinalProduction, error) {
	var resultFinalProductions []models.FinalProduction
	var modelFinalProduction models.FinalProduction
	collection := database.Db.GetCollection("FinalProduction")
	filter := bson.M{}

	finalProductions, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("error fiend all final productions: %v", err)
	}

	for finalProductions.Next(context.Background()) {
		err := finalProductions.Decode(&modelFinalProduction)
		if err != nil {
			return nil, fmt.Errorf("error decode final productions: %v", err)
		}

		var finalProdCopy models.FinalProduction
		err = utils.DeepCopy(modelFinalProduction, &finalProdCopy)
		if err != nil {
			return nil, fmt.Errorf("error deep copy object: %v", err)
		}

		resultFinalProductions = append(resultFinalProductions, finalProdCopy)
	}

	return resultFinalProductions, nil
}

func GetOneFinalProduction(finalProductionID string) (models.FinalProduction, error) {
	var modelFinalProduction models.FinalProduction
	collection := database.Db.GetCollection("FinalProduction")

	id, err := primitive.ObjectIDFromHex(finalProductionID)
	if err != nil {
		return models.FinalProduction{}, fmt.Errorf("error convert id: %v", err)
	}

	filter := bson.M{"_id": id}
	finalProduction := collection.FindOne(context.Background(), filter)
	err = finalProduction.Decode(&modelFinalProduction)
	if err == mongo.ErrNoDocuments {
		return models.FinalProduction{}, err
	}

	if err != nil {
		return models.FinalProduction{}, fmt.Errorf("error decode final production: %v", err)
	}

	return modelFinalProduction, nil
}

func PostNewFinalProduction(finalProductionReq models.FinalProduction) (models.FinalProduction, error) {
	collection := database.Db.GetCollection("FinalProduction")

	mapNewFinalProduction := bson.M{
		"date":            finalProductionReq.Date,
		"totalProduction": finalProductionReq.TotalProduction,
		"exportMarket":    finalProductionReq.ExportMarket,
		"nationalMarket":  finalProductionReq.NationalMarket,
		"waste":           finalProductionReq.Waste,
		"caliberDivision": finalProductionReq.CaliberDivision,
	}

	result, err := collection.InsertOne(context.Background(), mapNewFinalProduction)
	if err != nil {
		return models.FinalProduction{}, fmt.Errorf("error insert final production: %v", err)
	}

	finalProductionReq.ID = result.InsertedID.(primitive.ObjectID)

	return finalProductionReq, nil
}
