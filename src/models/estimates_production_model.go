package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EstimatesProduct struct {
	ID              primitive.ObjectID `bson:"_id" json:"id"`
	TotalEstimatedFruits int                `bson:"TotalEstimatedFruits" json:"TotalEstimatedFruits"`
	AgeFruits    int                `bson:"AgeFruits" json:"AgeFruits"`
	AverageFruits  int                `bson:"AverageFruits" json:"AverageFruits"`
	EstimatesProduction   int                `bson:"EstimatesProduction" json:"wEstimatesProductionaste"`
	ThreesAssesed []struct {
		NumFruits string `bson:"NumFruits" json:"NumFruits"`
		NumQuartiles int    `bson:"NumQuartiles" json:"NumQuartiles"`
	} `bson:"ThreesAssesed" json:"ThreesAssesed"`
}
