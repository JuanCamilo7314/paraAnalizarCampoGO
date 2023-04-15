package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EstimateModel struct {
	ID                   primitive.ObjectID `json:"id" bson:"_id"`
	Date                 time.Time          `json:"date" bson:"date"`
	NumTrees             int                `json:"numTrees" bson:"numTrees"`
	TotalFruitsEstimates int                `json:"totalFruitsEstimates" bson:"totalFruitsEstimates"`
	AverageFruits        float64            `json:"averageFruits" bson:"averageFruits"`
	EstimatedProduction  int                `json:"estimatedProduction" bson:"estimatedProduction"`
	TreesAssessed        []TreesAssessed    `json:"treesAssessed" bson:"treesAssessed"`
}

type TreesAssessed struct {
	NumFruits    int `json:"numFruits"`
	NumQuartiles int `json:"numQuartiles"`
}

type ReqEstimate struct {
	IdFarmLot     string          `json:"idFarm"`
	TreesAssessed []TreesAssessed `json:"treesAssessed"`
}
