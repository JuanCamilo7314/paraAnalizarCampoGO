package repositories

import (
	"AgroXpert-Backend/src/database"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func GetData() {
	collection := database.Db.GetCollection("FarmLot")
	filter := bson.M{}

	farms, err := collection.Find(context.Background(), filter)
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
