package services

import (
	"AgroXpert-Backend/src/database"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func GetDataTest() {
	collection := database.Db.GetCollection("FarmLot")

	farms, err := collection.Find(context.Background(), nil)
	if err != nil {
		fmt.Println(err, "Error en el find")

		return
	}

	for farms.Next(context.Background()) {
		var result bson.M
		err := farms.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
}
