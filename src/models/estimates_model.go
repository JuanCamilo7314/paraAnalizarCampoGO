package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EstimateModel struct {
	ID                   primitive.ObjectID `json:"id" bson:"_id"`
	Date                 time.Time          `json:"date" bson:"date"`
	NumTrees             int64              `json:"numTrees" bson:"numTrees"`
	TotalFruitsEstimates int64              `json:"totalFruitsEstimates" bson:"totalFruitsEstimates"`
	AverageFruits        int64              `json:"averageFruits" bson:"averageFruits"`
	EstimatedProduction  int64              `json:"estimatedProduction" bson:"estimatedProduction"`
	TreesAssessed        []TreesAssessed    `json:"treesAssessed" bson:"treesAssessed"`
}

type TreesAssessed struct {
	NumFruits    int64 `json:"numFruits"`
	NumQuartiles int64 `json:"numQuartiles"`
}

type ReqEstimate struct {
	IdFarmLot     primitive.ObjectID `json:"idFarm" bson:"_id"`
	TreesAssessed []TreesAssessed    `json:"treesAssessed" bson:"treesAssessed"`
}
