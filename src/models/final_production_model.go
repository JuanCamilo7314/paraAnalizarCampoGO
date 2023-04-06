package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FinalProduction struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Date            time.Time          `bson:"date,omitempty"`
	TotalProduction int                `bson:"totalProduction,omitempty"`
	ExportMarket    int                `bson:"exportMarket,omitempty"`
	NationalMarket  int                `bson:"nationalMarket,omitempty"`
	Waste           int                `bson:"waste,omitempty"`
	CaliberDivision []struct {
		Category string `bson:"category,omitempty"`
		Quantity int    `bson:"quantity,omitempty"`
	} `bson:"caliberDivision,omitempty"`
}
