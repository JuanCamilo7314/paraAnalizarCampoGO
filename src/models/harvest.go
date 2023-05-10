package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Harvest struct {
	ID                     primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Type                   string               `bson:"type,omitempty" json:"type,omitempty"`
	IDFarmLot              primitive.ObjectID   `bson:"idFarmLot,omitempty" json:"idFarmLot,omitempty"`
	EvaluationStartDate    time.Time            `bson:"evaluationStartDate,omitempty" json:"evaluationStartDate,omitempty"`
	EvaluationEndDate      time.Time            `bson:"evaluationEndDate,omitempty" json:"evaluationEndDate,omitempty"`
	Estimates              []primitive.ObjectID `bson:"estimates ,omitempty" json:"estimates ,omitempty"`
	SummaryFinalProduction primitive.ObjectID   `bson:"summaryFinalProduction,omitempty" json:"summaryFinalProduction,omitempty"`
}

type HarvestGraph struct {
	Type            string `bson:"type,omitempty" json:"type,omitempty"`
	Year            int    `bson:"year,omitempty" json:"year,omitempty"`
	TotalProduction int    `bson:"totalProduction,omitempty" json:"totalProduction,omitempty"`
}

type ProductionGraph struct {
	TotalProduction int       `bson:"totalProduction,omitempty" json:"totalProduction,omitempty"`
	Date            time.Time `bson:"date,omitempty" json:"date,omitempty"`
}
