package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FinalProduction struct {
	ID              primitive.ObjectID `bson:"_id" json:"id"`
	Date            time.Time          `bson:"date" json:"date"`
	TotalProduction int                `bson:"totalProduction" json:"totalProduction"`
	ExportMarket    int                `bson:"exportMarket" json:"exportMarket"`
	NationalMarket  int                `bson:"nationalMarket" json:"nationalMarket"`
	Waste           int                `bson:"waste" json:"waste"`
	CaliberDivision []struct {
		Category string `bson:"category" json:"category"`
		Quantity int    `bson:"quantity" json:"quantity"`
	} `bson:"caliberDivision" json:"caliberDivision"`
}
